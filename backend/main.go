package main

import (
//	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"lovelcode/router"
	"lovelcode/database"
"fmt"
"mime"
)


func main(){

	engine := html.New("../frontend", ".html")
	
	app:= fiber.New(fiber.Config{
		Views: engine,
	})
	// todo: must edited in production
	app.Use(cors.New())//cors.Config{
	//	AllowOrigins: "*",
	//	AllowHeaders: "*",
	//}))

	// remove
	_ = database.DB
	
	//if err:=database.Setup(); err!=nil{
	//	log.Fatal("can't connect to database")
	//}
	//settings.Setup()
	app.Static("/", "../frontend", fiber.Static{ModifyResponse: func(c *fiber.Ctx) error{
		if strings.Contains(c.GetRespHeader("Content-Type"), "text/plain"){
			c.Response().Header.Set("Content-Type", "application/javascript")
			c.Response().Header.Set("Connection", "keep-alive")
			c.Response().Header.Set("Keep-Alive", "timeout=5")
			c.Response().Header.Set("Cache-Control", "no-cache")
			fmt.Println(c.OriginalURL())
		}
		fmt.Println(c.GetRespHeader("Content-Type"))
		return nil
	}})
	//app.Get("/src/main.jsx", func (c *fiber.Ctx) error{
	//	c.Re
	//})
	
	router.Route(app)
	//log.Fatal(app.Listen(":8000"))
	fmt.Println("jsx",mime.TypeByExtension(".jsx"))
}