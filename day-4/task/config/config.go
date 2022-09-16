package config

import (
	"fmt"
	"os"

	"example.com/test/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {

	config := map[string]string{
		"DB_Username": os.Getenv("DB_USERNAME"),
		"DB_Password": os.Getenv("DB_PASSWORD"),
		"DB_Port":     os.Getenv("DB_PORT"),
		"DB_Host":     os.Getenv("DB_HOST"),
		"DB_Name":     os.Getenv("DB_NAME"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
}

func InitLog() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})
}
