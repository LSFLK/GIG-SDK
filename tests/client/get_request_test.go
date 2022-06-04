package client

import (
	"testing"
)

func TestThatGetRequestWorks(t *testing.T) {
	link := "https://www.buildings.gov.lk/index.php"
	result, _ := testClient.GetRequest(link)
	if result == "" {
		t.Errorf("get request failed. %s == %s", result, "")
	}
}
