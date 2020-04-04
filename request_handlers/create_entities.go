package request_handlers

import (
	"GIG-SDK"
	"GIG-SDK/models"
	"encoding/json"
)

/**
Create a list of new entities and save to GIG
 */
func CreateEntities(entities []models.Entity) ([]models.Entity, error) {

	resp, err := PostRequest(config.ApiUrl+"add-batch", entities)
	if err != nil {
		return entities, err
	}
	json.Unmarshal([]byte(resp), &entities)

	return entities, err
}
