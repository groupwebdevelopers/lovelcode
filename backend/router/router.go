package router

import (
	"github.com/gofiber/fiber/v2"
	"lovelcode/handlers"
)

func Route(app *fiber.App) {
	apiOnly := app.Group("/api", handlers.ApiOnly)
	apiV1 := apiOnly.Group("/v1")
	apiV1.Get("/", handlers.Home)
	
	// fileUploadAuthReq := app.Group("/api/v1/upload", handlers.AuthRequired)
	
	// auth not required
	
	apiV1.Post("/signin", handlers.Signin)
	apiV1.Post("/signup", handlers.Signup)
	
	// auth required
	authReq := apiV1.Group("/", handlers.AuthRequired)
	
	// Project Doing Request
	authReq.Post("/pdr/create", handlers.CreateProjectDoingRequest)
	authReq.Get("/pdr/get/:id", handlers.GetProjectDoingRequest)
	authReq.Get("/pdr/get-all", handlers.GetAllProjectDoingRequests)
	authReq.Put("/pdr/edit/:id", handlers.EditProjectDoingRequest)
	authReq.Delete("/pdr/delete/:id", handlers.DeleteProjectDoingRequest)
	
	
	// admin required
	adminReq := authReq.Group("/admin", handlers.AdminRequired)
	
	// user
	adminReq.Post("/user/ban/:id", handlers.BanUser)
	
	// Plan
	adminReq.Post("/plan/create", handlers.CreatePlan)
	adminReq.Post("/plan/create-features/:planId", handlers.CreateFeatures)
	adminReq.Put("/plan/edit/:planId", handlers.EditPlan)
	adminReq.Put("/plan/edit-feature/:featureId", handlers.EditFeature)
	adminReq.Get("/plan/get-all-plans", handlers.GetAllPlans)
	adminReq.Get("/plan/get-plan/:planId", handlers.GetPlan)
	adminReq.Get("/plan/get-all-features/:planId", handlers.GetAllFeatures)
	adminReq.Get("/plan/get-feature/:featureId", handlers.GetFeature)
	adminReq.Get("/plan/get-all-plans-and-features", handlers.GetAllPlansAndFeatures)
	adminReq.Delete("/plan/delete-plan/:planId", handlers.DeletePlan) // todo:image must deleted
	adminReq.Delete("/plan/delete-feature/:featureId", handlers.DeleteFeature)
	
	fileUploadAdminReq := app.Group("/upload", handlers.UploadAdminImage)
	fileUploadAdminReq.Post("/upload-plan-image/:planId", handlers.UploadPlanImage)
	
	
	// apt not found
	apiOnly.Use(func (c *fiber.Ctx) error{
		return c.Status(404).JSON(fiber.Map{"error":"page not found","routes":app.GetRoutes()})
	})
	
	// Static
	app.Static("/", "../frontend/dist")
	app.Static("*", "../frontend/dist/index.html")


	// static not found
	app.Use(func (c *fiber.Ctx) error{
		return c.SendStatus(404)
	})
	

}