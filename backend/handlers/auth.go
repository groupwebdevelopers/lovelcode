package handlers

import (
	
	
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"lovelcode/models"
	"lovelcode/database"
	"lovelcode/utils"
)




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

	// check user already exist

	var user models.User
	query := models.User{Email: ss.Email}
	if err:= database.DB.First(&user, &query).Error; err==gorm.ErrRecordNotFound{
		return c.Status(400).JSON(fiber.Map{"error":"user not found"})
	}else if err!=nil{
		return utils.ServerError(c, err)
	}

	// create token

	token, err := utils.CreateToken(user)
	if err!=nil{
		return utils.ServerError(c, err)
	}

	c.Cookie("token", token)
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

	token, err := utils.CreateToken(user)
	if err!=nil{
		return utils.ServerError(c, err)
	}

	c.Cookie("token", token)
	return c.Status(200).JSON(fiber.Map{"msg": "user created"})

}
