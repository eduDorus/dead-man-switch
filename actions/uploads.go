package actions

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"

	"github.com/edudorus/dead-man-switch/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"

	shell "github.com/ipfs/go-ipfs-api"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Upload)
// DB Table: Plural (uploads)
// Resource: Plural (Uploads)
// Path: Plural (/uploads)
// View Template Folder: Plural (/templates/uploads/)

// UploadsResource is the resource for the Upload model
type UploadsResource struct {
	buffalo.Resource
}

// List gets all Uploads. This function is mapped to the path
// GET /uploads
func (v UploadsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	uploads := &models.Uploads{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Uploads from the DB
	if err := q.All(uploads); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, uploads))
}

// Show gets the data for one Upload. This function is mapped to
// the path GET /uploads/{upload_id}
func (v UploadsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Upload
	upload := &models.Upload{}

	// To find the Upload the parameter upload_id is used.
	if err := tx.Find(upload, c.Param("upload_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, upload))
}

// New renders the form for creating a new Upload.
// This function is mapped to the path GET /uploads/new
func (v UploadsResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.Upload{}))
}

// Create adds a Upload to the DB. This function is mapped to the
// path POST /uploads
func (v UploadsResource) Create(c buffalo.Context) error {
	// Allocate an empty Upload
	upload := &models.Upload{}

	// Bind upload to the html form elements
	if err := c.Bind(upload); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	if err := UploadFile(upload); err != nil {
		log.Println("TILL HERE IT WORK")
		return errors.WithStack(err)
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(upload)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, upload))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Upload was created successfully")

	// and redirect to the uploads index page
	return c.Redirect(301, "/blocks")
}

// Edit renders a edit form for a Upload. This function is
// mapped to the path GET /uploads/{upload_id}/edit
func (v UploadsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Upload
	upload := &models.Upload{}

	if err := tx.Find(upload, c.Param("upload_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, upload))
}

// Update changes a Upload in the DB. This function is mapped to
// the path PUT /uploads/{upload_id}
func (v UploadsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Upload
	upload := &models.Upload{}

	if err := tx.Find(upload, c.Param("upload_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Upload to the html form elements
	if err := c.Bind(upload); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(upload)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, upload))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Upload was updated successfully")

	// and redirect to the uploads index page
	return c.Render(200, r.Auto(c, upload))
}

// Destroy deletes a Upload from the DB. This function is mapped
// to the path DELETE /uploads/{upload_id}
func (v UploadsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Upload
	upload := &models.Upload{}

	// To find the Upload the parameter upload_id is used.
	if err := tx.Find(upload, c.Param("upload_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(upload); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Upload was destroyed successfully")
	// Redirect to the uploads index page
	return c.Render(200, r.Auto(c, upload))
}

// AfterCreate will upload the file to IPFS Storage
func UploadFile(u *models.Upload) error {
	if !u.File.Valid() {
		return errors.New("File not valid")
	}

	var buf bytes.Buffer
	io.Copy(&buf, u.File)

	ed := encrypt(buf.Bytes(), u.Key)

	// upload and pin to IPFS
	sh := shell.NewShell("localhost:5001")
	ipfsHash, err := sh.Add(bytes.NewReader(ed))
	if err != nil {
		return errors.WithStack(err)
	}

	// Add ipfsHash to Blockchain
	err = UploadFileToBlockchain(ipfsHash, u.Address)
	if err != nil {
		return errors.WithStack(err)
	}

	// Update IPFSHash
	u.IPFSHash = ipfsHash
	return nil
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
