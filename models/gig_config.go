package models

type GigConfig struct {
	ApiUrl                 string `json:"api_url" bson:"api_url"`
	ApiKey                 string `json:"api_key" bson:"api_key"`
	NerServerUrl           string `json:"ner_server_url" bson:"ner_server_url"`
	NormalizationServerUrl string `json:"normalization_url" bson:"normalization_url"`
	OcrServerUrl           string `json:"ocr_server_url" bson:"ocr_server_url"`
}
