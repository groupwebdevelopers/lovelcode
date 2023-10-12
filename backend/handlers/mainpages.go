package handlers

import (
	// "time"
	
	"github.com/gofiber/fiber/v2"
	// "gorm.io/gorm"
	
	// "lovelcode/utils"
	// "lovelcode/database"
	// "lovelcode/models"
)

// GET
func Home(c *fiber.Ctx) error{
	return c.JSON(fiber.Map{"msg":"hello i'm working"})
}
