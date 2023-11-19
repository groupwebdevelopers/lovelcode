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
	fileUploadAdminReq.Post("/blog/image/:articleId", handlers.UploadArticleImage)
	fileUploadAdminReq.Post("/customer/image/:customerId", handlers.UploadCustomerImage)
	fileUploadAdminReq.Post("/portfolio/image/:portfolioId", handlers.UploadPortfolioImage)
	
	// auth not required
	
	// auth
	apiV1.Post("/signin", handlers.Signin)
	apiV1.Post("/signup", handlers.Signup)
	
	// plan
	apiV1.Get("/plan/get-all", handlers.GetAllPlansAndFeatures)
	apiV1.Get("/plan/get-featured", handlers.GetFeaturedPlans)
	apiV1.Get("/plan/get-all-plan-types", handlers.GetAllPlanTypes)
	apiV1.Get("/plan/get-plan-type/:planTypeId", handlers.GetPlanType)
	
	// member
	apiV1.Get("/member/get-all", handlers.GetAllMembers)
	
	// article
	apiV1.Get("/blog/get/:articleTitleUrl", handlers.GetArticle)
	apiV1.Get("/blog/get-all", handlers.GetAllArticlesTitles)
	apiV1.Get("/blog/get-featured", handlers.GetFeaturedArticlesTitle)
	apiV1.Get("/blog/get-categories", handlers.GetAllArticleCategories)
	apiV1.Get("/blog/search", handlers.SearchArticle)

	// portfolio
	apiV1.Get("/portfolio/get-all", handlers.GetAllPortfolios)
	apiV1.Get("/portfolio/get-featured", handlers.GetFeaturedPortfolios)
	
	// site info
	apiV1.Get("/site-social-media", handlers.GetSiteSocialMedia)
	apiV1.Get("/site-phone-numbers", handlers.GetSitePhoneNumbers)
	
	// comment
	apiV1.Get("/comment/get-all-for-article/:articleTitleUrl", handlers.GetAllArticleComments)
	
	// contact us
	apiV1.Post("/contactus/create", handlers.CreateContactUs)
	
	// customer
	apiV1.Get("/customer/get-all", handlers.GetAllCustomers)
	apiV1.Get("/customer/get-featured", handlers.GetFeaturedCustomers)
	
	// main pages
	apiV1.Get("/mainpage/:pageName", handlers.GetMainPage)
	
	// statistic
	apiV1.Get("/statistic/get-public", handlers.GetPublicStatistics)

	// order plan
	apiV1.Post("/order-plan/create", handlers.CreateOrderPlan)



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
	commentAuthReq.Post("/create/:articleTitleUrl", handlers.CreateComment)
	commentAuthReq.Put("/edit/:commentId", handlers.EditComment)
	commentAuthReq.Delete("/delete/:commentId", handlers.DeleteComment)
	
	// user
	userAuthReq := apiV1.Group("/user", handlers.AuthRequired)
	userAuthReq.Get("/get-state", handlers.GetUserState)

	// order plan
	orderPlanAuthReq := apiV1.Group("/order-plan", handlers.AuthRequired)
	orderPlanAuthReq.Put("/edit/:orderPlanId", handlers.EditOrderPlan)
	orderPlanAuthReq.Delete("/delete/:orderPlanId", handlers.DeleteOrderPlan)
	orderPlanAuthReq.Get("/get-all-user/", handlers.GetAllUserOrderPlans)

	
	// admin required
	
	// article
	adminArticleReq := apiV1.Group("/admin/blog", handlers.AdminArticleRequired)
	adminArticleReq.Post("/create", handlers.CreateArticle)
	adminArticleReq.Put("/edit/:articleId", handlers.EditArticle)
	adminArticleReq.Delete("/delete/:articleId", handlers.DeleteArticle)
	
	// article category
	adminArticleCategoryReq := apiV1.Group("/admin/article-category", handlers.AdminRequired)
	adminArticleCategoryReq.Post("/create", handlers.CreateArticleCategory)
	adminArticleCategoryReq.Put("/edit/:articleCategoryId", handlers.EditArticleCategory)
	adminArticleCategoryReq.Delete("/delete/:articleCategoryId", handlers.DeleteArticleCategory)

	
	// adminReq := authReq.Group("/admin", handlers.AdminRequired)
	
	// user
	userAdminReq := apiV1.Group("/admin/user", handlers.AdminRequired)
	userAdminReq.Post("/ban/:id", handlers.BanUser)
	
	// Plan
	planAdminReq := apiV1.Group("/admin/plan", handlers.AdminRequired)
	planAdminReq.Post("/create", handlers.CreatePlan)
	planAdminReq.Put("/edit/:planId", handlers.EditPlan)
	planAdminReq.Delete("/delete-plan/:planId", handlers.DeletePlan) // todo:image must deleted
	planAdminReq.Get("/get-all-plans", handlers.GetAllPlans)
	planAdminReq.Get("/get-plan/:planId", handlers.GetPlan)
	// feature
	planAdminReq.Post("/create-features/:planId", handlers.CreateFeatures)
	planAdminReq.Put("/edit-feature/:featureId", handlers.EditFeature)
	planAdminReq.Delete("/delete-feature/:featureId", handlers.DeleteFeature)
	planAdminReq.Get("/get-all-features/:planId", handlers.GetAllFeatures)
	planAdminReq.Get("/get-feature/:featureId", handlers.GetFeature)
	// plan type
	planAdminReq.Post("/create-plan-type", handlers.CreatePlanType)
	planAdminReq.Put("/edit-plan-type/:planTypeId", handlers.EditPlanType)
	planAdminReq.Delete("/delete-plan-type/:planTypeId", handlers.DeletePlanType)
	
	
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
	
	// portfolio
	adminPortfolioReq := apiV1.Group("/admin/portfolio", handlers.AdminRequired)
	adminPortfolioReq.Post("/create", handlers.CreatePortfolio)
	adminPortfolioReq.Put("/edit/:portfolioId", handlers.EditPortfolio)
	adminPortfolioReq.Delete("/delete/:portfolioId", handlers.DeletePortfolio)
	adminPortfolioReq.Get("/get/:portfolioId", handlers.GetPortfolio)
	
	
	// contactus
	contactusAdminReq := apiV1.Group("/admin/contactus", handlers.AdminRequired)
	
	// contactus
	contactusAuthReq := apiV1.Group("/admin/contactus", handlers.AdminRequired)
	contactusAuthReq.Get("/get/:contactusId", handlers.GetContactUs)
	contactusAuthReq.Delete("/delete/:contactusId", handlers.DeleteContactUs)
	contactusAdminReq.Get("/get-all/", handlers.GetAllContactUss) // with query
	
	// member
	statisticAdminReq := apiV1.Group("/admin/statistic", handlers.AdminRequired)
	statisticAdminReq.Post("/create/:statisticId", handlers.CreateStatistic)
	statisticAdminReq.Put("/edit/:statisticId", handlers.EditStatistic)
	statisticAdminReq.Delete("/delete/:statisticId", handlers.DeleteStatistic)
	statisticAdminReq.Get("/get-all/", handlers.GetAllStatistics)
	
	// mainpage
	mainpageAdminReq := apiV1.Group("/admin/mainpage", handlers.AdminRequired)
	mainpageAdminReq.Post("/create/:mainpageTextId", handlers.CreateMainpageTexts)
	mainpageAdminReq.Put("/edit/:mainpageTextId", handlers.EditMainpageText)
	mainpageAdminReq.Delete("/delete/:mainpageTextId", handlers.DeleteMainpageText)
	mainpageAdminReq.Get("/get-all/", handlers.GetAllMainpageText)

	// order plan
	orderPlanAdminReq := apiV1.Group("/admin/order-plan", handlers.AdminRequired)
	orderPlanAdminReq.Get("/get-all-order-plans", handlers.GetAllOrderPlans)
	


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