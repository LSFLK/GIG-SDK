package libraries

import (
	"github.com/lsflk/gig-sdk/libraries"
	"testing"
)

func TestThatStringContainsAnyInSliceReturnsTrue(t *testing.T) {
	testSlice := []string{"1", "2", "3", "4", "here"}
	result := libraries.StringContainsAnyInSlice(testSlice, "some value here")
	if result != true {
		t.Errorf("string contains in any slice true failed. %t != %t", result, true)
	}
}

func TestThatStringContainsAnyInSliceReturnsFalse(t *testing.T) {
	testSlice := []string{"1", "2", "3", "4", "some value here"}
	result := libraries.StringInSlice(testSlice, "else")
	if result != false {
		t.Errorf("string contains in any slice false failed. %t != %t", result, false)
	}
}
