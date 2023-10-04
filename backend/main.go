package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"lovelode/router"
)


func main(){
	app:= fiber.New()
	//add html views
	database.Setup()
	//settings.Setup()
	app.Static("/static", "../frontend/public")
	router.Route()
	log.Fatal(app.Listen(":8000"))
}