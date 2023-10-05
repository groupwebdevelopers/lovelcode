package main

import (
	"log"
	// "strings"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/template/html/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"lovelcode/router"
	"lovelcode/database"

)


func main(){

	// engine := html.New("../frontend", ".html")
	
	app:= fiber.New()//fiber.Config{
		// Views: engine,
	// })
	// todo: must edited in production
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// remove
	_ = database.DB
	
	//if err:=database.Setup(); err!=nil{
	//	log.Fatal("can't connect to database")
	//}
	//settings.Setup()
	app.Static("/", "../frontend/dist")
	
	router.Route(app)
	log.Fatal(app.Listen(":8000"))
	
}