package client

import (
	"github.com/lsflk/gig-sdk/models"
	"testing"
)

func TestThatAddEntityAsLinkWorks(t *testing.T) {
	linkEntity := models.Entity{Title: SriLanka}
	entity := models.Entity{Title: "test entity"}
	entity, _, _ = testClient.AddEntityAsLink(entity, linkEntity)

	if entity.Links[0].GetTitle() != SriLanka {
		t.Errorf("add entitiy as link failed. %s != %s", entity.Links[0].GetTitle(), SriLanka)
	}

}
