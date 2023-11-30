package session

import (
	"errors"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2"


	"lovelcode/utils"
	umodels "lovelcode/models/user"
)

var GlobalSession = session.New()


func GetUserFromSession(c *fiber.Ctx) (umodels.User, error) {
	var user umodels.User
	sess, err := GlobalSession.Get(c)
	if err!=nil{
		utils.LogError(err)
		return user, err
	}
	iAdminPermisions := sess.Get("adminPermsions")
	iUserID := sess.Get("userID")
	iUserName := sess.Get("userName")
	iUserFamily := sess.Get("userFamily")
	iToken := sess.Get("token")
	iEmail := sess.Get("email")
	
	if iAdminPermisions == nil || iUserID == nil || iUserName == nil || iUserFamily == nil || iToken == nil || iEmail == nil{
		return user, errors.New("auth required")
	}

	user.AdminPermisions = iAdminPermisions.(string)
	user.ID = iUserID.(uint64)
	user.Name = iUserName.(string)
	user.Family = iUserFamily.(string)
	user.Token = iToken.(string)
	user.Email = iEmail.(string)
 
	return user, nil

}
