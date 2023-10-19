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
	fileUploadAdminReq := app.Group("/admin/upload", handlers.AdminUploadImage)
	fileUploadAdminReq.Post("/plan/image/:planId", handlers.UploadPlanImage)
	fileUploadAdminReq.Post("/member/image/:memberId", handlers.UploadMemberImage)
	
	// auth not required
	
	// auth
	apiV1.Post("/signin", handlers.Signin)
	apiV1.Post("/signup", handlers.Signup)
	
	// plan
	apiV1.Get("/plan/get-all-plans", handlers.GetAllPlans)
	apiV1.Get("/plan/get-plan/:planId", handlers.GetPlan)
	apiV1.Get("/plan/get-all-features/:planId", handlers.GetAllFeatures)
	apiV1.Get("/plan/get-feature/:featureId", handlers.GetFeature)
	apiV1.Get("/plan/get-all-plans-and-features", handlers.GetAllPlansAndFeatures)
	
	// member
	apiV1.Get("/member/get-all", handlers.GetAllMembers)
	apiV1.Get("/member/get/:memberId", handlers.GetMember)

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
	adminReq.Delete("/plan/delete-plan/:planId", handlers.DeletePlan) // todo:image must deleted
	adminReq.Delete("/plan/delete-feature/:featureId", handlers.DeleteFeature)
	
	
	// member
	adminReq.Post("/member/create/:userId", handlers.CreateMember)
	adminReq.Put("/member/edit/:memberId", handlers.EditMember)
	adminReq.Delete("/member/delete/:memberId", handlers.DeleteMember)

	// file upload admin required
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