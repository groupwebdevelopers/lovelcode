package plan

import (
	"time"
	"fmt"
	
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	
	"lovelcode/utils"
	pmodels "lovelcode/models/plan"
	"lovelcode/database"
)


/////////////////// public /////////////////////////


/////////////////// admin /////////////////////////

// POST, Auth Required, /:id
// function getting list of features
func CreateFeatures(c *fiber.Ctx) error{

	id := utils.GetIDFromParams(c, "planId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the planId is invalid"})
	}
	
	// check plan is exist
	var plan pmodels.Plan
	if err:=database.DB.First(&plan, &pmodels.Plan{ID: id}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Plan not found"})
		}
		return utils.ServerError(c, err)
	}
	
	
	// get features from body
	var ft []pmodels.IFeature
		if err:= c.BodyParser(&ft); err!=nil{
			fmt.Println(err)
			return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
		}
	// check features validation
	for _, f := range ft{
		if err:= f.Check();err!=nil{
			return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
		}
	}

	// create features and fill its
	features := make([]pmodels.Feature, len(ft))
	for i:= range features{
		features[i].Fill(&ft[i])
		features[i].TimeCreated = time.Now()
		features[i].TimeModified = time.Now()
		features[i].PlanID = uint64(id)
	}

	
	if err:= database.DB.Create(&features).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 201, fiber.Map{"msg":"successfully created"})

}

func EditFeature(c *fiber.Ctx) error{
	// get id from params
	id := utils.GetIDFromParams(c, "featureId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the featureId is invalid"})
	}
	
	// check feature is exist
	var feature pmodels.Feature
	if err:=database.DB.First(&feature, &pmodels.Feature{ID: id}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Feature not found"})
		}
	}
	
	// get feature from body
	var ft pmodels.IFeature
	if err:= c.BodyParser(&ft); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}
	
	// check plan is exist
	var plan pmodels.Plan
	if err:=database.DB.First(&plan, &pmodels.Plan{ID: id}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Plan not found"})
		}
	}
	
	// check feature validation
	if err:= ft.Check();err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
	
	// fill the feature
	feature.Fill(&ft)
	feature.TimeModified = time.Now()
	feature.PlanID = id
	
	if err:= database.DB.Updates(&feature).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfully modified"})
}

func GetFeature(c *fiber.Ctx) error{

	id := utils.GetIDFromParams(c, "featureId")
	if id == 0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}
	var feature pmodels.Feature
	if err:= database.DB.First(&feature, &pmodels.Feature{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"plan not found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":feature})
}

func GetAllFeatures(c *fiber.Ctx) error{
	var features []pmodels.Feature

	if err:= database.DB.Find(&features).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no feature found"})
			}	
			return utils.ServerError(c, err)
		}	
			
	return utils.JSONResponse(c, 200, fiber.Map{"data":features})
			
}	


func DeleteFeature(c *fiber.Ctx) error{
	id := utils.GetIDFromParams(c, "featureId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}

	if err:= database.DB.Delete(&pmodels.Feature{}, id).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"feature not found"})
		}
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly deleted"})
}
