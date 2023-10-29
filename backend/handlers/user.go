package handlers

import (
	"time"	
	"crypto/sha256"
	"encoding/base64"
	

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"lovelcode/models"
	"lovelcode/database"
	"lovelcode/utils"
	utilstoken "lovelcode/utils/token"
)


func Signin(c *fiber.Ctx) error{

	sess, err := globalSession.Get(c)
	if err!=nil{
		return utils.ServerError(c, err)
	}
	defer sess.Save()
	
	// check json and extract data from it
	var ss models.SigninUser
	if err:= c.BodyParser(&ss); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"}, err.Error())
	}


	// check SignupUser fields	
	if err:= ss.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}

	// hash password
	ss.Password = hash(ss.Password)
	
	// check user already exist
	var user models.User
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
	token = hash(token)
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
	var ss models.SignupUser
	if err:= c.BodyParser(&ss); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check SignupUser fields	
	if err:= ss.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
	// hash password
	ss.Password = hash(ss.Password)
	
	// check user already exist
	var user models.User
	query := models.User{Email: ss.Email}
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
	token = hash(token)
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

// POST, Auth Required, Admin Required, /:id
func BanUser(c *fiber.Ctx) error{
	// get id from url
	id := utils.GetIDFromParams(c, "id")
	if id ==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}

	// user have permisoin and should ban target user

		var targetUser models.User
		if err:= database.DB.First(&targetUser, &models.User{ID: id}).Error;err!=nil{
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
	var users []models.User
	if err:=database.DB.Limit(10).Offset(int(page*10)).Find(&users).Error;err!=nil{
		if err == gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"not found"})
		}else{
			return utils.ServerError(c, err)
		}
	}

	return utils.JSONResponse(c, 200, fiber.Map{"users":users})
}


func hash(s string) string{
	h:= sha256.New()
	h.Write([]byte(s))
	return string(base64.URLEncoding.EncodeToString(h.Sum(nil)))
}

func hban(userID uint64) {
	err := database.DB.Model(&models.User{}).Where("id=?", userID).Update("isBanned", true).Error
	utils.LogError(err)
}