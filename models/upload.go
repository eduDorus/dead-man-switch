package models

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"time"

	"github.com/ipfs/go-ipfs-api"

	"github.com/gobuffalo/buffalo/binding"
	"github.com/pkg/errors"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type Upload struct {
	ID        uuid.UUID    `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	FilePath  string       `json:"file_path" db:"file_path"`
	Key       string       `json:"key" db:"key"`
	File      binding.File `db:"-" form:"File"`
}

// String is not required by pop and may be deleted
func (u Upload) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Uploads is not required by pop and may be deleted
type Uploads []Upload

// String is not required by pop and may be deleted
func (u Uploads) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *Upload) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *Upload) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *Upload) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// AfterCreate will upload the file to IPFS Storage
func (u *Upload) AfterCreate(tx *pop.Connection) error {
	if !u.File.Valid() {
		return nil
	}

	var buf bytes.Buffer
	io.Copy(&buf, u.File)

	ed := encrypt(buf.Bytes(), u.Key)
	// upload and pin to IPFS
	sh := shell.NewShell("localhost:5001")
	cid, err := sh.Add(bytes.NewReader(ed))
	if err != nil {
		log.Println(errors.WithStack(err))
	}

	// Update FilePath
	u.FilePath = cid
	verrs, err := tx.ValidateAndUpdate(u)
	if verrs != nil {
		log.Println(verrs)
	}
	log.Println(err)

	return err
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}
