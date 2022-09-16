package database

import (
	"example.com/test/config"
	"example.com/test/models"
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
	if user := config.DB.Where("id = ?", id).First(&user); user.Error != nil {
		return nil, user.Error
	}
	return user, nil
}
