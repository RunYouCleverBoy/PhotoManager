package user

import (
	"playgrounds.com/models"
)

func GetUserByEmail(email string) (*models.User, error) {
	return db.GetUserByEmail(email)
}

func CreateUserByUserObject(user *models.User) (*models.User, error) {
	return db.Create(user)
}

func UpdateToken(id *string, token *string, expiration int64) (*models.User, error) {
	return db.UpdateCredentials(*id, nil, token, &expiration)
}
