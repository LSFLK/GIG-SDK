package libraries

import (
	"github.com/lsflk/gig-sdk/libraries"
	"testing"
)

func TestThatPdfParserWorks(t *testing.T) {
	result := libraries.ParsePdf("../../cache/test.pdf")
	if len(result) != 143 {
		t.Errorf("PDF parse failed. %d != %d", len(result), 143)
	}
}
