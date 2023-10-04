package router

import (
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App){
	
	apiOnly := app.Group("/", handlers.ApiOnly)

	auth := apiOnly.Group("/auth")
	auth.Post("/signin", handlers.signin)
	auth.Post("/signup", handlers.signup)

	apiOnly.Get("/", handlers.home)
	
}