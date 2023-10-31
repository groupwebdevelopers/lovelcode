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

func GetSiteSocialMedia(c *fiber.Ctx) error{
	return utils.JSONResponse(c, 200, fiber.Map{"data":database.Settings.SocialMedias})
}


func GetSitePhoneNumbers(c *fiber.Ctx) error{
	return utils.JSONResponse(c, 200, fiber.Map{"data":database.Settings.SitePhoneNumbers})
}


func GetArticleCategories(c *fiber.Ctx) error{
	return utils.JSONResponse(c, 200, fiber.Map{"data":database.Settings.ArticleCategories})
}


// GET, /:pageName
func GetMainPage(c *fiber.Ctx) error{
	pageName := c.Params("pageName", "")
	if pageName == ""{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"empty pageName (pageName didn't send)"})
	}
	// check pageName
	if err:=utils.IsJustLetter(pageName, "-"); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid pageName:"+err.Error()})
	}

	// get main page from memory
	if database.Settings.MainpageInMemory{
		var result []models.OMainpageText
		result = make([]models.OMainpageText, 0, 10)
		for _, m := range database.MainpagesTexts{
			if m.PageName == pageName{
				result = append(result, m)
			}
		}

		if len(result) > 0{
			return utils.JSONResponse(c, 200, fiber.Map{"data":result, "test":true})
		}
	}

	// get main page from database
	var mpt []models.OMainpageText
	if err:= database.DB.Model(&models.MainpageText{}).Where(&models.MainpageText{PageName: pageName}).Scan(&mpt).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	if len(mpt) ==0{
		return utils.JSONResponse(c, 404, fiber.Map{"error":"page text not found"})
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":mpt})
}


// POST, auth required, admin required
func CreateMainpageTexts(c *fiber.Ctx) error{

	var mb models.IMainpageText
	if err:= c.BodyParser(&mb); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check validation
	if err:=mb.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
	// create MainpageText and fill it
	var mpt models.MainpageText
	mpt.Fill(&mb)
	if err:= database.DB.Create(&mpt).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 201, fiber.Map{"msg":"successfully created"})
}

// POST, Auth Required, Admin Required, /:id
// function getting mainpage id and a image
func UploadMainpageTextImage(c *fiber.Ctx) error{

	id := utils.GetIDFromParams(c, "MainpageTextId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the MainpageTextId didn't send"})
	}
	
	
	// check MainpageText is exist
	var mpt models.MainpageText
	if err:=database.DB.First(&mpt, &models.MainpageText{ID: id}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"MainpageText not found"})
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

	if err = database.DB.Model(&models.MainpageText{}).Where(&models.MainpageText{ID: id}).Update("image_path", imageURL).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"image added"})
}

// PUT, admin, /:MainpageTextId
func EditMainpageText(c *fiber.Ctx) error{
	// get id form params
	id := utils.GetIDFromParams(c, "mainpageTextId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the mainpageTextId didn't send"})
	}

	// check MainpageText is exist
	var mpt models.MainpageText
	if err:= database.DB.First(&mpt, &models.MainpageText{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"MainpageText not found"})
		}
		return utils.ServerError(c, err)
	}

	// get MainpageText from body
	var mb models.IMainpageText
	if err:= c.BodyParser(&mb); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check MainpageText validation
	if err:=mb.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	} 

	// fill the MainpageText
	mpt.Fill(&mb)

	// modify MainpageText in database
	if err:= database.DB.Updates(&mpt).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfully modified"})
}

// GET, admin required
func GetAllMainpageText(c *fiber.Ctx) error{
	var MainpageTexts []models.MainpageText
	if err:= database.DB.Model(&models.MainpageText{}).Scan(&MainpageTexts).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no MainpageText found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":MainpageTexts})
}

// DELETE, admin, /:id
func DeleteMainpageText(c *fiber.Ctx) error{
	id := utils.GetIDFromParams(c, "mainpageTextId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}

	var mpt models.MainpageText
	if err:= database.DB.First(&mpt, &models.MainpageText{ID: id}).Delete(&mpt).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"MainpageText not found"})
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

