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
)


func Signin(c *fiber.Ctx) error{
	type SigninStruct struct{
		Email string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// check json and extract data from it

	var ss SigninStruct
	if err:= c.BodyParser(&ss); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	if ss.Email != ""{
	// check email
		if err := utils.CheckEmail(ss.Email); err!=nil{
			return utils.JSONReponse(c, 400,fiber.Map{"error":"invalid email"})
		}
	}else {
		// check username
		if err:= utils.IsJustLetter(ss.Username, "-._"); err!=nil{
			return utils.JSONReponse(c, 400,fiber.Map{"error":err.Error})
		}	
	}
	
	// check password len
	if len(ss.Password) < 8{
		return utils.JSONReponse(c, 400, fiber.Map{"error": "small password (<8)"})
	}
	
	// hash password
	ss.Password = hash(ss.Password)
	
	// check user already exist
	var user models.User
	if err:= database.DB.First(&user, "email=? or username=?", ss.Email, ss.Username).Error; err==gorm.ErrRecordNotFound{
		return utils.JSONReponse(c, 400, fiber.Map{"error":"user not found"})
	}else if err!=nil{
		return utils.ServerError(c, err)
	}
	
	// check baned
	if user.IsBanned == true{
		return utils.JSONResponse(c, 403, fiber.Map{"error":"you are banned!"})
	}
	
	// create token
	token, err := utils.CreateJWTToken(user, tokenExpHours)
	if err!=nil{
		return utils.ServerError(c, err)
	}

	// update database token
	if err:= database.DB.Updates(&user); err!=nil{
		return utils.ServerError(c, err)
	}
		
	// set token to cookie
	c.Cookie(&fiber.Cookie{
		Name: "token",
		Value: token,
		Expires: time.Now().Add(time.Duration(database.Settings["tokenExpHours"])*time.Hour),
	})
	return utils.JSONResponse(c, 200, fiber.Map{"msg": "you signin"})
				
}


func Signup(c *fiber.Ctx) error{
	type SignupStruct struct{
		Email string `json:"email"`
		Password string `json:"password"`
		Name string `json:"name"`
		Family string `json:"family"`
		Username string `json:"username"`
	}

	// check json and extract data from it

	var ss SignupStruct
	if err:= c.BodyParser(&ss); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check email
	if err := utils.CheckEmail(ss.Email); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid email"})
	}

	// check password len
	if len(ss.Password) < 8{
		return utils.JSONReponse(c, 400, fiber.Map{"error": "small password (<8)"})
	}

	// hash password
	ss.Password = hash(ss.Password)

	// check name
	if err:= utils.IsJustLetter(ss.Name, "-,"); err!=nil{
		return utils.JSONReponse(c, 400, fiber.Map{"error":err.Error})
	}
	// check family
	if err:= utils.IsJustLetter(ss.Family, "-,"); err!=nil{
		return utils.JSONReponse(c, 400, fiber.Map{"error":err.Error})
	}
	// check username
	if err:= utils.IsJustLetter(ss.Username, "-._"); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error})
	}

	// check user already exist

	var user models.User
	query := models.User{Email: ss.Email}
	if err:= database.DB.First(&user, &query).Error; err==nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"user already exist"})
	}else if err!=gorm.ErrRecordNotFound{
		return utils.ServerError(c, err)
	}

	// create token
	
	token, err := utils.CreateJWTToken(user, tokenExpHours)
	if err!=nil{
		return utils.ServerError(c, err)
	}
		
	// create user
	user.Email = ss.Email
	user.Password = ss.Password
	user.Name = ss.Name
	user.Family = ss.Family
	user.Username = ss.Username
	user.AdminPermisions = ""
	user.IsDeleted = false
	user.IsBanned = false
	user.Token = token
	user.TokenExp = time.Now().Add(time.Duration(database.Settings["tokenExpHours"])*time.Hour)

	if err:= database.DB.Create(&user).Error; err!=nil{
		return utils.ServerError(c, err)
	}


	// set token into cookie
	c.Cookie(&fiber.Cookie{
		Name: "token",
		Value: token,
		Expires: time.Now().Add(time.Duration(database.Settings["tokenExpHours"])*time.Hour),
	})
	return utils.JSONResponse(c, 200, fiber.Map{"msg": "user created"})

}

func hash(s string) string{
	h:= sha256.New()
	h.Write([]byte(s))
	return string(base64.URLEncoding.EncodeToString(h.Sum(nil)))
}

