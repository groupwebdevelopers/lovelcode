package handlers

import (
	
	"time"

	"github.com/gofiber/fiber/v2"

	"lovelcode/utils"
	"lovelcode/models"
	"lovelcode/database"
)

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
// get list of features
func CreateFeatures(c *fiber.Ctx) error{
	id := utils.GetIDFromParams(c)
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
// func GetImage(c *fiber.Ctx) error{
// POST, Auth Required, Admin Required, /:id
	
	// }