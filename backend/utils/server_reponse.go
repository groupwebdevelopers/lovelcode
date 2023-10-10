package utils

import(
	"github.com/gofiber/fiber/v2"

	"log"
)

func ServerError(c *fiber.Ctx, err error, errmsg ...string) error{
	log.Println("###################################")
	log.Println("status:", 500, ", URL:", c.OriginalURL())
	log.Println(err)
	log.Println(errmsg[0])
	log.Println("###################################")
	return c.Status(500).JSON(fiber.Map{"error": "server error\ntry later"})
}

func JSONResponse(c *fiber.Ctx, status int, mp fiber.Map) error{
	log.Println("----------")
	log.Println("status:", status, ", URL:", c.OriginalURL(), "\nJSON:", mp)
	log.Println("----------")
	return c.Status(status).JSON(mp)
}