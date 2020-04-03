package tests

import (
	"GIG/sdk"
)

func (t *TestSDK) TestThatExtractDomainWorks() {
	link := "http://www.buildings.gov.lk/index.php"
	result := sdk.ExtractDomain(link)
	t.AssertEqual("www.buildings.gov.lk", result)
}
