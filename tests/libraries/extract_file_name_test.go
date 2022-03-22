package libraries

import (
	"github.com/lsflk/gig-sdk/libraries"
	"testing"
)

func TestThatExtractFilenameWorks(t *testing.T) {
	filename := "registered%20suppliers%20and%20service%20%20providers%20for%20the%20year%202019%20-%20g1-office%20stationery.pdf"
	link := "/images/pdf/" + filename
	result := libraries.ExtractFileName(link)
	if result != filename {
		t.Errorf("extract filename failed. %s != %s", result, filename)
	}
}
