package client

import (
	"testing"
)

func TestThatPostRequestWorks(t *testing.T) {
	link := "https://en.wikipedia.org/w/api.php?action=query&format=json&titles=Sri%20Lanka&prop=extracts&exintro&explaintext"
	result, _ := testClient.PostRequest(link, "")
	if result == "" {
		t.Errorf("post request failed. %s == %s", result, "")
	}
}
