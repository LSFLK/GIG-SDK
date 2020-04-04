package request_handlers

import (
	"GIG-SDK/request_handlers"
	"testing"
)

func TestThatGetRequestWorks(t *testing.T) {
	link := "http://www.buildings.gov.lk/index.php"
	result, _ := request_handlers.GetRequest(link)
	if result == "" {
		t.Errorf("match string for partially equal strings failed. %s == %s", result, "")
	}
}
