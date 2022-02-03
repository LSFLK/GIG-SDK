package models

import (
	"gopkg.in/mgo.v2/bson"
)

type NewUser struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Role     string        `json:"role" bson:"role"`
	Password string        `json:"password" bson:"password"`
}

func (e NewUser) GetId() bson.ObjectId {
	return e.Id
}
