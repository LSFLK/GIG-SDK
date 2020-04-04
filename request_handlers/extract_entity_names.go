package request_handlers

import (
	"GIG-SDK"
	"GIG-SDK/models"
	"encoding/json"
)

/**
NER extraction
 */
func ExtractEntityNames(textContent string) ([]models.NERResult, error) {

	apiResp, err := PostRequest(config.NERServerUrl, textContent)
	if err != nil {
		return nil, err
	}

	var (
		entityTitles [][]string
		results      []models.NERResult
	)
	if err = json.Unmarshal([]byte(apiResp), &entityTitles);err!=nil{
		return nil, err
	}

	for _, entity := range entityTitles {
		newNERResult := models.NERResult{}.SetCategory(entity[1]).SetEntityName(entity[0])
		results = append(results, newNERResult)
	}
	return results, nil
}
