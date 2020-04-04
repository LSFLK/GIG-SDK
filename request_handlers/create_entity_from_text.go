package request_handlers

import (
	"GIG-SDK/models"
)

func CreateEntityFromText(textContent string, title string, categories []string, entityTitles []NERResult) error {
	//decode to entity
	var entities []models.Entity
	entity := models.Entity{}.
		SetTitle(models.Value{}.SetType("string").SetValueString("gazette")).
		SetAttribute("content", models.Value{}.
			SetType("string").
			SetValueString(textContent)).
		AddCategories(categories)

	for _, entityObject := range entityTitles {
		//normalizedName, err := utils.NormalizeName(entityObject.EntityName)
		//if err == nil {
		entities = append(entities, models.Entity{Title: entityObject.EntityName}.AddCategory(entityObject.Category))
		//}
	}

	entity, err := AddEntitiesAsLinks(entity, entities)
	if err != nil {
		panic(err)
	}

	//save to db
	entity, saveErr := CreateEntity(entity)

	return saveErr
}
