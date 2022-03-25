package client

import (
	"github.com/lsflk/gig-sdk/models"
	"testing"
)

func TestThatCreateEntityFromTextWorks(t *testing.T) {
	err := testClient.CreateEntityFromText("Test content for create entity from text", "Entity created from text", []string{"test"}, []models.NERResult{})
	if err != nil {
		t.Error("create entity from text failed", err)
	}
}
