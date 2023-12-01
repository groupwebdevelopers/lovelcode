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
	"lovelcode/utils"
	"lovelcode/utils/s3"
	shandlers "lovelcode/handlers/statistics"
)


func main(){

	token.Setup()
	utils.Setup(shandlers.LogFunction)
	utils.Init()
	if err := s3.Init(); err!=nil{
		log.Println("cant create s3.")
	}
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
		log.Fatal("can't connect to database. err:", err)
	}
	//settings.Setup()
	
	router.Route(app)
	err := app.Listen(":3000")
	if err!=nil{
		log.Println(err)
		log.Fatal(app.Listen(":3131"))
	}
	
}