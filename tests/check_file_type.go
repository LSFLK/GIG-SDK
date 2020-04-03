package tests

import (
	"GIG-SDK"
)

func (t *TestSDK) TestThatFileTypeCheckTrueWorks() {
	link := "http://www.buildings.gov.lk/index.php"
	result := libraries.FileTypeCheck(link,"php")
	t.AssertEqual(result, true)
}

func (t *TestSDK) TestThatFileTypeCheckFalseWorks() {
	link := "http://www.buildings.gov.lk/index.php"
	result := libraries.FileTypeCheck(link,"pdf")
	t.AssertEqual(result, false)
}