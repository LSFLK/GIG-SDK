package tests

import "GIG-SDK"

func (t *TestSDK) TestThatMatchStringWorksForEqualStrings() {
	matchPercent:= libraries.StringMatchPercentage("long string","long string")
	t.AssertEqual(100, matchPercent)
}

func (t *TestSDK) TestThatMatchStringWorksForUnequalStrings() {
	matchPercent:= libraries.StringMatchPercentage("long string","some other string")
	t.AssertEqual(17, matchPercent)
}

func (t *TestSDK) TestThatMatchStringWorksForPartiallyEqualStrings() {
	matchPercent:= libraries.StringMatchPercentage("some what similar string","som wht similar stng")
	t.AssertEqual(66, matchPercent)
}
