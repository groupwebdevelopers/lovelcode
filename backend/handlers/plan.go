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

// POST, auth required, admin required
func CreatePlan(c *fiber.Ctx) error{

	// plan and features
	var pl models.IPlan
	if err:= c.BodyParser(&pl); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check plan validation
	if err:=pl.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
	// create plan and fill it
	var plan models.Plan
	plan.FillWithIPlan(pl)
	plan.TimeCreated = time.Now()
	plan.TimeModified = time.Now()

	if err:= database.DB.Create(&plan).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 201, fiber.Map{"msg":"successfully created", "id": plan.ID})
}

// POST, Auth Required, /:id
// function getting list of features
func CreateFeatures(c *fiber.Ctx) error{

	id := utils.GetIntFromParams(c, "planId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the planId is invalid"})
	}
	
	// check plan is exist
	var plan models.Plan
	if err:=database.DB.First(&plan, &models.Plan{ID: id}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Plan not found"})
		}
		return utils.ServerError(c, err)
	}
	
	
	// get features from body
	var ft []models.IFeature
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
	features := make([]models.Feature, len(ft))
	for i:= range features{
		features[i].FillWithIFeature(ft[i])
		features[i].TimeCreated = time.Now()
		features[i].TimeModified = time.Now()
		features[i].PlanID = uint64(id)
	}

	
	if err:= database.DB.Create(&features).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 201, fiber.Map{"msg":"successfully created"})

}

// POST, Auth Required, Admin Required, /:id
// function getting plan id and a image
func UploadPlanImage(c *fiber.Ctx) error{

	id := utils.GetIntFromParams(c, "planId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the planId didn't send"})
	}
	
	
	// check plan is exist
	var plan models.Plan
	if err:=database.DB.First(&plan, &models.Plan{ID: id}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Plan not found"})
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

	if err = database.DB.Model(&models.Plan{}).Where(&models.Plan{ID: id}).Update("image_path", imageURL).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"image added"})
}

// PUT, Auth Required, Admin Required, /:planId
func EditPlan(c *fiber.Ctx) error{
	// get id form params
	id := utils.GetIntFromParams(c, "planId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the planId didn't send"})
	}

	// check plan is exist
	var plan models.Plan
	if err:= database.DB.First(&plan, &models.Plan{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"plan not found"})
		}
		return utils.ServerError(c, err)
	}

	// get plan from body
	var pl models.IPlan
	if err:= c.BodyParser(&pl); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check plan validation
	if err:=pl.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	} 

	// fill the plan
	plan.FillWithIPlan(pl)
	plan.TimeModified = time.Now()

	// modify plan in database
	if err:= database.DB.Updates(&plan).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfully modified"})
}

func EditFeature(c *fiber.Ctx) error{
	// get id from params
	id := utils.GetIntFromParams(c, "featureId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the featureId is invalid"})
	}
	
	// check feature is exist
	var feature models.Feature
	if err:=database.DB.First(&feature, &models.Feature{ID: id}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Feature not found"})
		}
	}
	
	// get feature from body
	var ft models.IFeature
	if err:= c.BodyParser(&ft); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}
	
	// check plan is exist
	var plan models.Plan
	if err:=database.DB.First(&plan, &models.Plan{ID: id}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Plan not found"})
		}
	}

	// check feature validation
	if err:= ft.Check();err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	

	// fill the feature
	feature.FillWithIFeature(ft)
	feature.TimeModified = time.Now()
	feature.PlanID = id
	
	if err:= database.DB.Updates(&feature).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfully modified"})
}

func GetAllPlans(c *fiber.Ctx) error{
	var plans []models.Plan
	if err:= database.DB.Find(&plans).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no plan found"})
		}
		return utils.ServerError(c, err)
	}

	oplans := make([]models.OPlan, len(plans))
	for i, pl := range plans{
		oplans[i].FillWithPlan(pl)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":oplans})
}

func GetPlan(c *fiber.Ctx) error{
	
	id := utils.GetIntFromParams(c, "planId")
	if id == 0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}
	
	var plan models.Plan
	if err:= database.DB.First(&plan, &models.Plan{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"plan not found"})
		}
		return utils.ServerError(c, err)
	}

	var oplan models.OPlan
	oplan.FillWithPlan(plan)

	return utils.JSONResponse(c, 200, fiber.Map{"data":oplan})
	
}

func GetFeature(c *fiber.Ctx) error{

	id := utils.GetIntFromParams(c, "featureId")
	if id == 0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}
	var feature models.Feature
	if err:= database.DB.First(&feature, &models.Feature{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"plan not found"})
		}
		return utils.ServerError(c, err)
	}

	var ofeature models.OFeature
	ofeature.FillWithFeature(feature)

	return utils.JSONResponse(c, 200, fiber.Map{"data":ofeature})
}

func GetAllFeatures(c *fiber.Ctx) error{
	var features []models.Feature

	if err:= database.DB.Find(&features).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no feature found"})
		}
		return utils.ServerError(c, err)
	}

	ofeatures := make([]models.OFeature, len(features))
	for i, f := range features{
		ofeatures[i].FillWithFeature(f)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":ofeatures})

}

func GetAllPlansAndFeatures(c *fiber.Ctx) error{
	var plans []models.Plan
	var features []models.Feature
	if err:= database.DB.Find(&plans).Find(&features).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no record found"})
		}
		return utils.ServerError(c, err)
	}

	
	oplans := make([]models.OPlan, len(plans))
	for i, pl := range plans{
		oplans[i].FillWithPlan(pl)
	}

	ofeatures := make([]models.OFeature, len(features))
	for i, f := range features{
		ofeatures[i].FillWithFeature(f)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":fiber.Map{"features":ofeatures, "plans":oplans}})

}

func DeletePlan(c *fiber.Ctx) error{
	id := utils.GetIntFromParams(c, "planId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}

	if err:= database.DB.Delete(&models.Feature{}, &models.Feature{PlanID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"feature not found"})
		}
		return utils.ServerError(c, err)
	}

	var plan models.Plan
	if err:= database.DB.First(&plan, &models.Plan{ID: id}).Delete(&plan).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"plan not found"})
		}
		return utils.ServerError(c, err)
	}
	if strings.Contains(plan.ImagePath, "*"){
		return utils.ServerError(c, errors.New("one star is exist in image path. maybe hacker do this"))
	}
	if plan.ImagePath != ""{
		err := os.Remove(fmt.Sprintf(".%s", plan.ImagePath))
		if err!=nil{
			return utils.ServerError(c, err)
		}
	}
	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly deleted"})
}


func DeleteFeature(c *fiber.Ctx) error{
	id := utils.GetIntFromParams(c, "featureId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}

	if err:= database.DB.Delete(&models.Feature{}, &models.Feature{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"feature not found"})
		}
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly deleted"})
}