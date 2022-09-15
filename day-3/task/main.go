package main

import (
	"example.com/middleware/config"
	"example.com/middleware/routes"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	config.InitDB()
	config.InitLog()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
