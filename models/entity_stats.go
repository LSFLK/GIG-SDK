package models

import (
	"time"
)

/**
It is recommended to use get,set functions to access values of the entity.
Directly modify attributes only if you know what you are doing.
*/
type EntityStats struct {
	EntityCount            int                      `json:"entity_count" bson:"entity_count"`
	RelationCount          int                      `json:"relation_count" bson:"relation_count"`
	CategoryWiseCount      []CategoryWiseCount      `json:"category_wise_count" bson:"category_wise_count"`
	CategoryGroupWiseCount []CategoryGroupWiseCount `json:"category_group_wise_count" bson:"category_group_wise_count"`
	CreatedAt              time.Time                `json:"created_at" bson:"created_at"`
}

type CategoryGroupWiseCount struct {
	CategoryGroup []string `json:"_id" bson:"_id"`
	Count         int      `json:"category_count" bson:"category_count"`
}

type CategoryWiseCount struct {
	Category string `json:"_id" bson:"_id"`
	Count    int    `json:"category_count" bson:"category_count"`
}
