package router

import (
	"github.com/gofiber/fiber/v2"
	"lovelcode/handlers"
)

func Route(app *fiber.App) {
	apiOnly := app.Group("/api", handlers.ApiOnly)
	apiV1 := apiOnly.Group("/v1")
	apiV1.Get("/", handlers.Home)

	auth := apiV1.Group("/auth")
	auth.Post("/signin", handlers.Signin)
	auth.Post("/signup", handlers.Signup)
	
	authReq := apiV1.Group("/", handlers.AuthRequired)
	authReq.Post("/project-doing-request", handlers.ProjectDoingRequest)

	
	app.Static("/", "../frontend/dist")
	app.Static("*", "../frontend/dist/index.html")

	apiOnly.Use(func (c *fiber.Ctx) error{
		return c.Status(404).JSON(fiber.Map{"error":"page not found"})
	})


	

}