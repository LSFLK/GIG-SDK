package libraries

import (
	"github.com/lsflk/gig-sdk/libraries"
	"testing"
)

func TestThatStringInSliceTestReturnsTrue(t *testing.T) {
	testSlice := []string{"1", "2", "3", "4"}
	result := libraries.StringInSlice(testSlice, "2")
	if result != true {
		t.Errorf("string in slice true failed. %t != %t", result, true)
	}
}

func TestThatStringInSliceTestReturnsFalse(t *testing.T) {
	testSlice := []string{"1", "2", "3", "4"}
	result := libraries.StringInSlice(testSlice, "5")
	if result != false {
		t.Errorf("string in slice false failed. %t != %t", result, false)
	}
}
