package repository

import (
	"example.com/architecture/database"
	"example.com/architecture/internal/model"
)

func FindAll() (interface{}, error) {
	var users []model.User

	if e := database.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func FindByID(id uint) (interface{}, error) {
	user := model.User{}
	if user := database.DB.Where("id = ?", id).First(&user); user.Error != nil {
		return nil, user.Error
	}
	return user, nil
}
