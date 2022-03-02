package models

/**
It is recommended to use get,set functions to access values of the entity.
Directly modify attributes only if you know what you are doing.
 */
type EntityStats struct {
	EntityCount            int                      `json:"entity_count" bson:"entity_count"`
	RelationCount          int                      `json:"relation_count" bson:"relation_count"`
	CategoryWiseCount      map[string]int           `json:"category_wise_count" bson:"category_wise_count"`
	CategoryGroupWiseCount []CategoryGroupWiseCount `json:"category_group_wise_count" bson:"category_group_wise_count"`
}

type CategoryGroupWiseCount struct {
	CategoryGroup []string `json:"_id" bson:"_id"`
	Count         int      `json:"category_count" bson:"category_count"`
}
