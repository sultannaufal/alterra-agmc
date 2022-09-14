package main

import (
	"example.com/crud/config"
	"example.com/crud/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
