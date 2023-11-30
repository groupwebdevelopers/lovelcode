package user

import (
	"time"	

	

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	umodels "lovelcode/models/user"
	amodels "lovelcode/models/article"
	pmodels "lovelcode/models/plan"
	"lovelcode/database"
	"lovelcode/utils"
	utilstoken "lovelcode/utils/token"
	"lovelcode/session"
)


func Signin(c *fiber.Ctx) error{

	sess, err := session.GlobalSession.Get(c)
	if err!=nil{
		return utils.ServerError(c, err)
	}
	defer sess.Save()
	
	// check json and extract data from it
	var ss umodels.SigninUser
	if err:= c.BodyParser(&ss); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"}, err.Error())
	}


	// check SignupUser fields	
	if err:= ss.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}

	// hash password
	ss.Password = utils.Hash(ss.Password)
	
	// check user already exist
	var user umodels.User
	if err:= database.DB.First(&user, "email=? and password=?", ss.Email, ss.Password).Error; err==gorm.ErrRecordNotFound{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"user not found"})
	}else if err!=nil{
		return utils.ServerError(c, err)
	}
	
	// check baned
	if user.IsBanned{
		return utils.JSONResponse(c, 403, fiber.Map{"error":"you are banned!"})
	}
	
	// create token
	// token, err := utilstoken.CreateJWTToken(user, tokenExpHours)
	// if err!=nil{
		// return utils.ServerError(c, err)
	// }
	token := utilstoken.CreateRandomToken()
	token = utils.Hash(token)
	user.Token = token
	user.TokenExp = time.Now().Add(time.Duration(database.Settings.TokenExpHours) * time.Hour)
	// update database token
	if err:= database.DB.Updates(&user).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	// put user info into session
	sess.Set("userID", user.ID)
	sess.Set("userName", user.Name)
	sess.Set("userFamily", user.Family)
	sess.Set("token", user.Token)
	sess.Set("adminPermisions", user.AdminPermisions)	
	sess.Set("email", user.Email)

	// set token to cookie
	c.Cookie(&fiber.Cookie{
		Name: "token",
		Value: token,
		Expires: time.Now().Add(time.Duration(database.Settings.TokenExpHours)*time.Hour),
	})
	return utils.JSONResponse(c, 200, fiber.Map{"msg": "you signin"})
				
}


func Signup(c *fiber.Ctx) error{
	
	// check json and extract data from it
	var ss umodels.SignupUser
	if err:= c.BodyParser(&ss); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check SignupUser fields	
	if err:= ss.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
	// hash password
	ss.Password = utils.Hash(ss.Password)
	
	// check user already exist
	var user umodels.User
	query := umodels.User{Email: ss.Email}
	if err:= database.DB.First(&user, &query).Error; err==nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"user already exist"})
	}else if err!=gorm.ErrRecordNotFound{
		return utils.ServerError(c, err)
	}

	// create token
	
	// token, err := utilstoken.CreateJWTToken(user, tokenExpHours)
	// if err!=nil{
		// return utils.ServerError(c, err)
	// }
	token := utilstoken.CreateRandomToken()
	token = utils.Hash(token)
	// create user
	user.Email = ss.Email
	user.Password = ss.Password
	user.Name = ss.Name
	user.Family = ss.Family
	// user.Number = ss.Number
	user.AdminPermisions = ""
	user.IsDeleted = false
	user.IsBanned = false
	user.Token = token
	user.TokenExp = time.Now().Add(time.Duration(database.Settings.TokenExpHours)*time.Hour)
	user.TimeCreated = time.Now()


	if err:= database.DB.Create(&user).Error; err!=nil{
		return utils.ServerError(c, err)
	}


	// // set token into cookie
	// c.Cookie(&fiber.Cookie{
	// 	Name: "token",
	// 	Value: token,
	// 	Expires: time.Now().Add(time.Duration(database.Settings.TokenExpHours)*time.Hour),
	// })
	return utils.JSONResponse(c, 200, fiber.Map{"msg": "user created. go to signin"})

}

func GetUserState(c *fiber.Ctx) error{
	
	user := c.Locals("user").(umodels.User)

	// get commnets count
	var commentsCount int
	if err:= database.DB.Select("count(id)").Where(amodels.Comment{UserID: user.ID}).Scan(&commentsCount).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	// todo: change email if phone number added
	// get plan order count
	var planOrderCount int
	if err:= database.DB.Select("count(id)").Where(pmodels.PlanOrder{Email: user.Email}).Scan(&planOrderCount).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	var rt string = utils.ConvertTimeToString(utils.ConvertToPersianTime(user.TimeCreated))

	return utils.JSONResponse(c, 200, fiber.Map{"timeRigistered":rt, "totallComments":commentsCount, "totallProjectDoingRequest": planOrderCount})
}



// POST, Auth Required, Admin Required, /:id
func BanUser(c *fiber.Ctx) error{
	// get id from url
	id := utils.GetIDFromParams(c, "id")
	if id ==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}

	// user have permisoin and should ban target user

		var targetUser umodels.User
		if err:= database.DB.First(&targetUser, &umodels.User{ID: id}).Error;err!=nil{
			if err== gorm.ErrRecordNotFound{
				return utils.JSONResponse(c, 404, fiber.Map{"error":"user not found"})
			}else{
				return utils.ServerError(c, err)
			}
		}
		targetUser.IsBanned = true
		if err:=database.DB.Updates(&targetUser).Error; err!=nil {
			// todo: use update
			if err==gorm.ErrRecordNotFound{
				return utils.JSONResponse(c, 404, fiber.Map{"error":"user not found"})
			}else{
				return utils.ServerError(c, err)
			}
		}

		return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly banned"})
	


}


// GET, auth required, admin required, /:page
func GetUsersPaged(c *fiber.Ctx) error{

	page := utils.GetIDFromParams(c, "id")
	var users []umodels.User
	if err:=database.DB.Limit(10).Offset(int(page*10)).Find(&users).Error;err!=nil{
		if err == gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"not found"})
		}else{
			return utils.ServerError(c, err)
		}
	}

	return utils.JSONResponse(c, 200, fiber.Map{"users":users})
}




func Hban(userID uint64) {
	err := database.DB.Model(&umodels.User{}).Where("id=?", userID).Update("isBanned", true).Error
	utils.LogError(err)
}