package plan

import (
	"errors"	
	"time"
	
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	
	"lovelcode/utils"
	pmodels "lovelcode/models/plan"
	umodels "lovelcode/models/user"
	"lovelcode/database"
)

/////////////////// public /////////////////////////



// POST
func CreatePlanOrder(c *fiber.Ctx) error{
	
	
	var pl pmodels.IPlanOrder
	if err:= c.BodyParser(&pl); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check order plan validation
	if err:=pl.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
	// create order plan and fill it
	var op pmodels.PlanOrder
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
	var PlanOrder pmodels.PlanOrder
	if err:= database.DB.First(&PlanOrder, &pmodels.PlanOrder{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"order plan not found"})
		}
		return utils.ServerError(c, err)
	}

	// get order plan from body
	var pl pmodels.IPlanOrder
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
	user:=c.Locals("user").(umodels.User)
	if user.Email == ""{
		return utils.ServerError(c, errors.New("invalid email saved in session"))
	}
	var PlanOrders []pmodels.OPlanOrder
	if err:= database.DB.Model(&pmodels.PlanOrder{}).Where(&pmodels.PlanOrder{Email: user.Email}).Offset((page-1)*pageLimit).Limit(pageLimit).Find(&PlanOrders).Error; err!=nil{
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
	
	var PlanOrder pmodels.OPlanOrder
	if err:= database.DB.Model(&pmodels.PlanOrder{}).First(&PlanOrder, &pmodels.PlanOrder{ID: id}).Error; err!=nil{
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

	if err:= database.DB.Delete(&pmodels.PlanOrder{}, id).Error; err!=nil{
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
	
	var PlanOrders []pmodels.OPlanOrder
	
	if err:= database.DB.Model(&pmodels.PlanOrder{}).Offset((page-1)*pageLimit).Limit(pageLimit).Find(&PlanOrders).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no order plan found"})
		}
		return utils.ServerError(c, err)
	}
	
	return utils.JSONResponse(c, 200, fiber.Map{"data":PlanOrders})
}
