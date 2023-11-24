package handlers

import (
	"fmt"
	// "strconv"
	"strings"
	"time"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"

	"lovelcode/utils"
	// utilstoken "lovelcode/utils/token"
	"lovelcode/database"
	"lovelcode/models"
)


var globalSession = session.New()


func ApiOnly(c *fiber.Ctx) error{
	return c.Next()
	ct, ok := c.GetReqHeaders()["Content-Type"]
	if ct=="application/json" && ok{
		return c.Next()
	}
	return utils.JSONResponse(c, 400, fiber.Map{"error":"Content-Type must be application/json"})
}

func AuthRequired(c *fiber.Ctx) error{
	status, mp, err := authRequired(c)
	if err!=nil{
		return utils.ServerError(c, err)
	}
	if mp!=nil{
		return utils.JSONResponse(c, status, mp)
	}

	return c.Next()
}

func AdminRequired(c *fiber.Ctx) error{

	// auth requried
	status, mp, err := authRequired(c)
	if err!=nil{
		return utils.ServerError(c, err)
	}
	if mp!=nil{
		return utils.JSONResponse(c, status, mp)
	}
	
	// user := c.Locals("user").(models.User)
	userID, adminPermisions, err := getUserIDAndAdminPermisionsFromSession(c)
	if err!=nil{
		return utils.JSONResponse(c, 401, fiber.Map{"error":"auth required"})
	}
	// check user have permision
	splited := strings.Split(c.OriginalURL(), "/")
	if len(splited) < 5{
		return utils.JSONResponse(c, 404, fiber.Map{"error":"URL not found"})
	}
	field := splited[4]
	adminCode := utils.CheckAdminPermision(adminPermisions, field)
	if adminCode != 1{
		if adminCode == 2{
			hban(userID)
		}
		return utils.JSONResponse(c, 403, fiber.Map{"error":"Access Denied"})
	}

	return c.Next()
}

func AdminUploadImage(c *fiber.Ctx) error{
	
	// auth required
	status, mp, err := authRequired(c)
	if err!=nil{
		return utils.ServerError(c, err)
	}
	if mp!=nil{
		return utils.JSONResponse(c, status, mp)
	}

	token := c.Cookies("token", "")
	if token==""{
		return utils.JSONResponse(c, 401, fiber.Map{"error":"authentication required"})
	}
	
	// user, err := utilstoken.VerifyJWTToken(token)
	// if err!=nil{
		// return utils.JSONResponse(c, 401, fiber.Map{"error":"token invalid"})
		// }
		
		userID, adminPermisions, err := getUserIDAndAdminPermisionsFromSession(c)
	if err!=nil{
		return utils.JSONResponse(c, 401, fiber.Map{"error":"auth required"})
	}

		splited := strings.Split(c.OriginalURL(), "/")
		if len(splited) < 4{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"URL not found"})
		}
		field := splited[3];fmt.Println(field)
		adminCode := utils.CheckAdminPermision(adminPermisions, field)
		if adminCode != 1{
			if adminCode == 2{
				hban(userID)
			}
			return utils.JSONResponse(c, 403, fiber.Map{"error":"Access Denied"})
		}
		return c.Next()
	}

func AdminArticleRequired(c *fiber.Ctx) error{

	// auth requried
	status, mp, err := authRequired(c)
	if err!=nil{
		return utils.ServerError(c, err)
	}
	if mp!=nil{
		return utils.JSONResponse(c, status, mp)
	}

	
	userID, adminPermisions, err := getUserIDAndAdminPermisionsFromSession(c)
	if err!=nil{
		return utils.JSONResponse(c, 401, fiber.Map{"error":"auth required"})
	}
	
	
	// check user have permision
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

		if article.UserID != userID{
			field += "Other"
		}else{
			field += "My"
		}
		c.Locals("article", article)
	}
	field += "Article"
	adminCode := utils.CheckAdminPermision(adminPermisions, field)
	
	if adminCode != 1{
		if adminCode == 2{
			hban(userID)
		}
		return utils.JSONResponse(c, 403, fiber.Map{"error":"Access Denied"}, string(field))
	}
	return c.Next()

}


func authRequired(c *fiber.Ctx) (int, fiber.Map, error){
// auth required

token := c.Cookies("token", "")
if token==""{
	return 401, fiber.Map{"error":"authentication required"}, nil
}
var user models.User
// user, err := utilstoken.VerifyJWTToken(token)
// if err!=nil{
	// return utils.JSONResponse(c, 401, fiber.Map{"error":"token invalid"})
	// }
	// var user models.User = models.User{Token: token}
		

	sess, err := globalSession.Get(c)
	if err!=nil{
		return 500, nil, err
	}
	defer sess.Save()

	storedToken := sess.Get("token")
	if storedToken != nil{

		if token == sess.Get("token"){;fmt.Println("login with session sucesss")
		user.Name = sess.Get("userName").(string)
		user.Family = sess.Get("userFamily").(string)
		user.AdminPermisions = sess.Get("userAdminPermisions").(string)
		user.ID = sess.Get("userID").(uint64)
		user.Email = sess.Get("email")
		c.Locals("user", user)
		return 200, nil, nil
		}
		
	}
	
	if err:=database.DB.First(&user, &models.User{Token: token}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return 401, fiber.Map{"error":"authentication required"},nil
		}
		return 500, fiber.Map{}, err
	}

	// check banned
	if user.IsBanned{
		return 403, fiber.Map{"error":"you are banned!"}, nil
	}

	// check token
	if token != user.Token || user.TokenExp.Unix() < time.Now().Unix(){
		return 401, fiber.Map{"error":"authentication required"}, nil
	}

	c.Locals("user", user)

		
	return 200, nil, nil
}

func getUserIDAndAdminPermisionsFromSession(c *fiber.Ctx) (uint64, string, error) {
	
		sess, err := globalSession.Get(c)
		if err!=nil{
			utils.LogError(err)
			return 0, "", err
		}
	iAdminPermisions := sess.Get("adminPermsions")
	iUserID := sess.Get("userID")
	if iAdminPermisions == nil || iUserID == nil{
		return 0, "", errors.New("auth required")
	}
	 
	return iUserID.(uint64),iAdminPermisions.(string), nil

}