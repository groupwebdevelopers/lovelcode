package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func ApiOnly(c *fiber.Ctx) error{
	api, ok := c.GetReqHeaders()["api"]
	if api=="true" && ok==true{
		return c.Next()
	}
	return c.Render("index", fiber.Map{})
}