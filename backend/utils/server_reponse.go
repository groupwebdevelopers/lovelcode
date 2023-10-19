package utils

import(
	"github.com/gofiber/fiber/v2"

	"log"
)

func ServerError(c *fiber.Ctx, err error, errmsg ...string) error{
	log.Println("###################################")
	log.Println("status:", 500, ", URL:", c.OriginalURL())
	log.Println(err)
	log.Println(errmsg)
	log.Println("###################################")
	return c.JSON(fiber.Map{"status":500, "error": "server error\ntry later"})
}

func JSONResponse(c *fiber.Ctx, status int, mp fiber.Map, lg... string) error{
	log.Println("----------")
	log.Println("status:", status, ", URL:", c.OriginalURL(), "\nJSON:", mp, "\n", lg)
	log.Println("----------")
	mp["status"] = status
	return c.JSON(mp)
}