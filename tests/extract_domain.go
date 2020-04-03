package tests

import (
	"GIG-SDK/libraries"
)

func (t *TestSDK) TestThatExtractDomainWorks() {
	link := "http://www.buildings.gov.lk/index.php"
	result := libraries.ExtractDomain(link)
	t.AssertEqual("www.buildings.gov.lk", result)
}
