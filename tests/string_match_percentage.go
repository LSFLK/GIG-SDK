package tests

import "GIG/sdk"

func (t *TestSDK) TestThatMatchStringWorksForEqualStrings() {
	matchPercent:= sdk.StringMatchPercentage("long string","long string")
	t.AssertEqual(100, matchPercent)
}

func (t *TestSDK) TestThatMatchStringWorksForUnequalStrings() {
	matchPercent:= sdk.StringMatchPercentage("long string","some other string")
	t.AssertEqual(17, matchPercent)
}

func (t *TestSDK) TestThatMatchStringWorksForPartiallyEqualStrings() {
	matchPercent:= sdk.StringMatchPercentage("some what similar string","som wht similar stng")
	t.AssertEqual(66, matchPercent)
}
