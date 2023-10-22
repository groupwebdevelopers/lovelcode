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
	fileUploadAdminReq.Post("/article/image/:articleId", handlers.UploadArticleImage)
	
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
	
	// article
	apiV1.Get("/article/get/:articleTitleUrl", handlers.GetArticle)
	apiV1.Get("/article/get-all/:page", handlers.GetAllArticlesTitles)
	
	// set features
	apiV1.Get("/site-features/get-all", handlers.GetSiteFeatures)

	// auth required
	authReq := apiV1.Group("/a", handlers.AuthRequired)
	
	// Project Doing Request
	authReq.Post("/pdr/create", handlers.CreateProjectDoingRequest)
	authReq.Get("/pdr/get/:id", handlers.GetProjectDoingRequest)
	authReq.Get("/pdr/get-all", handlers.GetAllProjectDoingRequests)
	authReq.Put("/pdr/edit/:id", handlers.EditProjectDoingRequest)
	authReq.Delete("/pdr/delete/:id", handlers.DeleteProjectDoingRequest)
	
	
	// admin required

	// article
	adminArticleReq := authReq.Group("/admin/article", handlers.AdminArticleRequired)
	adminArticleReq.Post("/create", handlers.CreateArticle)
	adminArticleReq.Put("/edit/:articleId", handlers.EditArticle)
	adminArticleReq.Delete("/delete/:articleId", handlers.DeleteArticle)


	// adminReq := authReq.Group("/admin", handlers.AdminRequired)
	
	// user
	userAdminReq := authReq.Group("/admin/user", handlers.AdminRequired)
	userAdminReq.Post("/ban/:id", handlers.BanUser)
	
	// Plan
	planAdminReq := authReq.Group("/admin/plan", handlers.AdminRequired)
	planAdminReq.Post("/plan/create", handlers.CreatePlan)
	planAdminReq.Post("/plan/create-features/:planId", handlers.CreateFeatures)
	planAdminReq.Put("/plan/edit/:planId", handlers.EditPlan)
	planAdminReq.Put("/plan/edit-feature/:featureId", handlers.EditFeature)
	planAdminReq.Delete("/plan/delete-plan/:planId", handlers.DeletePlan) // todo:image must deleted
	planAdminReq.Delete("/plan/delete-feature/:featureId", handlers.DeleteFeature)
	
	
	// member
	memberAdminReq := authReq.Group("/admin/member", handlers.AdminRequired)
	memberAdminReq.Post("/create/:userId", handlers.CreateMember)
	memberAdminReq.Put("/edit/:memberId", handlers.EditMember)
	memberAdminReq.Delete("/delete/:memberId", handlers.DeleteMember)

	// settings
	settingsAdminReq := authReq.Group("/admin/settings", handlers.AdminRequired)
	settingsAdminReq.Post("/create-site-feature", handlers.CreateSiteFeature)


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