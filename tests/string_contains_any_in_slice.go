package tests

import (
	"GIG-SDK"
)

func (t *TestSDK) TestThatStringContainsAnyInSliceTestReturnsTrue() {
	testSlice := []string{"1", "2", "3", "4","here"}
	t.AssertEqual(libraries.StringContainsAnyInSlice(testSlice, "some value here"), true)
}

func (t *TestSDK) TestThatStringContainsAnyInSliceTestReturnsFalse() {
	testSlice := []string{"1", "2", "3", "4","some value here"}
	t.AssertEqual(libraries.StringInSlice(testSlice, "else"), false)
}
