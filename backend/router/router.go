package router

import (
	"github.com/gofiber/fiber/v2"
	"lovelcode/handlers"
)

func Route(app *fiber.App) {
	apiOnly := app.Group("/api", handlers.ApiOnly)
	apiV1 := apiOnly.Group("/v1")
	apiV1.Get("/", handlers.Home)

	fileUploadAuthReq := app.Group("/api/v1/upload", handlers.AuthRequired)
	
	// auth not required
	
	apiV1.Post("/signin", handlers.Signin)
	apiV1.Post("/signup", handlers.Signup)
	
	// auth required
	authReq := apiV1.Group("/", handlers.AuthRequired)
	
	
	// admin required
	adminReq := authReq.Group("/", handlers.AdminRequired)
	
	// user
	adminReq.Post("/user/ban/:id", handlers.BanUser)
	
	// Project Doing Request
	adminReq.Post("/pdr/create", handlers.CreateProjectDoingRequest)
	adminReq.Get("/pdr/get/:id", handlers.GetProjectDoingRequest)
	adminReq.Get("/pdr/get-all", handlers.GetAllProjectDoingRequests)
	adminReq.Put("/pdr/edit/:id", handlers.EditProjectDoingRequest)
	adminReq.Delete("/pdr/delete/:id", handlers.DeleteProjectDoingRequest)
	
	// Plan
	adminReq.Post("/plan/create", handlers.CreatePlan)
	adminReq.Post("/plan/create-features/:planId", handlers.CreateFeatures)
	fileUploadAuthReq.Post("/plan/upload-image/:planId", handlers.UploadPlanImage)
	adminReq.Put("/plan/edit", handlers.EditPlan)
	adminReq.Put("/plan/edit-feature/:featureId", handlers.EditFeature)
	fileUploadAuthReq.Put("/plan/edit-image/:planId", handlers.UploadPlanImage)
	adminReq.Get("/plan/get-all-plans", handlers.GetAllPlans)
	adminReq.Get("/plan/get-plan/:planId", handlers.GetPlan)
	adminReq.Get("/plan/get-all-features/:planId", handlers.GetAllFeatures)
	adminReq.Get("/plan/get-feature/:featureId", handlers.GetFeature)
	adminReq.Get("/plan/get-all-plans-and-features", handlers.GetAllPlansAndFeatures)
	adminReq.Delete("/plan/delete-plan/:planId", handlers.DeletePlan) // todo:image must deleted
	adminReq.Delete("/plan/delete-feature/:featureId", handlers.DeleteFeature)
	
	// Static
	app.Static("/", "../frontend/dist")
	app.Static("*", "../frontend/dist/index.html")

	// apt not found
	apiOnly.Use(func (c *fiber.Ctx) error{
		return c.Status(404).JSON(fiber.Map{"error":"page not found"})
	})

	// static not found
	app.Use(func (c *fiber.Ctx) error{
		return c.SendStatus(404)
	})
	

}