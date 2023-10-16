package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"lovelcode/utils"
	"lovelcode/database"
	"lovelcode/models"
)




// POST, Auth Required
func CreateProjectDoingRequest(c *fiber.Ctx) error{
	
	var pdr models.CEPDR
	if err:= c.BodyParser(&pdr); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	if err:= pdr.Check();err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err})
	}

	var pd models.ProjectDoingRequest
	pd.FillWithCEPDR(pdr)
	pd.TimeCreated = time.Now()
	pd.TimeModified = time.Now()
	pd.User = c.Locals("user").(models.User)

	if err:= database.DB.Create(&pd).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 201, fiber.Map{"msg": "request saved"})
}

// GET, auth required
func GetAllProjectDoingRequests(c *fiber.Ctx) error{
	var pdrs []models.ProjectDoingRequest
	query := models.ProjectDoingRequest{UserID: c.Locals("user").(models.User).ID}
	if err:= database.DB.Find(&pdrs, &query).Error; err!=nil{
		if err == gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"you don't have any request"})
		}else{
			return utils.ServerError(c, err)
		}
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":pdrs})
}

// GET, auth required, /:id
func GetProjectDoingRequest(c *fiber.Ctx) error{
	id := utils.GetIntFromParams(c, "id")
	var pdr models.ProjectDoingRequest
	query := models.ProjectDoingRequest{ID: uint64(id)}
	if err:= database.DB.First(&pdr, &query).Error; err!=nil{
		if err == gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"the Project Doing Request with sent id not found"})
		}else{
			return utils.ServerError(c, err)
		}
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":pdr})
} 

// POST, auth required, /:id
func EditProjectDoingRequest(c *fiber.Ctx) error {
	id := utils.GetIntFromParams(c, "id")

	var pdr models.CEPDR
	if err:= c.BodyParser(&pdr); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})//, err.String())
	}

	if err:=pdr.Check();err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err})
	}

	var pd models.ProjectDoingRequest
	query := models.ProjectDoingRequest{ID: uint64(id)}
	if err:=database.DB.First(&pd, &query).Error;err!=nil{
		if err == gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"the Project Doing Request with sent id not found"})
		}else{
			return utils.ServerError(c, err)
		}
	}

	pd.FillWithCEPDR(pdr)
	pd.TimeModified = time.Now()

	if err:=database.DB.Updates(pd).Error;err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly modified"})
}

func DeleteProjectDoingRequest(c *fiber.Ctx) error{
	id := utils.GetIntFromParams(c, "id")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}

	if err:= database.DB.Delete(&models.ProjectDoingRequest{}, id).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"the id not found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly deleted"})
}