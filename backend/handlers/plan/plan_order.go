package handlers

import (
	"errors"	
	"time"
	
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	
	"lovelcode/utils"
	"lovelcode/models"
	"lovelcode/database"
)

/////////////////// public /////////////////////////



// POST
func CreatePlanOrder(c *fiber.Ctx) error{
	
	
	var pl models.IPlanOrder
	if err:= c.BodyParser(&pl); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check order plan validation
	if err:=pl.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
	// create order plan and fill it
	var op models.PlanOrder
	op.Fill(&pl)
	op.TimeCreated = time.Now()
	op.TimeModified = time.Now()
	
	if err:= database.DB.Create(&op).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 201, fiber.Map{"msg":"successfully created", "id": op.ID})
}

// PUT, /:planTypeId
func EditPlanOrder(c *fiber.Ctx) error{
	// get id form params
	id := utils.GetIDFromParams(c, "PlanOrderId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the PlanOrderId didn't send"})
	}
	
	// check plan is exist
	var PlanOrder models.PlanOrder
	if err:= database.DB.First(&PlanOrder, &models.PlanOrder{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"order plan not found"})
		}
		return utils.ServerError(c, err)
	}

	// get order plan from body
	var pl models.IPlanOrder
	if err:= c.BodyParser(&pl); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}
	
	// check order plan validation
	if err:=pl.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
		} 
		
		// fill the order plan
	PlanOrder.Fill(&pl)
	PlanOrder.TimeModified = time.Now()
	
	// modify order plan in database
	if err:= database.DB.Updates(&PlanOrder).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	
	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfully modified"})
}



// auth req
func GetAllUserPlanOrders(c *fiber.Ctx) error{
	page, pageLimit, err := utils.GetPageAndPageLimitFromMap(c.Queries())
	if err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid page"})
	}
	user:=c.Locals("user").(models.User)
	if user.Email == ""{
		return utils.ServerError(c, errors.New("invalid email saved in session"))
	}
	var PlanOrders []models.OPlanOrder
	if err:= database.DB.Model(&models.PlanOrder{}).Where(&models.PlanOrder{Email: user.Email}).Offset((page-1)*pageLimit).Limit(pageLimit).Find(&PlanOrders).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no order plan found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":PlanOrders})
}

// auth req
func GetPlanOrder(c *fiber.Ctx) error{
	
	id := utils.GetIDFromParams(c, "PlanOrderId")
	if id == 0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}
	
	var PlanOrder models.OPlanOrder
	if err:= database.DB.Model(&models.PlanOrder{}).First(&PlanOrder, &models.PlanOrder{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"order plan not found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":PlanOrder})
	
}


// auth req
func DeletePlanOrder(c *fiber.Ctx) error{
	id := utils.GetIDFromParams(c, "PlanOrderId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}

	if err:= database.DB.Delete(&models.PlanOrder{}, id).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"order plan not found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly deleted"})
}
/////////////////// admin /////////////////////////

// GET, admin
func GetAllPlanOrders(c *fiber.Ctx) error{
	page, pageLimit, err := utils.GetPageAndPageLimitFromMap(c.Queries())
	if err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid page"})
	}
	
	var PlanOrders []models.OPlanOrder
	
	if err:= database.DB.Model(&models.PlanOrder{}).Offset((page-1)*pageLimit).Limit(pageLimit).Find(&PlanOrders).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no order plan found"})
		}
		return utils.ServerError(c, err)
	}
	
	return utils.JSONResponse(c, 200, fiber.Map{"data":PlanOrders})
}
