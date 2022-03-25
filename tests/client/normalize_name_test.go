package client

import (
	"testing"
)

func TestThatNormalizeNameWorks(t *testing.T) {
	normalizedText, err := testClient.NormalizeName("sri lanka")
	if err != nil {
		t.Error("create entity from text failed", err)
	}
	if normalizedText!="Sri Lanka"{
		t.Error("normalization result incorrect")
	}
}
