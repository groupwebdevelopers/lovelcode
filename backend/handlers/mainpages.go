package handlers

import (
	"time"
	
	"github.com/gofiber/fiber/v2"
	
	"lovelcode/utils"
	"lovelcode/database"
	"lovelcode/models"
)

// GET
func Home(c *fiber.Ctx) error{
	return c.JSON(fiber.Map{"msg":"hello i'm working"})
}


// POST, Auth Required
func CreateProjectDoingRequest(c *fiber.Ctx) error{
	type PDR struct{
		Title string `json:"title"`
		Description string `json:"description"`
		SuggestedPrice int `json:"suggestedprice`
	}

	var pdr PDR
	if err:= c.BodyParser(&pdr); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	if pdr.Title == "" || pdr.Description == ""{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"empty title or description"})
	}

	var pd models.ProjectDoingRequest
	pd.Title = pdr.Title
	pd.Description = pdr.Description
	pd.SuggestedPrice = uint(pdr.SuggestedPrice)
	pd.TimeCreated = time.Now()
	pd.TimeModified = time.Now()
	pd.User = c.Locals("user").(models.User)

	if err:= database.DB.Create(&pd).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 201, fiber.Map{"msg": "request saved"})
}