package libraries

import (
	"GIG-SDK/libraries"
	"testing"
)

func TestThatExtractDomainWorks(t *testing.T) {
	link := "http://www.buildings.gov.lk/index.php"
	result := libraries.ExtractDomain(link)
	expectedResult := "www.buildings.gov.lk"
	if result != expectedResult {
		t.Errorf("extract domain failed. %s != %s", result, expectedResult)
	}
}
