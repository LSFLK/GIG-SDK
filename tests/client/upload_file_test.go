package client

import (
	"github.com/lsflk/gig-sdk/models"
	"testing"
)

func TestThatUploadFileWorks(t *testing.T) {
	 err := testClient.UploadFile(models.Upload{Title:"test upload file", Source:"../../cache/test.pdf"})
	if err != nil {
		t.Error("file upload failed", err)
	}
}
