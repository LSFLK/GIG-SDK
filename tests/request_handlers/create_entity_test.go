package request_handlers

import (
	"GIG-SDK/models"
	"GIG-SDK/request_handlers"
	"testing"
)

func TestThatCreateEntityWorks(t *testing.T) {
	initialEntity := models.Entity{Title: "Sri Lanka"}
	entity, _ := request_handlers.CreateEntity(initialEntity)

	if entity.GetTitle() != "Sri Lanka" {
		t.Errorf("create entity failed. %s != %s", entity.GetTitle(), "Sri Lanka")
	}
}
