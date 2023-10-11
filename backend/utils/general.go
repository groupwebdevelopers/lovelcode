package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

)

func GetIDFromParams(c *fiber.Ctx) int{
	sid := c.Params("id", "")
	if sid==""{
		return 0
	}
	id, _ := strconv.Atoi(sid)
	return id
}