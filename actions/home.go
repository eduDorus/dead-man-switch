package actions

import (
	"bytes"
	"io"

	"github.com/gobuffalo/buffalo"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/pkg/errors"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}

func LogHandler(c buffalo.Context) error {
	ReadFilesFromBlockchain()
	return c.Render(200, r.HTML("index.html"))
}

func DecryptHandler(c buffalo.Context) error {
	ipfsHash := c.Param("ipfs_hash")
	secret := c.Param("secret")

	// upload and pin to IPFS
	sh := shell.NewShell("localhost:5001")
	fileReader, err := sh.Cat(ipfsHash)
	if err != nil {
		return errors.WithStack(err)
	}

	var buf bytes.Buffer
	io.Copy(&buf, u.File)

	content := decrypt(buf.Bytes(), secret)
	c.Response().Write(content)

}
