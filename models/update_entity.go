package models

type UpdateEntity struct {
	Title           string `json:"title"`
	SearchAttribute string `json:"search_attribute"`
	SearchValue     Value  `json:"search_value"`
	Attribute       string `json:"attribute"`
	Value           Value  `json:"value"`
}
