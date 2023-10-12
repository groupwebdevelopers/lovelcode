package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

)

func GetIDFromParams(c *fiber.Ctx) uint32{
	sid := c.Params("id", "")
	if sid==""{
		return 0
	}
	id, _ := strconv.Atoi(sid)
	return uint32(id)
}
// todo: add &models.User
// 0 means not access and 1 means access
// 2 means hacker trap and should report and ban user
// 3 means not found
func CheckAdminPermision(permisions string, p string) uint8{
	//  check hacker trap
	if p[2] == '1'{
		return 2
	}
	switch p{
	case "writeArticle":
		return (permisions[0]-'0')
	}

	return 3
}