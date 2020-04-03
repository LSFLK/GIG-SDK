package tests

import (
	"GIG/sdk"
)

func (t *TestSDK) TestThatStringContainsAnyInSliceTestReturnsTrue() {
	testSlice := []string{"1", "2", "3", "4","here"}
	t.AssertEqual(sdk.StringContainsAnyInSlice(testSlice, "some value here"), true)
}

func (t *TestSDK) TestThatStringContainsAnyInSliceTestReturnsFalse() {
	testSlice := []string{"1", "2", "3", "4","some value here"}
	t.AssertEqual(sdk.StringInSlice(testSlice, "else"), false)
}
