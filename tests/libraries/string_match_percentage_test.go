package libraries

import (
	"github.com/lsflk/gig-sdk/libraries"
	"testing"
)

const (
	LongString = "long string"
)

func TestThatMatchStringWorksForEqualStrings(t *testing.T) {
	matchPercent := libraries.StringMatchPercentage(LongString, LongString)
	if matchPercent != 100 {
		t.Errorf("match string for equal strings failed. %d != %d", matchPercent, 100)
	}
}

func TestThatMatchStringWorksForUnequalStrings(t *testing.T) {
	matchPercent := libraries.StringMatchPercentage(LongString, "some other string")
	if matchPercent != 17 {
		t.Errorf("match string for unequal strings failed. %d != %d", matchPercent, 17)
	}
}

func TestThatMatchStringWorksForPartiallyEqualStrings(t *testing.T) {
	matchPercent := libraries.StringMatchPercentage("some what similar string", "som wht similar string")
	if matchPercent != 66 {
		t.Errorf("match string for partially equal strings failed. %d != %d", matchPercent, 66)
	}
}
