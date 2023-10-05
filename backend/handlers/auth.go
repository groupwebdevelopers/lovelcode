package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Signin(c *fiber.Ctx) error{
	return c.JSON(fiber.Map{"msg":"not ready"})
}


func Signup(c *fiber.Ctx) error{
	return c.JSON(fiber.Map{"msg":"not ready"})
}