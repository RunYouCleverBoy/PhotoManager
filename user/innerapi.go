package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"playgrounds.com/models"
)

func GetUserByEmail(email string) (*models.User, error) {
	return db.GetUserByEmail(email)
}

func CreateUserByUserObject(user *models.User) (*models.User, error) {
	return db.Create(user)
}

func GetUserById(id *primitive.ObjectID) (*models.User, error) {
	return db.Get(id)
}

func UpdateCredentials(id *primitive.ObjectID, password *string, token *string, refreshToken *string) (*models.User, error) {
	return db.UpdateCredentials(id, password, token, refreshToken)
}
