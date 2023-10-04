package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"lovelcode/router"
	"lovelcode/database"
)


func main(){

	engine := html.New("../views", ".html")
	
	app:= fiber.New(fiber.Config{
		Views: engine,
	})

	


	//add html views
	if err:=database.Setup(); err!=nil{
		log.Fatal("can't connect to database")
	}
	//settings.Setup()
	app.Static("/static", "../frontend/public")
	router.Route(app)
	log.Fatal(app.Listen(":8000"))
}