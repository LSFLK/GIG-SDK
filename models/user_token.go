package models

// swagger:model
type UserToken struct {
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Role  string `json:"role" bson:"role"`
	Token string `json:"token" bson:"token"`
}
