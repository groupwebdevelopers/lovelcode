package utils

import(
	"github.com/gofiber/fiber/v2"

	"log"
)

func ServerError(c *fiber.Ctx, err error) error{
	log.Println("status:", code, ", URL:", c.OriginalURL(), "\nJSON:", mp)
	log.Println(err)
	return c.Status(500).JSON(fiber.Map{"error": "server error\ntry later"})
}

func JSONResponse(c *fiber.Ctx, status int, mp fiber.Map) error{
	// log.Println("status:", code, ", URL:", c.OriginalURL(), "\nJSON:", mp)
	return c.Status(status).JSON(mp)
}