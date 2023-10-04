package router

import (
	"github.com/gofiber/fiber/v2"
	"lovelcode/handlers"
)

func Route(app *fiber.App) error{
	
	apiOnly := app.Group("/", handlers.ApiOnly)

	auth := apiOnly.Group("/auth")
	auth.Post("/signin", handlers.Signin)
	auth.Post("/signup", handlers.Signup)

	apiOnly.Get("/", handlers.Home)

}