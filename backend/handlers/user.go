package handlers

import (
	"time"	
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"lovelcode/models"
	"lovelcode/database"
	"lovelcode/utils"
	utilstoken "lovelcode/utils/token"
)


func Signin(c *fiber.Ctx) error{
	

	// tokenExpHours, errr := getTokenExpHours()
	// tokenExpHours := 72
	var tokenExpHours uint16= 72


	// if errr!=nil{
		// return utils.ServerError(c, errr)
	// }

	// check json and extract data from it

	var ss models.SigninUser
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
	if err:= database.DB.First(&user, "email=? and password=?", ss.Email, ss.Password).Error; err==gorm.ErrRecordNotFound{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"user not found"})
	}else if err!=nil{
		return utils.ServerError(c, err)
	}
	
	// check baned
	if user.IsBanned == true{
		return utils.JSONResponse(c, 403, fiber.Map{"error":"you are banned!"})
	}
	
	// create token
	// token, err := utilstoken.CreateJWTToken(user, tokenExpHours)
	// if err!=nil{
		// return utils.ServerError(c, err)
	// }
	token := utilstoken.CreateToken()
	user.Token = token
	user.TokenExp = time.Now().Add(time.Duration(tokenExpHours) * time.Hour)
	// update database token
	if err:= database.DB.Updates(&user).Error; err!=nil{
		return utils.ServerError(c, err)
	}
		
	// set token to cookie
	c.Cookie(&fiber.Cookie{
		Name: "token",
		Value: token,
		Expires: time.Now().Add(time.Duration(tokenExpHours)*time.Hour),
	})
	return utils.JSONResponse(c, 200, fiber.Map{"msg": "you signin"})
				
}


func Signup(c *fiber.Ctx) error{
	

	// tokenExpHours, errr := getTokenExpHours()
	var tokenExpHours uint16= 72
	// if errr!=nil{
		// return utils.ServerError(c, errr)
	// }

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
	token := utilstoken.CreateToken()
		
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
	user.TokenExp = time.Now().Add(time.Duration(tokenExpHours)*time.Hour)

	if err:= database.DB.Create(&user).Error; err!=nil{
		return utils.ServerError(c, err)
	}


	// set token into cookie
	c.Cookie(&fiber.Cookie{
		Name: "token",
		Value: token,
		Expires: time.Now().Add(time.Duration(tokenExpHours)*time.Hour),
	})
	return utils.JSONResponse(c, 200, fiber.Map{"msg": "user created"})

}

// POST, Auth Required, Admin Required, /:id
func BanUser(c *fiber.Ctx) error{
	// get id from url
	id := utils.GetIntFromParams(c, "id")
	user := c.Locals("user").(models.User)
	if id ==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}

	// check user have permision for ban user
	ap := utils.CheckAdminPermision(user.AdminPermisions, "banuser")
	if ap == 0 || ap == 2{
		if ap == 2{
			hban(user)
		}
		return utils.JSONResponse(c, 403, fiber.Map{"error":"you don't access to do this"})
	}
	// user have permisoin and should ban target user
	if ap == 1{
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
		utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly banned"})
	}
	if ap == 3{
		return utils.ServerError(c, errors.New("the sent permision to function not found"))
	}

	return utils.ServerError(c, errors.New("can't do this (ap="+string(ap)+")"))
}


// GET, auth required, admin required, /:page
func GetUsersPaged(c *fiber.Ctx) error{
	user:= c.Locals("user").(models.User)
	adminCode := utils.CheckAdminPermision(user.AdminPermisions, "seeUsers")
	if adminCode != 1{
		if adminCode == 2{
			hban(user)
		}
		return utils.JSONResponse(c, 403, fiber.Map{"error":"Access Denied"})
	}

	page := utils.GetIntFromParams(c, "id")
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

func getTokenExpHours() (uint16, error){
	if i, err := strconv.Atoi(database.Settings["tokenExpHours"]);err!=nil{
		return 0, errors.New(err.Error() + "\nin database tokenExpHours setted to invalid intiger")
	}else{
		if i > int(^uint16(0)){
			return 0, errors.New("in database tokenExpHours value is too big")
		}
		return uint16(i), nil
	}
}

func hban(user models.User) {
	user.IsBanned = true
	database.DB.Updates(&user)
}