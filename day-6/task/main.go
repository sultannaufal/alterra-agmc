package main

import (
	"example.com/architecture/database"
	"example.com/architecture/internal/http"
	"example.com/architecture/internal/middleware"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	database.CreateConnection()
	database.CreateRedisConnection()
	e := echo.New()
	middleware.Init(e)
	http.NewHttp(e)
	e.Logger.Fatal(e.Start(":8000"))
}
