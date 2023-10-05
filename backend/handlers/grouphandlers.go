package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func ApiOnly(c *fiber.Ctx) error{
	c.Accepts("text/plain", "application/json")
	c.Accepts(fiber.MIMETextPlain, fiber.MIMETextPlainCharsetUTF8)
	api, ok := c.GetReqHeaders()["api"]
	if api=="true" && ok==true{
		return c.Next()
	}
	return c.Render("index", fiber.Map{})
}