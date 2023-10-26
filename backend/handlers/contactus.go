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
func CreateContactUs(c *fiber.Ctx) error{

	user := c.Locals("user").(models.User)

	var mb models.IContactUs
	if err:= c.BodyParser(&mb); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check check validation
	if err:=mb.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
	// create ContactUs and fill it
	var ContactUs models.ContactUs
	ContactUs.Fill(&mb)
	ContactUs.TitleUrl = utils.ConvertToUrl(ContactUs.Title)
	ContactUs.UserID = user.ID
	ContactUs.TimeCreated = time.Now()
	ContactUs.TimeModified = time.Now()

	if err:= database.DB.Create(&ContactUs).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 201, fiber.Map{"msg":"successfully created"})
}

// PUT, auth, /:ContactUsTitle
func EditContactUs(c *fiber.Ctx) error{
	
	// get title form params
	title := c.Params("title", "")
	if title == ""{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid title"})
	}

	// check ContactUs is exist
	var ContactUs models.ContactUs
	if err:= database.DB.First(&ContactUs, &models.ContactUs{Title: title}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"ContactUs not found"})
		}
		return utils.ServerError(c, err)
	}

	// get ContactUs from body
	var mb models.IContactUs
	if err:= c.BodyParser(&mb); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check ContactUs validation
	if err:=mb.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	} 

	// fill the ContactUs
	ContactUs.Fill(&mb)
	ContactUs.TitleUrl = utils.ConvertToUrl(ContactUs.Title)
	ContactUs.TimeModified = time.Now()

	// modify ContactUs in database
	if err:= database.DB.Updates(&ContactUs).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfully modified"})
}

// GET, auth
func GetAllUserContactUss(c *fiber.Ctx) error{
	user := c.Locals("user").(models.User)

	var ContactUss []models.OContactUs
	if err:= database.DB.Model(&models.ContactUs{}).Where(&models.ContactUs{UserID: user.ID}).Find(&ContactUss).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no ContactUs found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":ContactUss})
}

// GET, auth, /:title
func GetContactUsByTitle(c *fiber.Ctx) error{
	
	title := c.Params("title", "")
	if title == ""{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid title"})
	}
	
	var ContactUs models.ContactUs
	if err:= database.DB.First(&ContactUs, &models.ContactUs{Title: title}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"ContactUs not found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":ContactUs})
	
}

// DELETE, /:title
func DeleteContactUs(c *fiber.Ctx) error{
	
	title := c.Params("title", "")
	if title == ""{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid title"})
	}

	isHavePermison := false

	user := c.Locals("user").(models.User)
	
	var ContactUs models.ContactUs
	if err:= database.DB.First(&ContactUs, models.ContactUs{Title: title}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"ContactUs not found"})
		}
		return utils.ServerError(c, err)
	}

	if ContactUs.UserID == user.ID{
		isHavePermison = true
	}else{
		p:=utils.CheckAdminPermision(user.AdminPermisions, "deleteContactUs")
		if p!=1{
			if p==3{
				hban(user)
			}
		}else{
			isHavePermison = true
		}

	}

	if isHavePermison{
	if err:= database.DB.Delete(&ContactUs).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"ContactUs not found"})
		}
		return utils.ServerError(c, err)
	}
	
	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly deleted"})
	}

	return utils.JSONResponse(c, 403, fiber.Map{"error":"Access Denied"})
}