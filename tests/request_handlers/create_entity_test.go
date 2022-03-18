package request_handlers

import (
	"GIG-SDK/models"
	"GIG-SDK/request_handlers"
	"testing"
)

func TestThatCreateEntityWorks(t *testing.T) {
	initialEntity := models.Entity{Title: SriLanka}
	entity, _ := request_handlers.CreateEntity(initialEntity)

	if entity.GetTitle() != SriLanka {
		t.Errorf("create entity failed. %s != %s", entity.GetTitle(), SriLanka)
	}
}
