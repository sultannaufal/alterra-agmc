package database

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"example.com/architecture/internal/model"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Rdb *redis.Client
var books = []model.Book{
	{ID: 1, Title: "Go Tutorial", Isbn: strconv.Itoa(model.Isbn), Writer: "Anon", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

// var rDB

func CreateConnection() {

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
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

func CreateRedisConnection() {
	config := map[string]string{
		"REDIS_Username": os.Getenv("REDIS_USERNAME"),
		"REDIS_Password": os.Getenv("REDIS_PASSWORD"),
		"REDIS_Port":     os.Getenv("REDIS_PORT"),
		"REDIS_Host":     os.Getenv("REDIS_HOST"),
		"REDIS_Name":     os.Getenv("REDIS_NAME"),
	}

	connectionString := fmt.Sprintf("redis://%s:%s@%s:%s/%s",
		config["REDIS_Username"],
		config["REDIS_Password"],
		config["REDIS_Host"],
		config["REDIS_Port"],
		config["REDIS_Name"])

	opt, err := redis.ParseURL(connectionString)
	if err != nil {
		panic(err)
	}

	Rdb = redis.NewClient(opt)
	ttl := time.Duration(3) * time.Hour

	book_json, _ := json.Marshal(books)

	// store data
	op1 := Rdb.Set(context.Background(), "book", string(book_json), ttl)
	if err := op1.Err(); err != nil {
		fmt.Printf("unable to SET data. error: %v", err)
		return
	}
}

func InitMigrate() {
	DB.AutoMigrate(&model.User{})
}

func GetConnection() *gorm.DB {
	if DB == nil {
		CreateConnection()
	}
	if Rdb == nil {
		CreateRedisConnection()
	}
	return DB
}
