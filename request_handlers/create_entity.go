package request_handlers

import (
	"GIG-SDK/models"
	"GIG-Scripts"
	"encoding/json"
)

/**
Create a new entity and save to GIG
 */
func CreateEntity(entity models.Entity) (models.Entity, error) {

	resp, err := PostRequest(scripts.ApiUrl+"add", entity)
	if err != nil {
		return entity, err
	}
	json.Unmarshal([]byte(resp), &entity)
	return entity, err
}
