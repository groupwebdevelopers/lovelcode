package router

import (
	"github.com/gofiber/fiber/v2"
	"lovelcode/handlers"
)

func Route(app *fiber.App) {
	apiOnly := app.Group("/api", handlers.ApiOnly)
	apiV1 := apiOnly.Group("/v1")
	apiV1.Get("/", handlers.Home)

	// auth := apiV1.Group("/auth")
	apiV1.Post("/signin", handlers.Signin)
	apiV1.Post("/signup", handlers.Signup)

	authReq := apiV1.Group("/", handlers.AuthRequired)
	authReq.Post("/pdr/create", handlers.CreateProjectDoingRequest)
	authReq.Get("/pdr/get/:id", handlers.GetProjectDoingRequest)
	authReq.Get("/pdr/get-all", handlers.GetAllProjectDoingRequests)
	authReq.Post("/pdr/edit/:id", handlers.EditProjectDoingRequest)
	
	authReq.Post("/user/ban/:id", handlers.BanUser)

	authReq.Post("/plan/create", handlers.CreatePlan)

	app.Static("/", "../frontend/dist")
	app.Static("*", "../frontend/dist/index.html")

	apiOnly.Use(func (c *fiber.Ctx) error{
		return c.Status(404).JSON(fiber.Map{"error":"page not found"})
	})

	app.Use(func (c *fiber.Ctx) error{
		return c.SendStatus(404)
	})
	

}