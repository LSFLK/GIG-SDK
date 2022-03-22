package client

import (
	"GIG-SDK/models"
	"testing"
)

func TestThatCreateEntitiesWorks(t *testing.T) {
	initialEntity := models.Entity{Title: SriLanka}
	entities, _ := testClient.CreateEntities(append([]models.Entity{}, initialEntity))

	if entities[0].GetTitle() != SriLanka {
		t.Errorf("create entities failed. %s != %s", entities[0].GetTitle(), SriLanka)
	}
}
