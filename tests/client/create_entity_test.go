package client

import (
	"github.com/lsflk/gig-sdk/models"
	"testing"
)

func TestThatCreateEntityWorks(t *testing.T) {
	initialEntity := models.Entity{Title: SriLanka}
	entity, _ := testClient.CreateEntity(initialEntity)

	if entity.GetTitle() != SriLanka {
		t.Errorf("create entity failed. %s != %s", entity.GetTitle(), SriLanka)
	}
}
