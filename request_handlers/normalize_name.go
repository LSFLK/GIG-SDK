package request_handlers

import (
	"GIG-SDK"
	"encoding/json"
	"github.com/revel/revel"
	"log"
	"net/url"
	"strconv"
)

type Response struct {
	Status  int    `json:"status"`
	Content string `json:"content"`
}

/**
normalize entity title before appending
 */
func NormalizeName(title string) (string, error) {

	normalizedName, err := GetRequest(config.NormalizeServer + "?searchText=" + url.QueryEscape(title))

	if err != nil {
		return "", err
	}
	var response Response
	if json.Unmarshal([]byte(normalizedName), &response); err != nil {
		return "", err
	}
	if response.Status != 200 {
		return "", revel.NewErrorFromPanic("Server responded with" + strconv.Itoa(response.Status))
	}
	log.Println("normalized", title, "to", response.Content)
	return response.Content, nil
}
