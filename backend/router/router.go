package router

import (
	"github.com/gofiber/fiber/v2"
	ahandlers "lovelcode/handlers/article"
	cohandlers "lovelcode/handlers/contactus"
	cuhandlers "lovelcode/handlers/customer"
	ghandlers "lovelcode/handlers/grouphandlers"
	mhandlers "lovelcode/handlers/mainpage"
	phandlers "lovelcode/handlers/plan"
	pfhandlers "lovelcode/handlers/portfolio"
	shandlers "lovelcode/handlers/settings"
	sshandlers "lovelcode/handlers/statistics"
	uhandlers "lovelcode/handlers/user"
	"lovelcode/utils"
)
//todo: image file size limit

func Route(app *fiber.App) {
	apiOnly := app.Group("/api", ghandlers.ApiOnly)
	apiV1 := apiOnly.Group("/v1")
	apiV1.Get("/", mhandlers.Home)
	
	// fileUploadAuthReq := app.Group("/api/v1/upload", handlers.AuthRequired)
	fileUploadAdminReq := app.Group("/admin/upload", ghandlers.AdminUploadImage)
	fileUploadAdminReq.Post("/plan/image/:planId", phandlers.UploadPlanImage)
	fileUploadAdminReq.Post("/member/image/:memberId", uhandlers.UploadMemberImage)
	fileUploadAdminReq.Post("/blog/image/:articleTitleUrl", ahandlers.UploadArticleImage)
	fileUploadAdminReq.Post("/customer/image/:customerId", cuhandlers.UploadCustomerImage)
	fileUploadAdminReq.Post("/portfolio/image/:portfolioId", pfhandlers.UploadPortfolioImage)
	
	// auth not required
	
	// auth
	apiV1.Post("/signin", uhandlers.Signin)
	apiV1.Post("/signup", uhandlers.Signup)
	
	// plan
	apiV1.Get("/plan/get-all", phandlers.GetAllPlansAndFeatures)
	apiV1.Get("/plan/get-featured", phandlers.GetFeaturedPlans)
	apiV1.Get("/plan/get-all-plan-types", phandlers.GetAllPlanTypes)
	apiV1.Get("/plan/get-plan-type/:planTypeId", phandlers.GetPlanType)
	
	// member
	apiV1.Get("/member/get-all", uhandlers.GetAllMembers)
	
	// article
	apiV1.Get("/blog/get/:articleTitleUrl", ahandlers.GetArticle)
	apiV1.Get("/blog/get-all", ahandlers.GetAllArticlesTitles)
	apiV1.Get("/blog/get-featured", ahandlers.GetFeaturedArticlesTitle)
	apiV1.Get("/blog/get-categories", ahandlers.GetAllArticleCategories)
	apiV1.Get("/blog/search", ahandlers.SearchArticle)

	// portfolio
	apiV1.Get("/portfolio/get-all", pfhandlers.GetAllPortfolios)
	apiV1.Get("/portfolio/get-featured", pfhandlers.GetFeaturedPortfolios)
	
	// site info
	apiV1.Get("/site-social-media", mhandlers.GetSiteSocialMedia)
	apiV1.Get("/site-phone-numbers", mhandlers.GetSitePhoneNumbers)
	
	// comment
	apiV1.Get("/comment/get-all-for-article/:articleTitleUrl", ahandlers.GetAllArticleComments)
	
	// contact us
	apiV1.Post("/contactus/create", cohandlers.CreateContactUs)
	
	// customer
	apiV1.Get("/customer/get-all", cuhandlers.GetAllCustomers)
	apiV1.Get("/customer/get-featured", cuhandlers.GetFeaturedCustomers)
	
	// main pages
	apiV1.Get("/mainpage/:pageName", mhandlers.GetMainPage)
	
	// statistic
	apiV1.Get("/statistic/get-public", sshandlers.GetPublicStatistics)

	// plan order
	apiV1.Post("/order-plan/create", phandlers.CreatePlanOrder)



	// auth required
	
	
	// comment
	commentAuthReq := apiV1.Group("/blog-comment", ghandlers.AuthRequired)
	commentAuthReq.Post("/create/:articleTitleUrl", ahandlers.CreateComment)
	commentAuthReq.Put("/edit/:commentId", ahandlers.EditComment)
	commentAuthReq.Delete("/delete/:commentId", ahandlers.DeleteComment)
	
	// user
	userAuthReq := apiV1.Group("/user", ghandlers.AuthRequired)
	userAuthReq.Get("/get-state", uhandlers.GetUserState)

	// plan order
	PlanOrderAuthReq := apiV1.Group("/order-plan", ghandlers.AuthRequired)
	PlanOrderAuthReq.Put("/edit/:PlanOrderId", phandlers.EditPlanOrder)
	PlanOrderAuthReq.Delete("/delete/:PlanOrderId", phandlers.DeletePlanOrder)
	PlanOrderAuthReq.Get("/get-all-user/", phandlers.GetAllUserPlanOrders)

	
	// admin required
	
	// article
	adminArticleReq := apiV1.Group("/admin/blog", ghandlers.AdminArticleRequired)
	adminArticleReq.Post("/create/:articleTitleUrl", ahandlers.CreateArticle)
	adminArticleReq.Put("/edit/:articleTitleUrl", ahandlers.EditArticle)
	adminArticleReq.Delete("/delete/:articleTitleUrl", ahandlers.DeleteArticle)
	
	// article category
	adminArticleCategoryReq := apiV1.Group("/admin/article-category", ghandlers.AdminRequired)
	adminArticleCategoryReq.Post("/create", ahandlers.CreateArticleCategory)
	adminArticleCategoryReq.Put("/edit/:articleCategoryId", ahandlers.EditArticleCategory)
	adminArticleCategoryReq.Delete("/delete/:articleCategoryId", ahandlers.DeleteArticleCategory)

	
	// adminReq := authReq.Group("/admin", handlers.AdminRequired)
	
	// user
	userAdminReq := apiV1.Group("/admin/user", ghandlers.AdminRequired)
	userAdminReq.Post("/ban/:id", uhandlers.BanUser)
	
	// Plan
	planAdminReq := apiV1.Group("/admin/plan", ghandlers.AdminRequired)
	planAdminReq.Post("/create", phandlers.CreatePlan)
	planAdminReq.Put("/edit/:planId", phandlers.EditPlan)
	planAdminReq.Delete("/delete-plan/:planId", phandlers.DeletePlan) // todo:image must deleted
	planAdminReq.Get("/get-all-plans", phandlers.GetAllPlans)
	planAdminReq.Get("/get-plan/:planId", phandlers.GetPlan)
	// feature
	planAdminReq.Post("/create-features/:planId", phandlers.CreateFeatures)
	planAdminReq.Put("/edit-feature/:featureId", phandlers.EditFeature)
	planAdminReq.Delete("/delete-feature/:featureId", phandlers.DeleteFeature)
	planAdminReq.Get("/get-all-features/:planId", phandlers.GetAllFeatures)
	planAdminReq.Get("/get-feature/:featureId", phandlers.GetFeature)
	// plan type
	planAdminReq.Post("/create-plan-type", phandlers.CreatePlanType)
	planAdminReq.Put("/edit-plan-type/:planTypeId", phandlers.EditPlanType)
	planAdminReq.Delete("/delete-plan-type/:planTypeId", phandlers.DeletePlanType)
	
	
	// member
	memberAdminReq := apiV1.Group("/admin/member", ghandlers.AdminRequired)
	memberAdminReq.Post("/create/:userId", uhandlers.CreateMember)
	memberAdminReq.Put("/edit/:memberId", uhandlers.EditMember)
	memberAdminReq.Delete("/delete/:memberId", uhandlers.DeleteMember)
	memberAdminReq.Get("/get/:memberId", uhandlers.GetMember)
	
	// settings
	settingsAdminReq := apiV1.Group("/admin/settings", ghandlers.AdminRequired)
	settingsAdminReq.Post("/create", shandlers.CreateSetting)
	settingsAdminReq.Put("/edit/:settingId", shandlers.EditSetting)
	settingsAdminReq.Delete("/delete/:settingId", shandlers.DeleteSetting)
	settingsAdminReq.Get("/get-all", shandlers.GetAllSettings)
	
	// customer
	customerAdminReq := apiV1.Group("/admin/customer", ghandlers.AdminRequired)
	customerAdminReq.Post("/create", cuhandlers.CreateCustomer)
	customerAdminReq.Put("/edit/:customerId", cuhandlers.EditCustomer)
	customerAdminReq.Delete("/delete/:customerId", cuhandlers.DeleteCustomer)
	customerAdminReq.Get("/get/:customerId", cuhandlers.GetCustomer)
	
	// portfolio
	adminPortfolioReq := apiV1.Group("/admin/portfolio", ghandlers.AdminRequired)
	adminPortfolioReq.Post("/create", pfhandlers.CreatePortfolio)
	adminPortfolioReq.Put("/edit/:portfolioId", pfhandlers.EditPortfolio)
	adminPortfolioReq.Delete("/delete/:portfolioId", pfhandlers.DeletePortfolio)
	adminPortfolioReq.Get("/get/:portfolioId", pfhandlers.GetPortfolio)
	
	
	// contactus
	// contactusAuthReq := apiV1.Group("/admin/contactus", ghandlers.AuthRequired)
	
	// contactus
	contactusAdminReq := apiV1.Group("/admin/contactus", ghandlers.AdminRequired)
	contactusAdminReq.Get("/get/:contactusId", cohandlers.GetContactUs)
	contactusAdminReq.Delete("/delete/:contactusId", cohandlers.DeleteContactUs)
	contactusAdminReq.Get("/get-all/", cohandlers.GetAllContactUss) // with query
	contactusAdminReq.Get("/get-all-unseen/", cohandlers.GetAllUnseenContactUss) // with query
	contactusAdminReq.Post("/save-as-seen/:contactusId", cohandlers.SaveAsSeen) // with query
	contactusAdminReq.Post("/save-as-unseen/:contactusId", cohandlers.SaveAsUnSeen) // with query
	

	// member
	statisticAdminReq := apiV1.Group("/admin/statistic", ghandlers.AdminRequired)
	statisticAdminReq.Post("/create/:statisticId", sshandlers.CreateStatistic)
	statisticAdminReq.Put("/edit/:statisticId", sshandlers.EditStatistic)
	statisticAdminReq.Delete("/delete/:statisticId", sshandlers.DeleteStatistic)
	statisticAdminReq.Get("/get-all/", sshandlers.GetAllStatistics)
	
	// mainpage
	mainpageAdminReq := apiV1.Group("/admin/mainpage", ghandlers.AdminRequired)
	mainpageAdminReq.Post("/create/:mainpageTextId", mhandlers.CreateMainpageTexts)
	mainpageAdminReq.Put("/edit/:mainpageTextId", mhandlers.EditMainpageText)
	mainpageAdminReq.Delete("/delete/:mainpageTextId", mhandlers.DeleteMainpageText)
	mainpageAdminReq.Get("/get-all/", mhandlers.GetAllMainpageText)

	// plan order
	PlanOrderAdminReq := apiV1.Group("/admin/order-plan", ghandlers.AdminRequired)
	PlanOrderAdminReq.Get("/get-all", phandlers.GetAllPlanOrders)
	
	// todo: must remove in production
	apiV1.Get("/routes", func (c *fiber.Ctx)error{return utils.JSONResponse(c, 200, fiber.Map{"routes":app.GetRoutes()})})

	// file upload admin required
	// apt not found
	apiOnly.Use(func (c *fiber.Ctx) error{
		return utils.JSONResponse(c, 404, fiber.Map{"error":"API not found.", "apis": app.GetRoutes()})
	})
	
	// Static
	app.Static("/", "../frontend/dist")
	app.Static("/", "/frontend/dist")
	app.Static("*", "../frontend/dist/index.html")
	app.Static("*", "/frontend/dist/index.html")

	// todo:add robots.txt

	// static not found
	app.Use(func (c *fiber.Ctx) error{
		return utils.JSONResponse(c, 404, fiber.Map{"error":"Page not found."})
	})
	

}