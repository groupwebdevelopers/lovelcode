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
	fileUploadAdminReq.Post("/customer/image/:customerId", handlers.UploadCustomerImage)
	fileUploadAdminReq.Post("/work-sample/image/:workSampleId", handlers.UploadWorkSampleImage)
	
	// auth not required
	
	// auth
	apiV1.Post("/signin", handlers.Signin)
	apiV1.Post("/signup", handlers.Signup)
	
	// plan
	apiV1.Get("/plan/get-all-plans-and-features", handlers.GetAllPlansAndFeatures)
	apiV1.Get("/plan/get-featured", handlers.GetFeaturedPlans)
	
	// member
	apiV1.Get("/member/get-all", handlers.GetAllMembers)
	
	// article
	apiV1.Get("/article/get/:articleTitleUrl", handlers.GetArticle)
	apiV1.Get("/article/get-all/:page", handlers.GetAllArticlesTitles)
	apiV1.Get("/article/get-featured", handlers.GetFeaturedArticlesTitle)
	
	
	// work sample
	apiV1.Get("/work-sample/get-all/:page", handlers.GetAllWorkSamples)
	apiV1.Get("/work-sample/get-featured", handlers.GetFeaturedWorkSamples)
	
	// get features
	apiV1.Get("/site-features/get-all", handlers.GetSiteFeatures)
	
	// comment
	apiV1.Get("/comment/get-all-for-article/:articleTitleUrl", handlers.GetAllArticleComments)
	
	// contact us
	apiV1.Get("/contactus/get-all-for-user", handlers.GetAllUserContactUss)

	// customer
	apiV1.Get("/customer/get-all", handlers.GetAllCustomers)
	apiV1.Get("/customer/get-featured", handlers.GetFeaturedCustomers)
	
	// main pages
	apiV1.Get("/mainpage/:pageName", handlers.GetMainPage)

	// auth required
	
	// Project Doing Request
	pdrAuthReq := apiV1.Group("/pdr", handlers.AuthRequired)
	pdrAuthReq.Post("/create", handlers.CreateProjectDoingRequest)
	pdrAuthReq.Get("/get/:id", handlers.GetProjectDoingRequest)
	pdrAuthReq.Get("/get-all", handlers.GetAllProjectDoingRequests)
	pdrAuthReq.Put("/edit/:id", handlers.EditProjectDoingRequest)
	pdrAuthReq.Delete("/delete/:id", handlers.DeleteProjectDoingRequest)
	
	// comment
	commentAuthReq := apiV1.Group("/comment", handlers.AuthRequired)
	commentAuthReq.Post("/create", handlers.CreateComment)
	commentAuthReq.Put("/edit/:id", handlers.EditComment)
	commentAuthReq.Delete("/delete/:id", handlers.DeleteComment)
	
	// contactus
	contactusAuthReq := apiV1.Group("/contactus", handlers.AuthRequired)
	contactusAuthReq.Post("/create", handlers.CreateContactUs)
	contactusAuthReq.Get("/get/:contactUsTitle", handlers.GetContactUsByTitle)
	contactusAuthReq.Put("/edit/:id", handlers.EditContactUs)
	contactusAuthReq.Delete("/delete/:id", handlers.DeleteContactUs)
	
	
	// admin required
	
	// article
	adminArticleReq := apiV1.Group("/admin/article", handlers.AdminArticleRequired)
	adminArticleReq.Post("/create", handlers.CreateArticle)
	adminArticleReq.Put("/edit/:articleId", handlers.EditArticle)
	adminArticleReq.Delete("/delete/:articleId", handlers.DeleteArticle)
	
	
	// adminReq := authReq.Group("/admin", handlers.AdminRequired)
	
	// user
	userAdminReq := apiV1.Group("/admin/user", handlers.AdminRequired)
	userAdminReq.Post("/ban/:id", handlers.BanUser)
	
	// Plan
	planAdminReq := apiV1.Group("/admin/plan", handlers.AdminRequired)
	planAdminReq.Post("/create", handlers.CreatePlan)
	planAdminReq.Post("/create-features/:planId", handlers.CreateFeatures)
	planAdminReq.Put("/edit/:planId", handlers.EditPlan)
	planAdminReq.Put("/edit-feature/:featureId", handlers.EditFeature)
	planAdminReq.Delete("/delete-plan/:planId", handlers.DeletePlan) // todo:image must deleted
	planAdminReq.Delete("/delete-feature/:featureId", handlers.DeleteFeature)
	planAdminReq.Get("/get-all-plans", handlers.GetAllPlans)
	planAdminReq.Get("/get-plan/:planId", handlers.GetPlan)
	planAdminReq.Get("/get-all-features/:planId", handlers.GetAllFeatures)
	planAdminReq.Get("/get-feature/:featureId", handlers.GetFeature)
	
	
	// member
	memberAdminReq := apiV1.Group("/admin/member", handlers.AdminRequired)
	memberAdminReq.Post("/create/:userId", handlers.CreateMember)
	memberAdminReq.Put("/edit/:memberId", handlers.EditMember)
	memberAdminReq.Delete("/delete/:memberId", handlers.DeleteMember)
	memberAdminReq.Get("/get/:memberId", handlers.GetMember)
	
	// settings
	settingsAdminReq := apiV1.Group("/admin/settings", handlers.AdminRequired)
	settingsAdminReq.Post("/create", handlers.CreateSetting)
	settingsAdminReq.Put("/edit/:settingId", handlers.EditSetting)
	settingsAdminReq.Delete("/delete/:settingId", handlers.DeleteSetting)
	settingsAdminReq.Get("/get-all", handlers.GetAllSettings)
	
	// customer
	customerAdminReq := apiV1.Group("/admin/customer", handlers.AdminRequired)
	customerAdminReq.Post("/create", handlers.CreateCustomer)
	customerAdminReq.Put("/edit/:customerId", handlers.EditCustomer)
	customerAdminReq.Delete("/delete/:customerId", handlers.DeleteCustomer)
	customerAdminReq.Get("/get/:customerId", handlers.GetCustomer)
	
	// work sample
	adminWorkSampleReq := apiV1.Group("/admin/work-sample", handlers.AdminRequired)
	adminWorkSampleReq.Post("/create", handlers.CreateWorkSample)
	adminWorkSampleReq.Put("/edit/:workSampleId", handlers.EditWorkSample)
	adminWorkSampleReq.Delete("/delete/:workSmapleId", handlers.DeleteWorkSample)
	adminWorkSampleReq.Get("/get/:workSampleId", handlers.GetWorkSample)
	

	// contactus
	contactusAdminReq := apiV1.Group("/admin/contactus", handlers.AdminRequired)
	contactusAdminReq.Get("/get-all", handlers.GetAllUserContactUss)
	//todo: anwser


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