package actions

import (
	"os"

	"github.com/edudorus/dead-man-switch/models"
	"github.com/markbates/willie"
)

func (as *ActionSuite) Test_UploadsResource_List() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_UploadsResource_Show() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_UploadsResource_New() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_UploadsResource_Create() {
	// clear out the uploads directory
	os.RemoveAll("./public/uploads")

	// setup a new Widget
	u := &models.Upload{FilePath: "/Example/IPFS/path", Key: "OurSecret"}

	// find the file we want to upload
	r, err := os.Open("./logo.svg")
	as.NoError(err)
	// setup a new willie.File to hold the file information
	f := willie.File{
		// ParamName is the name of the form parameter
		ParamName: "File",
		// FileName is the name of the file being uploaded
		FileName: r.Name(),
		// Reader is the file that is to be uploaded, any io.Reader works
		Reader: r,
	}

	// Post the Upload and the File(s) to /uploads
	res, err := as.HTML("/uploads").MultiPartPost(u, f)
	as.NoError(err)
	as.Equal(302, res.Code)

	// assert the file exists on disk
	_, err = os.Stat("./uploads/logo.svg")
	as.NoError(err)

	// assert the Widget was saved to the DB correctly
	as.NoError(as.DB.First(u))
	as.Equal("/Example/IPFS/path", u.FilePath)
	as.Equal("OurSecret", u.Key)
	as.NotZero(u.ID)
}

func (as *ActionSuite) Test_UploadsResource_Edit() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_UploadsResource_Update() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_UploadsResource_Destroy() {
	as.Fail("Not Implemented!")
}
