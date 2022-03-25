package client

import (
	"testing"
)

func TestThatExtractEntityNamesWorks(t *testing.T) {
	_, err := testClient.ExtractEntityNames("John who lives in Colombo")
	if err != nil {
		t.Error("extract entity names failed:", err)
	}
}
