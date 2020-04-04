package request_handlers

import (
	"GIG-SDK/models"
	"GIG-SDK/request_handlers"
	"testing"
)

func TestThatAddEntityAsAttributeWorks(t *testing.T) {
	attributeEntity, err := request_handlers.GetEntity("Sri Lanka")
	if err != nil {
		t.Errorf("add entitiy as attribute, entity %s not found", "Sri Lanka")
	}

	entity := models.Entity{Title: "test entity"}
	entity, attributeEntity, _ = request_handlers.AddEntityAsAttribute(entity, "testAttribute", attributeEntity)
	result, expectedValue := entity.Attributes["testAttribute"].GetValue().GetValueString(), attributeEntity.GetTitle()

	if result != expectedValue || attributeEntity.Id.Hex() == "" {
		t.Errorf("add entitiy as attribute failed. %s != %s", result, expectedValue)
	}
}
