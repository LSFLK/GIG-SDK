package tests

import (
	"GIG/sdk"
)

func (t *TestSDK) TestThatStringInSliceTestReturnsTrue() {
	testSlice := []string{"1", "2", "3", "4"}
	t.AssertEqual(sdk.StringInSlice(testSlice, "2"), true)
}

func (t *TestSDK) TestThatStringInSliceTestReturnsFalse() {
	testSlice := []string{"1", "2", "3", "4"}
	t.AssertEqual(sdk.StringInSlice(testSlice, "5"), false)
}
