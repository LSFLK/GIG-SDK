package request_handlers

import (
	"GIG-SDK"
	"GIG-SDK/models"
)

/**
Upload an image through API
 */
func UploadImage(payload models.Upload) error {

	if _, err := PostRequest(config.ApiUrl+"upload", payload); err != nil {
		return err
	}
	return nil
}
