package libraries

import (
	"GIG-SDK/libraries"
	"os"
	"testing"
)

func TestThatDownloadFileWorks(t *testing.T) {
	filePath := "../../cache/test.pdf"
	os.Remove(filePath)
	link := "https://s1.q4cdn.com/806093406/files/doc_downloads/test.pdf"
	result := libraries.DownloadFile(filePath, link)
	if result != nil {
		t.Errorf("download file failed. %s != %#v", result.Error(), nil)
	}
}
