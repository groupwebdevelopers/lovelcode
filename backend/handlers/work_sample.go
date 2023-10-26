package handlers

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"lovelcode/database"
	"lovelcode/models"
	"lovelcode/utils"
)

// POST, auth required, admin required
func CreateWorkSample(c *fiber.Ctx) error{
	
	var al models.IWorkSample
	if err:= c.BodyParser(&al); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}
	
	// check check validation
	if err:=al.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
	// create WorkSample and fill it
	var WorkSample models.WorkSample
	WorkSample.Fill(&al)

	if err:= database.DB.Create(&WorkSample).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 201, fiber.Map{"msg":"successfully created"})
}

// POST, Auth Required, Admin Required, /:id
// function getting WorkSample id and a image
func UploadWorkSampleImage(c *fiber.Ctx) error{

	id := utils.GetIntFromParams(c, "workSampleId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the workSampleId didn't send"})
	}
	
	
	// check WorkSample is exist
	var WorkSample models.WorkSample
	if err:=database.DB.First(&WorkSample, &models.WorkSample{ID: id}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"WorkSample not found"})
		}
		return utils.ServerError(c, err)
	}


	// delete last image if exist
	if WorkSample.ImagePath != ""{
		if strings.Contains(WorkSample.ImagePath, "*"){
			return utils.ServerError(c, errors.New("one star is exist in image path. maybe hacker do this"))
		}
		if WorkSample.ImagePath != ""{
			err := os.Remove(fmt.Sprintf(".%s", WorkSample.ImagePath))
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
	err = c.SaveFile(file, database.Settings.ImageSaveUrl+image)

	if err!=nil{
		return utils.ServerError(c, err)
	}
	
	imageURL := fmt.Sprintf("/images/%s", image)

	if err = database.DB.Model(&models.WorkSample{}).Where(&models.WorkSample{ID: id}).Update("image_path", imageURL).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"image added"})
}

// PUT, Auth Required, Admin Required, /:WorkSampleId
func EditWorkSample(c *fiber.Ctx) error{
	// get id form params
	id := utils.GetIntFromParams(c, "WorkSampleId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the WorkSampleId didn't send"})
	}

	// check WorkSample is exist
	var WorkSample models.WorkSample
	if err:= database.DB.First(&WorkSample, &models.WorkSample{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"WorkSample not found"})
		}
		return utils.ServerError(c, err)
	}
	
	// get WorkSample from body
	var al models.IWorkSample
	if err:= c.BodyParser(&al); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}
	
	// check WorkSample validation
	if err:=al.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	} 

	// fill the WorkSample
	WorkSample.Fill(&al)

	// modify WorkSample in database
	if err:= database.DB.Updates(&WorkSample).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfully modified"})
}

// GET /:page
func GetAllWorkSamples(c *fiber.Ctx) error{

	page := utils.GetIntFromParams(c, "page")
	if page==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid page"})
	}
	var WorkSamples []models.OWorkSample
	if err:= database.DB.Model(&models.WorkSample{}).Order("id DESC").Offset((int(page)-1)*database.Settings.PageLength).Limit(database.Settings.PageLength).Find(&WorkSamples).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no WorkSample found"})
		}
		return utils.ServerError(c, err)
	}

	for i := range WorkSamples{
		date := strings.Split(WorkSamples[i].DoneTime, "T")[0]
		WorkSamples[i].DoneTime = strings.Split(utils.ConvertToPersianTime(utils.ConvertStringToTime(date, time.UTC)).String(), " ")[0]
	}

	return utils.JSONResponse(c, 200, fiber.Map{"WorkSamples":WorkSamples}) //user
}


// GET
func GetFeaturedWorkSamples(c *fiber.Ctx) error{

	var WorkSamples []models.OWorkSample
	if err:= database.DB.Model(&models.WorkSample{}).Where(&models.WorkSample{IsFeatured: true}).Find(&WorkSamples).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no WorkSample found"})
		}
		return utils.ServerError(c, err)
	}

	
	for i := range WorkSamples{
		date := strings.Split(WorkSamples[i].DoneTime, "T")[0]
		WorkSamples[i].DoneTime = strings.Split(utils.ConvertToPersianTime(utils.ConvertStringToTime(date, time.UTC)).String(), " ")[0]
	}

	return utils.JSONResponse(c, 200, fiber.Map{"WorkSamples":WorkSamples}) //user
}



// GET, admin, /:WorkSampleId
func GetWorkSample(c *fiber.Ctx) error{

	// get id form params
	id := utils.GetIntFromParams(c, "WorkSampleId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the WorkSampleId didn't send"})
	}

	
	var WorkSample models.WorkSample
	if err:= database.DB.First(&WorkSample, &models.WorkSample{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"WorkSample not found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":WorkSample})
	
}

// DELETE, /:WorkSampleId
func DeleteWorkSample(c *fiber.Ctx) error{
	id := utils.GetIntFromParams(c, "WorkSampleId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}
	WorkSample := c.Locals("WorkSample").(models.WorkSample)
	
	if err:= database.DB.Delete(&WorkSample).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"WorkSample not found"})
		}
		return utils.ServerError(c, err)
	}
	if strings.Contains(WorkSample.ImagePath, "*"){
		return utils.ServerError(c, errors.New("one star is exist in image path. maybe hacker do this"))
	}
	if WorkSample.ImagePath != ""{
		err := os.Remove(fmt.Sprintf(".%s", WorkSample.ImagePath))
		if err!=nil{
			return utils.ServerError(c, err)
		}
	}
	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly deleted"})
}