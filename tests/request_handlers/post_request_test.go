package request_handlers

import (
	"GIG-SDK/request_handlers"
	"testing"
)

func TestThatPostRequestWorks(t *testing.T) {
	link := "https://en.wikipedia.org/w/api.php?action=query&format=json&titles=Sri%20Lanka&prop=extracts&exintro&explaintext"
	result, _ := request_handlers.PostRequest(link, "")
	t.AssertNotEqual(result, "")
}
