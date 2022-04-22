package models

import (
	"time"

	"github.com/lsflk/gig-sdk/libraries"
)

type NormalizedName struct {
	SearchText     string    `json:"search_text" bson:"search_text"`
	NormalizedText string    `json:"normalized_text" bson:"normalized_text"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
}

func (n NormalizedName) NewNormalizedName() NormalizedName {
	n.CreatedAt = time.Now()
	return n
}

func (n NormalizedName) GetSearchText() string {
	return n.SearchText
}

func (n NormalizedName) SetSearchText(value string) NormalizedName {
	n.SearchText = libraries.ProcessNameString(value)
	return n
}

func (n NormalizedName) GetNormalizedText() string {
	return n.NormalizedText
}

func (n NormalizedName) SetNormalizedText(value string) NormalizedName {
	n.NormalizedText = value
	return n
}

func (n NormalizedName) GetCreatedDate() time.Time {
	return n.CreatedAt
}
