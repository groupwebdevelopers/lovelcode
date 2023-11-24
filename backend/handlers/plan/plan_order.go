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

/////////////////// public /////////////////////////



// POST
func CreateOrderPlan(c *fiber.Ctx) error{
	
	
	var pl models.IOrderPlan
	if err:= c.BodyParser(&pl); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check order plan validation
	if err:=pl.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
	// create order plan and fill it
	var op models.OrderPlan
	op.Fill(&pl)
	op.TimeCreated = time.Now()
	op.TimeModified = time.Now()
	
	if err:= database.DB.Create(&op).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 201, fiber.Map{"msg":"successfully created", "id": op.ID})
}

// PUT, /:planTypeId
func EditOrderPlan(c *fiber.Ctx) error{
	// get id form params
	id := utils.GetIDFromParams(c, "orderPlanId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the orderPlanId didn't send"})
	}
	
	// check plan is exist
	var orderPlan models.OrderPlan
	if err:= database.DB.First(&orderPlan, &models.OrderPlan{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"order plan not found"})
		}
		return utils.ServerError(c, err)
	}

	// get order plan from body
	var pl models.IOrderPlan
	if err:= c.BodyParser(&pl); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}
	
	// check order plan validation
	if err:=pl.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
		} 
		
		// fill the order plan
	orderPlan.Fill(&pl)
	orderPlan.TimeModified = time.Now()
	
	// modify order plan in database
	if err:= database.DB.Updates(&orderPlan).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	
	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfully modified"})
}



// auth req
func GetAllUserOrderPlans(c *fiber.Ctx) error{
	page, pageLimit, err := utils.GetPageAndPageLimitFromMap(c.Queries())
	if err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid page"})
	}
	user:=c.Locals("user").(models.User)
	if user.Email == ""{
		return utils.ServerError(c, errors.New("invalid email saved in session"))
	}
	var orderPlans []models.OOrderPlan
	if err:= database.DB.Model(&models.OrderPlan{}).Where(&models.OrderPlan{Email: user.Email}).Offset((page-1)*pageLimit).Limit(pageLimit).Find(&orderPlans).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no order plan found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":orderPlans})
}

// auth req
func GetOrderPlan(c *fiber.Ctx) error{
	
	id := utils.GetIDFromParams(c, "orderPlanId")
	if id == 0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}
	
	var orderPlan models.OOrderPlan
	if err:= database.DB.Model(&models.OrderPlan{}).First(&orderPlan, &models.OrderPlan{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"order plan not found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":orderPlan})
	
}

func GetAllUserOrderPlans(c *fiber.Ctx) error{
	page, pageLimit, err := utils.GetPageAndPageLimitFromMap(c.Queries())
	if err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid page"})
	}
	
	var orderPlans []models.OOrderPlan

	user := c.Locals("user").(models.User)
	
	if err:= database.DB.Model(&models.OrderPlan{}).Where(&models.OrderPlan{Email: user.Email}).Offset((page-1)*pageLimit).Limit(pageLimit).Find(&orderPlans).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no order plan found"})
		}
		return utils.ServerError(c, err)
	}
	
	return utils.JSONResponse(c, 200, fiber.Map{"data":orderPlans})
}



// auth req
func DeleteOrderPlan(c *fiber.Ctx) error{
	id := utils.GetIDFromParams(c, "orderPlanId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}

	if err:= database.DB.Delete(&models.OrderPlan{}, id).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"order plan not found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly deleted"})
}
/////////////////// admin /////////////////////////

// GET, admin
func GetAllOrderPlans(c *fiber.Ctx) error{
	page, pageLimit, err := utils.GetPageAndPageLimitFromMap(c.Queries())
	if err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid page"})
	}
	
	var orderPlans []models.OOrderPlan
	
	if err:= database.DB.Model(&models.OrderPlan{}).Offset((page-1)*pageLimit).Limit(pageLimit).Find(&orderPlans).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no order plan found"})
		}
		return utils.ServerError(c, err)
	}
	
	return utils.JSONResponse(c, 200, fiber.Map{"data":orderPlans})
}
