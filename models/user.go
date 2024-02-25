package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type Role string

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Email    string             `json:"email" bson:"email,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
	Token    *string            `json:"token" bson:"token,omitempty"`
	Role     *Role              `json:"role" bson:"role,omitempty"`
}
