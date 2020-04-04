package request_handlers

import (
	"GIG-SDK/models"
	"GIG-SDK/request_handlers"
	"testing"
)

func TestThatAddEntityAsLinkWorks(t *testing.T) {
	linkEntity := models.Entity{Title: "Sri Lanka"}
	entity := models.Entity{Title: "test entity"}
	entity, _, _ = request_handlers.AddEntityAsLink(entity, linkEntity)

	if entity.Links[0].GetTitle() != "Sri Lanka" {
		t.Errorf("add entitiy as link failed. %s != %s", entity.Links[0].GetTitle(), "Sri Lanka")
	}

}
