package request_handlers

import (
	"GIG-SDK/models"
	"GIG-Scripts"
	"encoding/json"
)

/**
GetEntity
 */
func GetEntity(title string) (models.Entity, error) {
	var entity models.Entity
	resp, err := GetRequest(scripts.ApiUrl + "get/" + title)
	if err != nil {
		return entity, err
	}
	json.Unmarshal([]byte(resp), &entity)
	return entity, err
}
