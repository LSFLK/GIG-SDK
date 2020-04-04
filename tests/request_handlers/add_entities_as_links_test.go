package request_handlers

import (
	"GIG-SDK/models"
	"GIG-SDK/request_handlers"
	"testing"
)

func TestThatAddEntitiesAsLinksWorks(t *testing.T) {
	linkEntity := models.Entity{Title: "Sri Lanka"}
	entity := models.Entity{Title: "test entity"}
	entity, _ = request_handlers.AddEntitiesAsLinks(entity, append([]models.Entity{}, linkEntity))

	if entity.Links[0].GetTitle() != "Sri Lanka" {
		t.Errorf("add entities as links failed. %s != %s", entity.Links[0].GetTitle(), "Sri Lanka")
	}

}
