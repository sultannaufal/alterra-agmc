package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `json:"-" form:"password"`
}
