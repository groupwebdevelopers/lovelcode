package handlers

import (
	"time"	
	"net/mail"
	"crypto/sha256"
	
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"lovelcode/models"
	"lovelcode/database"
	"lovelcode/utils"
)


const tokenExpHours uint8 = 72

func Signin(c *fiber.Ctx) error{
	type SigninStruct struct{
		Email string `json:email`
		Password string `json:password`
	}

	// check json and extract data from it

	var ss SigninStruct
	if err:= c.BodyParser(&ss).Error; err!=nil{
		return c.Status(400).JSON(fiber.Map{"error":"invalid json"})
	}

	
	// check email
	if err := checkEmail(ss.Email); err!=nil{
		return c.Status(400).JSON(fiber.Map{"error":"invalid email"})
	}
	// hash password
	ss.Password = hash(ss.Password)

	// check user already exist

	var user models.User
	query := models.User{Email: ss.Email}
	if err:= database.DB.First(&user, &query).Error; err==gorm.ErrRecordNotFound{
		return c.Status(400).JSON(fiber.Map{"error":"user not found"})
	}else if err!=nil{
		return utils.ServerError(c, err)
	}

	// create token

	token, err := utils.CreateToken(user, tokenExpHours)
	if err!=nil{
		return utils.ServerError(c, err)
	}

	c.Cookie(&fiber.Cookie{
		Name: "token",
		Value: token,
		Expires: time.Now().Add(time.Duration(tokenExpHours)*time.Hour),
	})
	return c.Status(200).JSON(fiber.Map{"msg": "you signin"})

}


func Signup(c *fiber.Ctx) error{
	type SignupStruct struct{
		Email string `json:email`
		Password string `json:password`
		Name string `json:name`
		Family string `json:family`
	}

	// check json and extract data from it

	var ss SignupStruct
	if err:= c.BodyParser(&ss).Error; err!=nil{
		return c.Status(400).JSON(fiber.Map{"error":"invalid json"})
	}

	// check email
	if err := checkEmail(ss.Email); err!=nil{
		return c.Status(400).JSON(fiber.Map{"error":"invalid email"})
	}
	// hash password
	ss.Password = hash(ss.Password)
	

	// check user already exist

	var user models.User
	query := models.User{Email: ss.Email}
	if err:= database.DB.First(&user, &query).Error; err==nil{
		return c.Status(400).JSON(fiber.Map{"error":"user already exist"})
	}else if err!=gorm.ErrRecordNotFound{
		return utils.ServerError(c, err)
	}

	// create user

	user.Email = ss.Email
	user.Password = ss.Password
	user.Name = ss.Name
	user.Family = ss.Family
	user.AdminPermisions = ""
	user.IsDeleted = false

	if err:= database.DB.Create(&user); err!=nil{
		return c.SendStatus(500)
	}

	// create token

	token, err := utils.CreateToken(user, tokenExpHours)
	if err!=nil{
		return utils.ServerError(c, err)
	}


	c.Cookie(&fiber.Cookie{
		Name: "token",
		Value: token,
		Expires: time.Now().Add(time.Duration(tokenExpHours)*time.Hour),
	})
	return c.Status(200).JSON(fiber.Map{"msg": "user created"})

}

func hash(s string) string{
	h:= sha256.New()
	h.Write([]byte(s))
	return string(h.Sum(nil))
}

func checkEmail(e string) error{
	_, err := mail.ParseAddress(e)
	return err
}