package database

import (
	"example.com/crud/config"
	"example.com/crud/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func GetUserByID(id string) (interface{}, error) {
	user := models.User{}
	if user := config.DB.First(&user).Where("id = ?", id); user.Error != nil {
		return nil, user.Error
	}
	return user, nil
}
