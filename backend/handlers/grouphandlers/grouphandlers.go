package grouphandlers

import (
	"fmt"
	// "strconv"
	"strings"
	// "time"
	// "errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"lovelcode/utils"
	// utilstoken "lovelcode/utils/token"
	"lovelcode/database"
	// umodels "lovelcode/models/user"
	amodels "lovelcode/models/article"
	uhandlers "lovelcode/handlers/user"
	"lovelcode/session"
)





func ApiOnly(c *fiber.Ctx) error{
	// return c.Next()
	ct, ok := c.GetReqHeaders()["Content-Type"]
	if ok && ct[0]=="application/json"{
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
	
	// user := c.Locals("user").(umodels.User)
	user, err := session.GetUserFromSession(c)
	if err!=nil{
		return utils.JSONResponse(c, 401, fiber.Map{"error":"auth required"})
	}
	// check user have permision
	splited := strings.Split(c.OriginalURL(), "/")
	if len(splited) < 5{
		return utils.JSONResponse(c, 404, fiber.Map{"error":"URL not found"})
	}
	field := splited[4]
	adminCode := utils.CheckAdminPermision(user.AdminPermisions, field)
	if adminCode != 1{
		if adminCode == 2{
			uhandlers.Hban(user.ID)
		}
		return utils.JSONResponse(c, 403, fiber.Map{"error":"Access Denied"})
	}

	return c.Next()
}

func AdminUploadImage(c *fiber.Ctx) error{
	ct, ok := c.GetReqHeaders()["Content-Type"]	
	if (len(ct) > 0 && len(ct[0]) >= 19 && ct[0][:19]!="multipart/form-data") || len(ct) == 0 || len(ct[0]) < 19 || !ok{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid header. header must be multipart/form-data"})
	}
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
		
	// user := c.Locals("user").(umodels.User)
	user, err := session.GetUserFromSession(c)
	if err!=nil{
		return utils.JSONResponse(c, 401, fiber.Map{"error":"auth required"})
	}

		splited := strings.Split(c.OriginalURL(), "/")
		if len(splited) < 6{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"URL not found"})
		}
		field := splited[3];fmt.Println(field)
		if field == "blog"{


			titleUrl := splited[5]//c.Params("articleTitleUrl")
		if titleUrl=="" || strings.Contains(titleUrl, "'"){
			return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid titleUrl"})
		}
			
			var article amodels.Article
		if err:=database.DB.First(&article, &amodels.Article{TitleUrl: titleUrl}).Error;err!=nil{
			if err== gorm.ErrRecordNotFound{
				return utils.JSONResponse(c, 404, fiber.Map{"error":"article not found"})
			}
			return utils.ServerError(c, err)
		}
			field = "edit"
			if article.UserID != user.ID{
				field += "Other"
			}else{
				field += "My"
			}
			field += "Article"
			c.Locals("article", article)
			// c.Locals("articleID", article.ID)
		}
		adminCode := utils.CheckAdminPermision(user.AdminPermisions, field)
		if adminCode != 1{
			if adminCode == 2{
			uhandlers.Hban(user.ID)
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

	
	// user := c.Locals("user").(umodels.User)
	user, err := session.GetUserFromSession(c)
	if err!=nil{
		return utils.JSONResponse(c, 401, fiber.Map{"error":"Auth Required"})
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
			return utils.JSONResponse(c, 400, fiber.Map{"error":"articleTitleUrl didn't send"})
		}

		titleUrl := splited[6]//c.Params("articleTitleUrl")
		if titleUrl=="" || strings.Contains(titleUrl, "'"){
			return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid titleUrl"})
		}
		var article amodels.Article
		if err:=database.DB.First(&article, &amodels.Article{TitleUrl: titleUrl}).Error;err!=nil{
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
			uhandlers.Hban(user.ID)
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
	// var user umodels.User
	// user, err := utilstoken.VerifyJWTToken(token)
	// if err!=nil{
	// return utils.JSONResponse(c, 401, fiber.Map{"error":"token invalid"})
	// }
	// var user gmodels.User = gmodels.User{Token: token}
		

	sess, err := session.GlobalSession.Get(c)
	if err!=nil{
		return 403, fiber.Map{"error":"you must signin! s"}, nil
	}
	defer sess.Save()

	storedToken := sess.Get("token")
	if storedToken != nil{

		if token == storedToken.(string){;fmt.Println("login with session sucesss")
		_, err := session.GetUserFromSession(c)
		if err !=nil{
			return 403, fiber.Map{"error":"you must signin! s"}, err
		}
		// c.Locals("user", user)
		return 200, nil, nil
		}
		
	}



	return 401, fiber.Map{"error":"Auth Required"}, nil
	
	// if err:=database.DB.First(&user, &umodels.User{Token: token}).Error;err!=nil{
	// 	if err==gorm.ErrRecordNotFound{
	// 		return 401, fiber.Map{"error":"authentication required"},nil
	// 	}
	// 	return 500, fiber.Map{}, err
	// }

	// // check banned
	// if user.IsBanned{
	// 	return 403, fiber.Map{"error":"you are banned!"}, nil
	// }

	// // check token
	// if token != user.Token || user.TokenExp.Unix() < time.Now().Unix(){
	// 	return 401, fiber.Map{"error":"authentication required"}, nil
	// }

	// c.Locals("user", user)

		
}

