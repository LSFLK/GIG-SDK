package client

import (
	"testing"
)

func TestThatGetEntityWorks(t *testing.T) {
	result, err := testClient.GetEntity("Sri Lanka")
	if err != nil {
		t.Errorf("match string for partially equal strings failed. %s == %s", result, "")
	}
}
