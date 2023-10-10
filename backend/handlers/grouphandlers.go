package handlers

import (
	"time"


	"github.com/gofiber/fiber/v2"

	"lovelcode/utils"
	"lovelcode/database"
)

func ApiOnly(c *fiber.Ctx) error{
	
	ct, ok := c.GetReqHeaders()["Content-Type"]
	if ct=="application/json" && ok==true{
		return c.Next()
	}
	return c.Status(400).JSON(fiber.Map{"error":"Content-Type must be application/json"})
}

func AuthRequired(c *fiber.Ctx) error{
	token := c.Cookies("token", "")
	if token==""{
		return utls.JSONResponse(c, 401, fiber.Map{"error":"authentication required"})
	}
	user, err := utils.VerifyJWTToken(token)
	if err!=nil{
		return utils.JSONResponse(c, 401, fiber.Map{"error":"token invalid"})
	}
	// var user models.User = models.User{Token: token}
	if err:=database.DB.First(&user, &user).Error;err!=nil{
		utils.ServerError(c, err)
	}

	// check banned
	if user.IsBanned == true{
		return utils.JSONResponse(c, 403, fiber.Map{"error":"you are banned!"})
	}

	// check token
	if token != user.Token && user.TokenExp.Unix() < time.Now().Unix(){
		return utils.JSONResponse(c, 401, fiber.Map{"error":"authentication required"})
	}

	c.Locals("user", user)
	return c.Next()
}