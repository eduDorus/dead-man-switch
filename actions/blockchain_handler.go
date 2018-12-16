package actions

import "github.com/gobuffalo/buffalo"

func BlockHandler(c buffalo.Context) error {
	files := ReadFilesFromBlockchain()
	c.Set("files", files)
	return c.Render(200, r.HTML("blocklist.html"))
}
