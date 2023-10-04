package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"lovelcode/router"
)


func main(){

	engine := html.New("../views", ".html")
	
	app:= fiber.New(fiber.Config{
		Views: engine,
	})

	


	//add html views
	database.Setup()
	//settings.Setup()
	app.Static("/static", "../frontend/public")
	router.Route()
	log.Fatal(app.Listen(":8000"))
}