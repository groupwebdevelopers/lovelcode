package plan

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
	pmodels "lovelcode/models/plan"
	"lovelcode/database"
	"lovelcode/utils/s3"
)

/////////////////////  public /////////////////////



func GetAllPlansAndFeatures(c *fiber.Ctx) error{
	type Result struct{
		pmodels.OPlan
		Name string
		Value string
		IsHave bool
		FeatureIsFeatured bool
	}

	page, pageLimit, err := utils.GetPageAndPageLimitFromMap(c.Queries())
	if err != nil {
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}

	var result []Result
	if err:= database.DB.Model(&pmodels.Plan{}).Select("plans.id, plans.title, plans.price, plans.image_path, plans.type, plans.is_featured, plans.is_compare, features.name, features.value, features.is_have, features.is_featured as feature_is_featured").Joins("INNER JOIN features ON plans.id=features.plan_id").Offset((page-1)*pageLimit).Limit(pageLimit).Scan(&result).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no record found"})
		}
		return utils.ServerError(c, err)
	}

	type Result2 struct{
		pmodels.OPlan
		Features []pmodels.OFeature
	}
	
	var result2 []Result2
	
	for _, r := range result{
		res1 := Result2{}
		res1.IsCompare = r.IsCompare
		res1.Title = r.Title
		res1.Price = r.Price
		res1.ImagePath = r.ImagePath
		res1.Type = r.Type


		// check id is exist
		isExist := false
		for _, r := range result2{
			if r.Title == res1.Title{
				isExist = true
				break
			}
		}
		if !isExist{
			result2 = append(result2, res1)
		}

	}
	
	for _, r := range result{
		feature := pmodels.OFeature{Name: r.Name, Value: r.Value, IsHave: r.IsHave, IsFeatured: r.FeatureIsFeatured}
		for i, r2 := range result2{
			if r2.Title == r.Title{
				result2[i].Features = append(result2[i].Features, feature)
				break
			}
		}
		
		}	
		
	return utils.JSONResponse(c, 200, fiber.Map{"data":result2})

}

func GetFeaturedPlans(c *fiber.Ctx) error{
	type Result struct{
		pmodels.OPlan
		Name string
		Value string
		IsHave bool
		FeatureIsFeatured bool
	}

	var result []Result
	if err:= database.DB.Model(&pmodels.Plan{}).Select("plans.id, plans.title, plans.price, plans.image_path, plans.type, plans.is_featured, plans.is_compare, features.name, features.value, features.is_have, features.is_featured as feature_is_featured").Joins("INNER JOIN features ON plans.id=features.plan_id").Where(&pmodels.Plan{IsFeatured: true}).Scan(&result).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no record found"})
		}
		return utils.ServerError(c, err)
	}

	type Result2 struct{
		pmodels.OPlan
		
		Features []pmodels.OFeature
	}
	
	var result2 []Result2
	
	for _, r := range result{
		res1 := Result2{}
		res1.IsCompare = r.IsCompare
		res1.Title = r.Title
		res1.Price = r.Price
		res1.ImagePath = r.ImagePath
		res1.Type = r.Type
		
		// check id is exist
		isExist := false
		for _, r := range result2{
			if r.Title == res1.Title{
				isExist = true
				break
			}
		}
		if !isExist{
			result2 = append(result2, res1)
		}

		
		result2 = append(result2, res1)

	}
	
	for _, r := range result{
		feature := pmodels.OFeature{Name: r.Name, Value: r.Value, IsHave: r.IsHave, IsFeatured: r.FeatureIsFeatured}
		for i, r2 := range result2{
			if r2.Title == r.Title{
				result2[i].Features = append(result2[i].Features, feature)
				break
			}
		}
		
		}
	
	
	return utils.JSONResponse(c, 200, fiber.Map{"data":result2})

}


func SearchPlan(c *fiber.Ctx) error{
// q := c.Queries()
// type Result struct{
// 	pmodels.OPlan

// }
// if title, ok := q["title"]; ok{
// 	// search by title
// 	if err:= 

// }else if ty, ok:= q["type"]; ok{
// 	// search by type

// }
return nil

}


/////////////////  admin   ///////////////////////////

// ---------------- plan --------------------------
// POST, auth required, admin required
func CreatePlan(c *fiber.Ctx) error{

	// plan and features
	var pl pmodels.IPlan
	if err:= c.BodyParser(&pl); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check plan validation
	if err:=pl.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
	// create plan and fill it
	var plan pmodels.Plan
	plan.Fill(&pl)
	plan.TimeCreated = time.Now()
	plan.TimeModified = time.Now()

	if err:= database.DB.Create(&plan).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 201, fiber.Map{"msg":"successfully created", "id": plan.ID})
}

// POST, Auth Required, Admin Required, /:id
// function getting plan id and a image
func UploadPlanImage(c *fiber.Ctx) error{

	id := utils.GetIDFromParams(c, "planId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the planId didn't send"})
	}
	
	
	// check plan is exist
	var plan pmodels.Plan
	if err:=database.DB.First(&plan, &pmodels.Plan{ID: id}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Plan not found"})
		}
	}

	// delete last image if exist
	if plan.ImagePath !=""{
		if strings.Contains(plan.ImagePath, "*"){
			return utils.ServerError(c, errors.New("one star is exist in image path. maybe hacker do this"))
		}
		if plan.ImagePath != ""{
			err := os.Remove(fmt.Sprintf(".%s", plan.ImagePath))
			if err!=nil{
				return utils.ServerError(c, err)
			}
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

	fl, err := file.Open()
	defer fl.Close()

	err = s3.PutObject(fl, fmt.Sprintf("/images/plan/%s", image))
	// err = c.SaveFile(file, database.Settings.ImageSaveUrl+image)

	if err!=nil{
		return utils.ServerError(c, err)
	}
	
	imageURL := fmt.Sprintf("/images/plan/%s", image)

	if err = database.DB.Model(&pmodels.Plan{}).Where(&pmodels.Plan{ID: id}).Update("image_path", imageURL).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"image added"})
}

// PUT, Auth Required, Admin Required, /:planId
func EditPlan(c *fiber.Ctx) error{
	// get id form params
	id := utils.GetIDFromParams(c, "planId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the planId didn't send"})
	}

	// check plan is exist
	var plan pmodels.Plan
	if err:= database.DB.First(&plan, &pmodels.Plan{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"plan not found"})
		}
		return utils.ServerError(c, err)
	}

	// get plan from body
	var pl pmodels.IPlan
	if err:= c.BodyParser(&pl); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}
	
	// check plan validation
	if err:=pl.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
		} 
		
	// fill the plan
	plan.Fill(&pl)
	plan.TimeModified = time.Now()
	
	// modify plan in database
	if err:= database.DB.Updates(&plan).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	
	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfully modified"})
}



// GET, admin
func GetAllPlans(c *fiber.Ctx) error{
	var plans []pmodels.Plan
	if err:= database.DB.Find(&plans).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no plan found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":plans})
}

func GetPlan(c *fiber.Ctx) error{
	
	id := utils.GetIDFromParams(c, "planId")
	if id == 0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}
	
	var plan pmodels.Plan
	if err:= database.DB.First(&plan, &pmodels.Plan{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"plan not found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":plan})
	
}

func DeletePlan(c *fiber.Ctx) error{
	id := utils.GetIDFromParams(c, "planId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}

	if err:= database.DB.Delete(&pmodels.Feature{}, &pmodels.Feature{PlanID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"feature not found"})
		}
		return utils.ServerError(c, err)
	}

	var plan pmodels.Plan
	if err:= database.DB.First(&plan, &pmodels.Plan{ID: id}).Delete(&plan).Error; err!=nil{
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
