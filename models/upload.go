package models

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gobuffalo/buffalo/binding"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"

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
	t := time.Now().Unix()
	s := strconv.FormatInt(t, 10)
	u.Key = generateHash(s)
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
	dir := filepath.Join(".", "public", "uploads")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.WithStack(err)
	}
	f, err := os.Create(filepath.Join(dir, u.File.Filename))
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	_, err = io.Copy(f, u.File)

	// Update FilePath
	u.FilePath = f.Name()
	verrs, err := tx.ValidateAndUpdate(u)
	if verrs != nil {
		log.Println(verrs)
	}
	log.Println(err)

	return err
}

func generateHash(s string) string {
	// A hash needs to be 64 bytes long to have 256-bit collision resistance.
	h := make([]byte, 64)
	// Compute a 64-byte hash of buf and put it in h.
	sha3.ShakeSum256(h, []byte(s))
	return fmt.Sprintf("%x", h)
}
