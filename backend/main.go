package main

import (
	"log"
	"time"
	// "strings"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/template/html/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	
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

	app.Use(helmet.New())
	app.Use(csrf.New(csrf.Config{
		CookieName: "r",
		CookieSecure: true,
		CookieSessionOnly: true,
		CookieHTTPOnly: true,
		Expiration: 5 * time.Minute,
	}))
	app.Use(limiter.New(limiter.Config{
		Max: 20,
		Expiration: 2 * time.Minute,
	}))
	app.Use(logger.New(logger.Config{
		Format: "${green}${time}, ${blue}[${ip}]:${port} ${yellow}${status} - ${cyan}${method} ${magenta}${path}\n",
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