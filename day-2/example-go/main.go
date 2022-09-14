package main

import (
	"example.com/go/config"
	"example.com/go/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
