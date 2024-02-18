package models

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type Role string

type User struct {
	ID          string  `json:"id" bson:"id, omitempty"`
	Name        string  `json:"name" bson:"name"`
	Email       string  `json:"email" bson:"email"`
	Password    string  `json:"password" bson:"password, omitempty"`
	Token       *string `json:"token" bson:"token, omitempty"`
	TokenExpiry *int64  `json:"token_expiry" bson:"token_expiry, omitempty"`
	Role        *Role   `json:"role" bson:"role, omitempty"`
}
