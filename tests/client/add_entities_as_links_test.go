package client

import (
	"github.com/lsflk/gig-sdk/models"
	"testing"
)

func TestThatAddEntitiesAsLinksWorks(t *testing.T) {
	linkEntity := models.Entity{Title: SriLanka}
	entity := models.Entity{Title: "test entity"}
	err := testClient.AddEntitiesAsLinks(&entity, append([]models.Entity{}, linkEntity))
	if err != nil {
		t.Error("error adding entities as links:", err)
	}
	if entity.Links[0].GetTitle() != SriLanka {
		t.Errorf("add entities as links failed. %s != %s", entity.Links[0].GetTitle(), SriLanka)
	}

}
