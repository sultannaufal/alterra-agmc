package migration

import (
	"example.com/architecture/database"
	"example.com/architecture/internal/model"
)

var tables = []interface{}{
	&model.User{},
}

func Migrate() {
	conn := database.GetConnection()
	conn.AutoMigrate(tables...)
}
