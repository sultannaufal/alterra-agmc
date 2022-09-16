package main

import (
	"example.com/test/config"
	"example.com/test/routes"
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
