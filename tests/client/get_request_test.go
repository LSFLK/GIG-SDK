package client

import (
	"testing"
)

func TestThatGetRequestWorks(t *testing.T) {
	link := "http://www.buildings.gov.lk/index.php"
	result, _ := testClient.GetRequest(link)
	if result == "" {
		t.Errorf("match string for partially equal strings failed. %s == %s", result, "")
	}
}
