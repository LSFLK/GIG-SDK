package request_handlers

import (
	"GIG-SDK/models"
	"GIG-Scripts"
	"encoding/json"
)

/**
Create a list of new entities and save to GIG
 */
func CreateEntities(entities []models.Entity) ([]models.Entity, error) {

	resp, err := PostRequest(scripts.ApiUrl+"add-batch", entities)
	if err != nil {
		return entities, err
	}
	json.Unmarshal([]byte(resp), &entities)

	return entities, err
}
