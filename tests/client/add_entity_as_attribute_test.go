package client

import (
	"github.com/lsflk/gig-sdk/models"
	"testing"
)

func TestThatAddEntityAsAttributeWorks(t *testing.T) {
	attributeEntity, err := testClient.GetEntity(SriLanka)
	if err != nil {
		t.Errorf("add entitiy as attribute, entity %s not found", SriLanka)
	}

	entity := models.Entity{Title: "test entity"}
	entity, attributeEntity, _ = testClient.AddEntityAsAttribute(entity, "testAttribute", attributeEntity)
	result, expectedValue := entity.Attributes["testAttribute"].GetValue().GetValueString(), attributeEntity.GetTitle()

	if result != expectedValue {
		t.Errorf("add entitiy as attribute failed. %s != %s", result, expectedValue)
	}
}
