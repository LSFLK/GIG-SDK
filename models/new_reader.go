package models

// swagger:model
type NewReader struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
