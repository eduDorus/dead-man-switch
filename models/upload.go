package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gobuffalo/buffalo/binding"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
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
	return validate.Validate(
		&validators.StringIsPresent{Field: u.FilePath, Name: "FilePath"},
		&validators.StringIsPresent{Field: u.Key, Name: "Key"},
	), nil
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
	return err
}

func generateHash(v []byte) string {
	// A hash needs to be 64 bytes long to have 256-bit collision resistance.
	h := make([]byte, 64)
	// Compute a 64-byte hash of buf and put it in h.
	sha3.ShakeSum256(h, v)
	return fmt.Sprintf("%x", h)
}
