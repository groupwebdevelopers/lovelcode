package handlers

import (
	"fmt"
	// "strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"lovelcode/utils"
	// utilstoken "lovelcode/utils/token"
	"lovelcode/database"
	"lovelcode/models"
)

func ApiOnly(c *fiber.Ctx) error{
	return c.Next()
	ct, ok := c.GetReqHeaders()["Content-Type"]
	if ct=="application/json" && ok{
		return c.Next()
	}
	return utils.JSONResponse(c, 400, fiber.Map{"error":"Content-Type must be application/json"})
}

func AuthRequired(c *fiber.Ctx) error{
	token := c.Cookies("token", "")
	if token==""{
		return utils.JSONResponse(c, 401, fiber.Map{"error":"authentication required"})
	}

	var user models.User
	// user, err := utilstoken.VerifyJWTToken(token)
	// if err!=nil{
		// return utils.JSONResponse(c, 401, fiber.Map{"error":"token invalid"})
	// }
	// var user models.User = models.User{Token: token}
	if err:=database.DB.First(&user, &models.User{Token: token}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 401, fiber.Map{"error":"authentication required"})
		}
		return utils.ServerError(c, err)
	}

	// check banned
	if user.IsBanned{
		return utils.JSONResponse(c, 403, fiber.Map{"error":"you are banned!"})
	}

	// check token
	if token != user.Token || user.TokenExp.Unix() < time.Now().Unix(){
		return utils.JSONResponse(c, 401, fiber.Map{"error":"authentication required"})
	}

	c.Locals("user", user)
	return c.Next()
}

func AdminRequired(c *fiber.Ctx) error{
	// auth required
	token := c.Cookies("token", "")
	if token==""{
		return utils.JSONResponse(c, 401, fiber.Map{"error":"authentication required"})
	}

	var user models.User
	// user, err := utilstoken.VerifyJWTToken(token)
	// if err!=nil{
		// return utils.JSONResponse(c, 401, fiber.Map{"error":"token invalid"})
	// }
	// var user models.User = models.User{Token: token}
	if err:=database.DB.First(&user, &models.User{Token: token}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 401, fiber.Map{"error":"authentication required"})
		}
		return utils.ServerError(c, err)
	}

	// check banned
	if user.IsBanned{
		return utils.JSONResponse(c, 403, fiber.Map{"error":"you are banned!"})
	}

	// check token
	if token != user.Token || user.TokenExp.Unix() < time.Now().Unix(){
		return utils.JSONResponse(c, 401, fiber.Map{"error":"authentication required"})
	}

	c.Locals("user", user)

	// check user have permision
	splited := strings.Split(c.OriginalURL(), "/")
	if len(splited) < 5{
		return utils.JSONResponse(c, 404, fiber.Map{"error":"URL not found"})
	}
	field := splited[4]
	adminCode := utils.CheckAdminPermision(user.AdminPermisions, field)
	if adminCode != 1{
		if adminCode == 2{
			hban(user)
		}
		return utils.JSONResponse(c, 403, fiber.Map{"error":"Access Denied"}, string(field))
	}
	return c.Next()
}

func AdminUploadImage(c *fiber.Ctx) error{
	ct, err := c.GetReqHeaders()["Content-Type"]
	fmt.Println(ct, err)
	token := c.Cookies("token", "")
	if token==""{
		return utils.JSONResponse(c, 401, fiber.Map{"error":"authentication required"})
	}
	
	// user, err := utilstoken.VerifyJWTToken(token)
	// if err!=nil{
		// return utils.JSONResponse(c, 401, fiber.Map{"error":"token invalid"})
		// }
	var user models.User
		if err:=database.DB.First(&user, &models.User{Token: token}).Error;err!=nil{
			if err==gorm.ErrRecordNotFound{
				return utils.JSONResponse(c, 401, fiber.Map{"error":"authentication required"})
			}
			return utils.ServerError(c, err)
		}
		
		// check banned
		if user.IsBanned{
			return utils.JSONResponse(c, 403, fiber.Map{"error":"you are banned!"})
		}
		
		// check token
		if token != user.Token || user.TokenExp.Unix() < time.Now().Unix(){
			return utils.JSONResponse(c, 401, fiber.Map{"error":"authentication required"})
		}
		
		c.Locals("user", user)
		user= c.Locals("user").(models.User)

		splited := strings.Split(c.OriginalURL(), "/")
		if len(splited) < 4{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"URL not found"})
		}
		field := splited[3];fmt.Println(field)
		adminCode := utils.CheckAdminPermision(user.AdminPermisions, field)
		if adminCode != 1{
			if adminCode == 2{
				hban(user)
			}
			return utils.JSONResponse(c, 403, fiber.Map{"error":"Access Denied"})
		}
		return c.Next()
	}

func AdminArticleRequired(c *fiber.Ctx) error{

	// check user have permision
	user:= c.Locals("user").(models.User)
	splited := strings.Split(c.OriginalURL(), "/")
	if len(splited) < 6{
		return utils.JSONResponse(c, 404, fiber.Map{"error":"URL not found"})
	}
	field := splited[5]
	if field != "create"{
		// check the article is for user or not. if not add other to field
		
		if len(splited) < 7{
			return utils.JSONResponse(c, 400, fiber.Map{"error":"articleId didn't send"})
		}

		titleUrl := c.Params("titleUrl")
		if titleUrl==""{
			return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid titleUrl"})
		}
		var article models.Article
		if err:=database.DB.First(&article, &models.Article{TitleUrl: titleUrl}).Error;err!=nil{
			if err== gorm.ErrRecordNotFound{
				return utils.JSONResponse(c, 404, fiber.Map{"error":"article not found"})
			}
			return utils.ServerError(c, err)
		}

		if article.UserID != user.ID{
			field += "Other"
		}else{
			field += "My"
		}
		c.Locals("article", article)
	}
	field += "Article"
	adminCode := utils.CheckAdminPermision(user.AdminPermisions, field)
	
	if adminCode != 1{
		if adminCode == 2{
			hban(user)
		}
		return utils.JSONResponse(c, 403, fiber.Map{"error":"Access Denied"}, string(field))
	}
	return c.Next()

}