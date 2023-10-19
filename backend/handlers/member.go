package handlers

import (
	"fmt"
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
func CreateMember(c *fiber.Ctx) error{

	var mb models.IMember
	if err:= c.BodyParser(&mb); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check check validation
	if err:=mb.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
	// create member and fill it
	var member models.Member
	member.FillWithIMember(mb)
	member.TimeCreated = time.Now()
	member.TimeModified = time.Now()

	if err:= database.DB.Create(&member).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 201, fiber.Map{"msg":"successfully created", "id": member.ID})
}

// POST, Auth Required, Admin Required, /:id
// function getting member id and a image
func UploadMemberImage(c *fiber.Ctx) error{

	id := utils.GetIntFromParams(c, "memberId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the memberId didn't send"})
	}
	
	
	// check member is exist
	var member models.Member
	if err:=database.DB.First(&member, &models.Member{ID: id}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Member not found"})
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

	if err = database.DB.Model(&models.Member{}).Where(&models.Member{ID: id}).Update("image_path", imageURL).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"image added"})
}

// PUT, Auth Required, Admin Required, /:memberId
func EditMember(c *fiber.Ctx) error{
	// get id form params
	id := utils.GetIntFromParams(c, "memberId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the memberId didn't send"})
	}

	// check member is exist
	var member models.Member
	if err:= database.DB.First(&member, &models.Member{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"member not found"})
		}
		return utils.ServerError(c, err)
	}

	// get member from body
	var mb models.IMember
	if err:= c.BodyParser(&mb); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check member validation
	if err:=mb.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	} 

	// fill the member
	member.FillWithIMember(mb)
	member.TimeModified = time.Now()

	// modify member in database
	if err:= database.DB.Updates(&member).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfully modified"})
}
