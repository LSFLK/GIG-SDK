package models

// swagger:model
type User struct {
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Role     string `json:"role" bson:"role"`
	Password []byte `json:"-"`
	ApiKey   string `json:"apikey" bson:"apikey"`
}
