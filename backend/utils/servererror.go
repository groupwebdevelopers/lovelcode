package utils

import(
	"github.com/gofiber/fiber/v2"

	"log"
)

func ServerError(c *fiber.Ctx, err error) error{
	log.Println(err)
	return c.Status(500).JSON(fiber.Map{"error": "server error\ntry later"})
}