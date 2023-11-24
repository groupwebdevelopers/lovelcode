package handlers

import (

	"gorm.io/gorm"

	"lovelcode/database"
	"lovelcode/models"
	"lovelcode/utils"

	"github.com/gofiber/fiber/v2"

)

const requestNumberLimit = 100

////////////////////////// public //////////////////////////////

// GET
func GetPublicStatistics(c *fiber.Ctx) error{
	
	var mpt []models.OStatistic
	if err:= database.DB.Model(&models.Statistic{}).Where(&models.Statistic{IsPublic: true}).Scan(&mpt).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	if len(mpt) ==0{
		return utils.JSONResponse(c, 404, fiber.Map{"error":"public statistic not found"})
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":mpt})
}

///////////////////////// admin ////////////////////////////////



// POST, auth required, admin required
func CreateStatistic(c *fiber.Ctx) error{

	var mb models.IStatistic
	if err:= c.BodyParser(&mb); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check validation
	if err:=mb.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
	// create Statistics and fill it
	var mpt models.Statistic
	mpt.Fill(&mb)
	if err:= database.DB.Create(&mpt).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 201, fiber.Map{"msg":"successfully created"})
}

// PUT, admin, /:StatisticsId
func EditStatistic(c *fiber.Ctx) error{
	// get id form params
	id := utils.GetIDFromParams(c, "statisticId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the StatisticsId didn't send"})
	}

	// check Statistics is exist
	var mpt models.Statistic
	if err:= database.DB.First(&mpt, &models.Statistic{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Statistics not found"})
		}
		return utils.ServerError(c, err)
	}

	// get Statistics from body
	var mb models.IStatistic
	if err:= c.BodyParser(&mb); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check Statistics validation
	if err:=mb.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	} 

	// fill the Statistics
	mpt.Fill(&mb)

	// modify Statistics in database
	if err:= database.DB.Updates(&mpt).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfully modified"})
}

// GET, admin required
func GetAllStatistics(c *fiber.Ctx) error{
	var Statisticss []models.Statistic
	if err:= database.DB.Model(&models.Statistic{}).Scan(&Statisticss).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no Statistics found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":Statisticss})
}

// DELETE, admin, /:id
func DeleteStatistic(c *fiber.Ctx) error{
	id := utils.GetIDFromParams(c, "statisticId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}

	if err:= database.DB.Delete(&models.Statistic{}, id).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Statistic not found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly deleted"})
}



var requestNumber uint8
// dynamic statistic
// cout number of request to server
// every {requestNumberLimit} request will add to database
func AddOneRequest(){
	if requestNumber > requestNumberLimit{
		// add one request to database
		var n uint64
		if err:= database.DB.Model(&models.Statistic{}).Select("number").Where(&models.Statistic{Name: "requestNumber"}).Scan(&n).Error;err!=nil{
			if err== gorm.ErrRecordNotFound{
				database.DB.Create(&models.Statistic{Name: "requestNumber", Name2:"request number", Number: float64(requestNumber), IsPublic: false})
				return
			}else{
				utils.LogError(err)
				return
			}
		}
		if err:= database.DB.Model(&models.Statistic{}).Where(&models.Statistic{Name: "requestNumber"}).Update("number", uint64(requestNumber)+n).Error; err!=nil{
			utils.LogError(err)
			return
		}
		requestNumber =0
	}else{
		requestNumber++
	}
}

func LogFunction() {
	AddOneRequest()
}