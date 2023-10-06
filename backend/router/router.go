package router

import (
	"github.com/gofiber/fiber/v2"
	"lovelcode/handlers"
)

func Route(app *fiber.App) {
	
	apiOnly := app.Group("/", handlers.ApiOnly)
	apiV1 := apiOnly.Group("/v1")

	auth := apiV1.Group("/auth")
	auth.Post("/signin", handlers.Signin)
	auth.Post("/signup", handlers.Signup)

	apiV1.Get("/", handlers.Home)
	
}