package client

import (
	"testing"
)

func TestThatExtractEntityNamesWorks(t *testing.T) {
	_, err := testClient.ExtractEntityNames("Test content for create entity from text")
	if err != nil {
		t.Error("create entity from text failed", err)
	}
}
