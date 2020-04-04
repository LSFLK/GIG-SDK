package request_handlers

import (
	"GIG-SDK/models"
	"GIG-SDK/request_handlers"
	"testing"
)

func TestThatCreateEntitiesWorks(t *testing.T) {
	initialEntity := models.Entity{Title: "Sri Lanka"}
	entities, _ := request_handlers.CreateEntities(append([]models.Entity{}, initialEntity))

	if entities[0].GetTitle() != "Sri Lanka" {
		t.Errorf("create entities failed. %s != %s", entities[0].GetTitle(), "Sri Lanka")
	}
}
