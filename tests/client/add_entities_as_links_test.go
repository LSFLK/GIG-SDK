package client

import (
	"github.com/lsflk/gig-sdk/models"
	"testing"
)

func TestThatAddEntitiesAsLinksWorks(t *testing.T) {
	linkEntity := models.Entity{Title: SriLanka}
	entity := models.Entity{Title: "test entity"}
	testClient.AddEntitiesAsLinks(&entity, append([]models.Entity{}, linkEntity))

	if entity.Links[0].GetTitle() != SriLanka {
		t.Errorf("add entities as links failed. %s != %s", entity.Links[0].GetTitle(), SriLanka)
	}

}
