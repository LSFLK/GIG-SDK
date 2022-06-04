package libraries

import (
	"github.com/lsflk/gig-sdk/libraries"
	"testing"
)

func TestThatExtractDomainWorks(t *testing.T) {
	link := "https://www.buildings.gov.lk/index.php"
	result := libraries.ExtractDomain(link)
	expectedResult := "www.buildings.gov.lk"
	if result != expectedResult {
		t.Errorf("extract domain failed. %s != %s", result, expectedResult)
	}
}
