package handlers

import (
	// "time"

	"lovelcode/database"
	"lovelcode/utils"

	"github.com/gofiber/fiber/v2"
	// "gorm.io/gorm"
)

// GET
func Home(c *fiber.Ctx) error{
	return c.JSON(fiber.Map{"msg":"hello i'm working 1234"})
}

func GetSiteFeatures(c *fiber.Ctx) error{
	return utils.JSONResponse(c, 200, fiber.Map{"data":database.Settings.SiteFeatures})
}

