package main

import (
	"log"
	// "strings"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/template/html/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"lovelcode/router"
	"lovelcode/database"
	"lovelcode/utils/token"
)


func main(){

	token.Setup()

	// engine := html.New("../frontend", ".html")
	
	app:= fiber.New()//fiber.Config{
		// Views: engine,
	// })
	// todo: must edited in production
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// remove
	// _ = database.DB
	
	if err:=database.Setup(); err!=nil{
		log.fatal("can't connect to database. err:", err)
	}
	//settings.Setup()
	
	router.Route(app)
	log.Fatal(app.Listen(":3000"))
	
}