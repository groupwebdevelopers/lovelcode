package handlers

import (
	// "time"

	"lovelcode/database"
	"lovelcode/utils"

	"github.com/gofiber/fiber/v2"
	// "gorm.io/gorm"
	"lovelcode/models"
)

// GET
func Home(c *fiber.Ctx) error{
	return c.JSON(fiber.Map{"msg":"hello i'm working"})
}

func GetSiteFeatures(c *fiber.Ctx) error{
	return utils.JSONResponse(c, 200, fiber.Map{"data":database.Settings.SiteFeatures})
}

func CreateSiteFeature(c *fiber.Ctx) error{
	
	var st models.SettingsDB
	if err:= c.BodyParser(&st); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}
	
	// check siteFeature validation
	if err:=st.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
	if err:= database.DB.Create(&st).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 201, fiber.Map{"msg":"successfully created"})
}
