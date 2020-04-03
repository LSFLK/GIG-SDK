package tests

import (
	"GIG/sdk"
)

func (t *TestSDK) TestThatFileTypeCheckTrueWorks() {
	link := "http://www.buildings.gov.lk/index.php"
	result := sdk.FileTypeCheck(link,"php")
	t.AssertEqual(result, true)
}

func (t *TestSDK) TestThatFileTypeCheckFalseWorks() {
	link := "http://www.buildings.gov.lk/index.php"
	result := sdk.FileTypeCheck(link,"pdf")
	t.AssertEqual(result, false)
}