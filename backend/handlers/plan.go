package handlers

import (
	
	"time"
	"fmt"
	"strings"

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
	var pl models.CEPlan
	if err:= c.BodyParser(&pl); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check plan validation
	if err:=pl.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err})
	}
	
	// create plan and fill it
	var plan models.Plan
	plan.FillWithCEPlan(pl)
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
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the planId didn't send"})
	}
	
	// check plan is exist
	var plan models.Plan
	if err:=database.DB.First(&plan, &models.Plan{ID: id}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Plan not found"})
		}
	}
	
	
	// get features from body
	var ft []models.CEFeature
		if err:= c.BodyParser(&ft); err!=nil{
			return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
		}
	// check features validation
	for _, f := range ft{
		if err:= f.Check();err!=nil{
			return utils.JSONResponse(c, 400, fiber.Map{"error":err})
		}
	}

	// create features and fill its
	features := make([]models.Feature, len(ft))
	for i, _ := range features{
		features[i].FillWithCEFeature(ft[i])
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
	

	file, err := c.FormFile("image")
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

	if err = database.DB.Where(&models.Plan{ID: id}).Update("image_url", imageURL).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"image added"})
}

// PUT, Auth Required, Admin Required, /:planId
func EditPlan(c *fiber.Ctx) error{
	// get id form params
	id := utils.GetIntFromParams(c)
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the planId didn't send"})
	}

	// check plan is exist
	var plan models.Plan
	if err:= database.DB.First(&plan, &models.Plan{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordnotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"plan not found"})
		}
		return utils.ServerError(c, err)
	}

	// get plan from body
	var pl models.CEPlan
	if err:= c.BodyParser(&pl); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check plan validation
	if err:=pl.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err})
	} 

	// fill the plan
	plan.FillWithCEPlan(pl)
	plan.TimeModified = time.Now()

	// modify plan in database
	if err:= database.DB.Updates(&plan).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfully modified"})
}

func EditFeature(c *fiber.Ctx) error{
	// get id from params
	id := utils.GetIntFromParams(c, "planId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the planId didn't send"})
	}
	
	// check feature is exist
	var feature models.Feature
	if err:=database.DB.First(&feature, &models.Feature{ID: id}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Feature not found"})
		}
	}
	
	// get feature from body
	var ft models.CEFeature
	if err:= c.BodyParser(&ft); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}
	
	// check plan is exist
	var plan models.Plan
	if err:=database.DB.First(&plan, &models.Plan{ID: ft.PlanID}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Plan not found"})
		}
	}

	// check feature validation
	if err:= ft.Check();err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err})
	}
	

	// fill the feature
	feature.FillWithCEFeature(ft)
	feature.TimeModified = time.Now()

	
	if err:= database.DB.Updates(&feature).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfully modified"})
}

func GetAllPlans(c *fiber.Ctx) error{
	var plans []models.Plan
	if err:= database.DB.Find(&plans).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return JSONResponse(c, 404, fiber.Map{"error":"no plan found"})
		}
		return uitls.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":plans})
}

func GetPlan(c *fiber.Ctx) error{
	var plan models.Plan
	if err:= database.DB.First(&plan, &models.Plan{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONRespone(c, 404, fiber.Map{"error":"plan not found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":plan})
	
}