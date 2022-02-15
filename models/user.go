package models

import (
	"gopkg.in/mgo.v2/bson"
)

// swagger:model
type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Role     string        `json:"role" bson:"role"`
	Password []byte        `json:"-"`
	ApiKey   string        `json:"apikey" bson:"apikey"`
}

func (e User) GetId() bson.ObjectId {
	return e.Id
}
