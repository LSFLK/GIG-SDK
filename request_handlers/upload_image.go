package request_handlers

import (
	"GIG-SDK/models"
	"GIG-Scripts"
)

/**
Upload an image through API
 */
func UploadImage(payload models.Upload) error {

	if _, err := PostRequest(scripts.ApiUrl+"upload", payload); err != nil {
		return err
	}
	return nil
}
