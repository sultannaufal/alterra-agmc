package repository

import (
	"example.com/architecture/database"
	"example.com/architecture/internal/model"
)

func FindByEmail(email string) (*model.User, error) {
	user := model.User{}
	if user := database.DB.Where("email = ?", email).First(&user); user.Error != nil {
		return nil, user.Error
	}
	return &user, nil
}
