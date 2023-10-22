package handlers

import (
	"lovelcode/database"
	"lovelcode/models"
	"lovelcode/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateSetting(c *fiber.Ctx) error{
	
	var st models.ISettingsDB
	if err:= c.BodyParser(&st); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}
	
	// check siteFeature validation
	if err:=st.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	var setting models.SettingsDB
	setting.FillWithISettingsDB(st)
	if err:= database.DB.Create(&setting).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	go database.RegetSettings()
	return utils.JSONResponse(c, 201, fiber.Map{"msg":"successfully created."})
}


// PUT, Auth Required, Admin Required, /:settingId
func EditSetting(c *fiber.Ctx) error{
	// get id form params
	id := utils.GetIntFromParams(c, "settingId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the settingId didn't send or invalid"})
	}

	// check setting is exist
	var setting models.SettingsDB
	if err:= database.DB.First(&setting, &models.SettingsDB{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Article not found"})
		}
		return utils.ServerError(c, err)
	}
	
	// get setting from body
	var st models.ISettingsDB
	if err:= c.BodyParser(&st); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}
	
	
	// check setting validation
	if err:=st.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	} 

	// fill the setting
	setting.FillWithISettingsDB(st)

	// modify setting in database
	if err:= database.DB.Updates(&setting).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	go database.RegetSettings()
	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfully modified"})
}

// GET , admin required
func GetAllSettings(c *fiber.Ctx) error{

	var settings []models.SettingsDB
	if err:= database.DB.Find(&settings).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no settings found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":settings}) //user
}

// GET, admin requrid /:settingId
func GetSetting(c *fiber.Ctx) error{
	
	id := utils.GetIntFromParams(c, "settingId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the settingId didn't send or invalid"})
	}
	
	var setting models.SettingsDB
	if err:= database.DB.First(&setting, &models.SettingsDB{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Setting not found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":setting})
	
}

// DELETE, admin required /:settingId
func DeleteSetting(c *fiber.Ctx) error{
	id := utils.GetIntFromParams(c, "settingId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}
	
	if err:= database.DB.Delete(&models.SettingsDB{}, id).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Article not found"})
		}
		return utils.ServerError(c, err)
	}
	go database.RegetSettings()	
	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly deleted"})
}