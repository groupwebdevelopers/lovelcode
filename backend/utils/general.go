package utils

import (
	"strconv"
	"github.com/gofiber/fiber/v2"

)

func GetIntFromParams(c *fiber.Ctx, name string) uint64{
	sid := c.Params(name, "")
	if sid==""{
		return 0
	}
	id, _ := strconv.Atoi(sid)
	return uint64(id) // todo: id is int
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
	case "plan":
		return (permisions[10]-'0')
	}

	return 3
}