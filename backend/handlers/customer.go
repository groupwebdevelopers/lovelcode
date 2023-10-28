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

/////////////////  public    //////////////////////////////////


// GET
func GetAllCustomers(c *fiber.Ctx) error{
	var Customers []models.OCustomer
	if err:= database.DB.Find(&Customers).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no Customer found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":Customers})
}


// GET
func GetFeaturedCustomers(c *fiber.Ctx) error{
	var Customers []models.OCustomer
	if err:= database.DB.Find(&Customers, models.Customer{IsFeatured: true}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no Customer found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":Customers})
}


////////////////////   admin    /////////////////////////////////////

// POST, auth required, admin required 
func CreateCustomer(c *fiber.Ctx) error{

	var mb models.ICustomer
	if err:= c.BodyParser(&mb); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check check validation
	if err:=mb.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
	// create Customer and fill it
	var Customer models.Customer
	Customer.Fill(&mb)
	Customer.TimeCreated = time.Now()
	Customer.TimeModified = time.Now()

	if err:= database.DB.Create(&Customer).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 201, fiber.Map{"msg":"successfully created", "id": Customer.ID})
}

// POST, Auth Required, Admin Required, /:id
// function getting Customer id and a image
func UploadCustomerImage(c *fiber.Ctx) error{

	id := utils.GetIntFromParams(c, "customerId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the CustomerId didn't send"})
	}
	
	
	// check Customer is exist
	var Customer models.Customer
	if err:=database.DB.First(&Customer, &models.Customer{ID: id}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Customer not found"})
		}
	}

	// delete last image if exist
	if Customer.ImagePath!=""{
		if strings.Contains(Customer.ImagePath, "*"){
			return utils.ServerError(c, errors.New("one star is exist in image path. maybe hacker do this"))
		}
		if Customer.ImagePath != ""{
			err := os.Remove(fmt.Sprintf(".%s", Customer.ImagePath))
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

	if err = database.DB.Model(&models.Customer{}).Where(&models.Customer{ID: id}).Update("image_path", imageURL).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"image added"})
}

// PUT, admin, /:CustomerId
func EditCustomer(c *fiber.Ctx) error{
	// get id form params
	id := utils.GetIntFromParams(c, "customerId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the CustomerId didn't send"})
	}

	// check Customer is exist
	var Customer models.Customer
	if err:= database.DB.First(&Customer, &models.Customer{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Customer not found"})
		}
		return utils.ServerError(c, err)
	}

	// get Customer from body
	var mb models.ICustomer
	if err:= c.BodyParser(&mb); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check Customer validation
	if err:=mb.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	} 

	// fill the Customer
	Customer.Fill(&mb)
	Customer.TimeModified = time.Now()

	// modify Customer in database
	if err:= database.DB.Updates(&Customer).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfully modified"})
}


// GET, admin, /:id
func GetCustomer(c *fiber.Ctx) error{
	
	id := utils.GetIntFromParams(c, "customerId")
	if id == 0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}
	
	var Customer models.Customer
	if err:= database.DB.First(&Customer, &models.Customer{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Customer not found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":Customer})
	
}

// DELETE, admin, /:id
func DeleteCustomer(c *fiber.Ctx) error{
	id := utils.GetIntFromParams(c, "customerId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}

	var Customer models.Customer
	if err:= database.DB.First(&Customer, &models.Customer{ID: id}).Delete(&Customer).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Customer not found"})
		}
		return utils.ServerError(c, err)
	}
	if strings.Contains(Customer.ImagePath, "*"){
		return utils.ServerError(c, errors.New("one star is exist in image path. maybe hacker do this"))
	}
	if Customer.ImagePath != ""{
		err := os.Remove(fmt.Sprintf(".%s", Customer.ImagePath))
		if err!=nil{
			return utils.ServerError(c, err)
		}
	}
	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly deleted"})
}