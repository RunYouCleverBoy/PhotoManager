package models

type User struct {
	ID          string  `json:"id" bson:"id, omitempty"`
	Name        string  `json:"name" bson:"name"`
	Email       string  `json:"email" bson:"email"`
	Token       *string `json:"token" bson:"token, omitempty"`
	TokenExpiry *int64  `json:"token_expiry" bson:"token_expiry, omitempty"`
}
