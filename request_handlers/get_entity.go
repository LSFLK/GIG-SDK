package request_handlers

import (
	"GIG-SDK"
	"GIG-SDK/models"
	"encoding/json"
)

/**
GetEntity
 */
func GetEntity(title string) (models.Entity, error) {
	var entity models.Entity
	resp, err := GetRequest(config.ApiUrl + "get/" + title)
	if err != nil {
		return entity, err
	}
	json.Unmarshal([]byte(resp), &entity)
	return entity, err
}
