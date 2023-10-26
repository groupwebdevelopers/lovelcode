package handlers

import (
	// "time"
	"errors"
	"fmt"
	"os"
	"strings"

	"gorm.io/gorm"

	"lovelcode/database"
	"lovelcode/models"
	"lovelcode/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	// "gorm.io/gorm"
)

// GET
func Home(c *fiber.Ctx) error{
	return c.JSON(fiber.Map{"msg":"hello i'm working 1234"})
}

func GetSiteFeatures(c *fiber.Ctx) error{
	return utils.JSONResponse(c, 200, fiber.Map{"data":database.Settings.SiteFeatures})
}

// GET, /:pageName
func GetMainPage(c *fiber.Ctx) error{
	pageName := c.Params("pageName", "")
	if pageName == ""{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"empty pageName (pageName didn't send)"})
	}
	// check pageName
	if err:=utils.IsJustLetter(pageName); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid pageName:"+err.Error()})
	}
	var mpt []models.OMainpagesTexts
	if err:= database.DB.Model(&models.MainpagesTexts{}).Where(&models.MainpagesTexts{PageName: pageName}).Scan(&mpt).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	if len(mpt) ==0{
		return utils.JSONResponse(c, 404, fiber.Map{"error":"page text not found"})
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":mpt})
}


// POST, auth required, admin required
func CreateMainpage(c *fiber.Ctx) error{

	var mb models.IMainpagesTexts
	if err:= c.BodyParser(&mb); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check validation
	if err:=mb.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
	// create MainpagesTexts and fill it
	var mpt models.MainpagesTexts
	mpt.Fill(&mb)
	if err:= database.DB.Create(&mpt).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 201, fiber.Map{"msg":"successfully created"})
}

// POST, Auth Required, Admin Required, /:id
// function getting mainpage id and a image
func UploadMainpagesTextsImage(c *fiber.Ctx) error{

	id := utils.GetIntFromParams(c, "mainpagesTextsId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the mainpagesTextsId didn't send"})
	}
	
	
	// check MainpagesTexts is exist
	var mpt models.MainpagesTexts
	if err:=database.DB.First(&mpt, &models.MainpagesTexts{ID: id}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"MainpagesTexts not found"})
		}
	}
	

	file, err := c.FormFile("i")
	if err!=nil{
		return utils.ServerError(c, err)
	}

	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt	:= strings.Split(file.Filename, ".")[1]
	image := fmt.Sprintf("%s.%s", filename, fileExt)
	err = c.SaveFile(file, fmt.Sprintf("../frontend/dist/images/%s", image))

	if err!=nil{
		return utils.ServerError(c, err)
	}
	
	imageURL := fmt.Sprintf("/images/%s", image)

	if err = database.DB.Model(&models.MainpagesTexts{}).Where(&models.MainpagesTexts{ID: id}).Update("image_path", imageURL).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"image added"})
}

// PUT, admin, /:mainpagesTextsId
func EditMainpagesTexts(c *fiber.Ctx) error{
	// get id form params
	id := utils.GetIntFromParams(c, "mainpagesTextsId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the mainpagesTextsId didn't send"})
	}

	// check MainpagesTexts is exist
	var mpt models.MainpagesTexts
	if err:= database.DB.First(&mpt, &models.MainpagesTexts{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"MainpagesTexts not found"})
		}
		return utils.ServerError(c, err)
	}

	// get MainpagesTexts from body
	var mb models.IMainpagesTexts
	if err:= c.BodyParser(&mb); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check MainpagesTexts validation
	if err:=mb.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	} 

	// fill the MainpagesTexts
	mpt.Fill(&mb)

	// modify MainpagesTexts in database
	if err:= database.DB.Updates(&mpt).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfully modified"})
}

// GET, admin required
func GetAllMainpagesTexts(c *fiber.Ctx) error{
	var MainpagesTextss []models.MainpagesTexts
	if err:= database.DB.Model(&models.MainpagesTexts{}).Scan(&MainpagesTextss).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no MainpagesTexts found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":MainpagesTextss})
}

// DELETE, admin, /:id
func DeleteMainpagesTexts(c *fiber.Ctx) error{
	id := utils.GetIntFromParams(c, "mainpagesTextsId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}

	var mpt models.MainpagesTexts
	if err:= database.DB.First(&mpt, &models.MainpagesTexts{ID: id}).Delete(&mpt).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"MainpagesTexts not found"})
		}
		return utils.ServerError(c, err)
	}
	if strings.Contains(mpt.ImagePath, "*"){
		return utils.ServerError(c, errors.New("one star is exist in image path. maybe hacker do this"))
	}
	if mpt.ImagePath != ""{
		err := os.Remove(fmt.Sprintf(".%s", mpt.ImagePath))
		if err!=nil{
			return utils.ServerError(c, err)
		}
	}
	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly deleted"})
}

