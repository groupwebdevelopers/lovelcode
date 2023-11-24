package handlers

import (
	"errors"	
	"time"
	"fmt"
	"strings"
	"os"
	
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	
	"lovelcode/utils"
	"lovelcode/models"
	"lovelcode/database"
)


////////////////  public //////////////////////////

// GET
func GetAllPlanTypes(c *fiber.Ctx) error{
	var planTypes []models.OPlanType
	if err:= database.DB.Model(&models.PlanType{}).Find(&planTypes).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no plan type found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":planTypes})
}


////////////////  admin //////////////////////////
// POST, auth required, admin required
func CreatePlanType(c *fiber.Ctx) error{

	
	var pl models.IPlanType
	if err:= c.BodyParser(&pl); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check plan type validation
	if err:=pl.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
	// create plan type and fill it
	var planType models.PlanType
	planType.Fill(&pl)
	planType.TimeCreated = time.Now()
	planType.TimeModified = time.Now()

	if err:= database.DB.Create(&planType).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 201, fiber.Map{"msg":"successfully created", "id": planType.ID})
}

// PUT, Auth Required, Admin Required, /:planTypeId
func EditPlanType(c *fiber.Ctx) error{
	// get id form params
	id := utils.GetIDFromParams(c, "planTypeId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the planTypeId didn't send"})
	}

	// check plan is exist
	var planType models.PlanType
	if err:= database.DB.First(&planType, &models.PlanType{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"plan type not found"})
		}
		return utils.ServerError(c, err)
	}

	// get plan type from body
	var pl models.IPlanType
	if err:= c.BodyParser(&pl); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}
	
	// check plan type validation
	if err:=pl.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
		} 
		
	// fill the plan type
	planType.Fill(&pl)
	planType.TimeModified = time.Now()
	
	// modify plan type in database
	if err:= database.DB.Updates(&planType).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	
	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfully modified"})
}



func GetPlanType(c *fiber.Ctx) error{
	
	id := utils.GetIDFromParams(c, "planTypeId")
	if id == 0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}
	
	var planType models.OPlanType
	if err:= database.DB.Model(&models.PlanType{}).First(&planType, &models.PlanType{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"plan type not found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":planType})
	
}

func DeletePlanType(c *fiber.Ctx) error{
	id := utils.GetIDFromParams(c, "planTypeId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}

	if err:= database.DB.Delete(&models.PlanType{}, id).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"plan type not found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly deleted"})
}
