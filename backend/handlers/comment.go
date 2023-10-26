package handlers

import (
	// "errors"
	// "fmt"
	// "os"
	// "strings"
	"time"

	"github.com/gofiber/fiber/v2"
	// "github.com/google/uuid"
	"gorm.io/gorm"

	"lovelcode/database"
	"lovelcode/models"
	"lovelcode/utils"
)

// POST, auth required /:articleTitleUrl
func CreateComment(c *fiber.Ctx) error{

	articleTitleUrl := c.Params("articleTitleUrl")

	user := c.Locals("user").(models.User)

	var mb models.IComment
	if err:= c.BodyParser(&mb); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	var articleID uint64
	if err:= database.DB.Model(&models.Article{}).Select("id").Where(models.Article{TitleUrl: articleTitleUrl}).Scan(&articleID).Error;err!=nil{
		if err == gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"article not found"})
		}
		return utils.ServerError(c, err)
	}

	// check check validation
	if err:=mb.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
	// create Comment and fill it
	var Comment models.Comment
	Comment.Fill(&mb)
	Comment.UserID = user.ID
	Comment.ArticleID = articleID
	Comment.TimeCreated = time.Now()
	Comment.TimeModified = time.Now()

	if err:= database.DB.Create(&Comment).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 201, fiber.Map{"msg":"successfully created"})
}

// PUT, auth, /:commentId
func EditComment(c *fiber.Ctx) error{
	// get id form params
	id := utils.GetIntFromParams(c, "commentId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the CommentId didn't send"})
	}

	// check Comment is exist
	var Comment models.Comment
	if err:= database.DB.First(&Comment, &models.Comment{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Comment not found"})
		}
		return utils.ServerError(c, err)
	}

	// get Comment from body
	var mb models.IComment
	if err:= c.BodyParser(&mb); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check Comment validation
	if err:=mb.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	} 

	// fill the Comment
	Comment.Fill(&mb)
	Comment.TimeModified = time.Now()

	// modify Comment in database
	if err:= database.DB.Updates(&Comment).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfully modified"})
}

// GET
func GetAllArticleComments(c *fiber.Ctx) error{
	var Comments []models.OComment
	if err:= database.DB.Model(&models.Comment{}).Select("comments.body, users.name, users.family").Joins("INNER JOIN users ON comments.user_id=users.id").Scan(&Comments).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no Comment found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":Comments})
}


// GET, admin, /:id
func GetComment(c *fiber.Ctx) error{
	
	id := utils.GetIntFromParams(c, "commentId")
	if id == 0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}
	
	var Comment models.Comment
	if err:= database.DB.First(&Comment, &models.Comment{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Comment not found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":Comment})
	
}

// DELETE, /:id
func DeleteComment(c *fiber.Ctx) error{
	id := utils.GetIntFromParams(c, "commentId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}

	isHavePermison := false

	user := c.Locals("user").(models.User)
	
	var Comment models.Comment
	if err:= database.DB.First(&Comment, models.Comment{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Comment not found"})
		}
		return utils.ServerError(c, err)
	}

	if Comment.UserID == user.ID{
		isHavePermison = true
	}else{
		p:=utils.CheckAdminPermision(user.AdminPermisions, "deleteComment")
		if p!=1{
			if p==3{
				hban(user)
			}
		}else{
			isHavePermison = true
		}

	}

	if isHavePermison{
	if err:= database.DB.Delete(&Comment).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Comment not found"})
		}
		return utils.ServerError(c, err)
	}
	
	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly deleted"})
	}

	return utils.JSONResponse(c, 403, fiber.Map{"error":"Access Denied"})
}