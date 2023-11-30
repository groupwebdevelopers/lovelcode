package user

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
	umodels "lovelcode/models/user"
	"lovelcode/utils"
	"lovelcode/utils/s3"
)

/////////////////////   public   ///////////////////////////////



// GET
func GetAllMembers(c *fiber.Ctx) error{
	var members []umodels.OMember
	if err:= database.DB.Model(&umodels.Member{}).Select("members.id, members.job_title, members.image_path, members.work_exp, members.contact, users.name, users.family, users.email").Joins("INNER JOIN users ON members.user_id=users.id").Scan(&members).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no member found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":members})
}

///////////////////   admin   ///////////////////////////////


// POST, auth required, admin required /:userId
func CreateMember(c *fiber.Ctx) error{

	id := utils.GetIDFromParams(c, "userId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the uesrId didn't send"})
	}

	// check user is exist
	var user umodels.User
	if err:=database.DB.First(&user, &umodels.User{ID: id}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"User not found"})
		}
	}


	var mb umodels.IMember
	if err:= c.BodyParser(&mb); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check check validation
	if err:=mb.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
	// create member and fill it
	var member umodels.Member
	member.Fill(&mb)
	member.UserID = id
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

	id := utils.GetIDFromParams(c, "memberId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the memberId didn't send"})
	}
	
	
	// check member is exist
	var member umodels.Member
	if err:=database.DB.First(&member, &umodels.Member{ID: id}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Member not found"})
		}
	}

	// delete last image if exist
	if member.ImagePath!=""{
		if strings.Contains(member.ImagePath, "*"){
			return utils.ServerError(c, errors.New("one star is exist in image path. maybe hacker do this"))
		}
		if member.ImagePath != ""{
			err := os.Remove(fmt.Sprintf(".%s", member.ImagePath))
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

	err = s3.PutObject(fl, fmt.Sprintf("/images/member/%s", image))
	// err = c.SaveFile(file, database.Settings.ImageSaveUrl+image)

	if err!=nil{
		return utils.ServerError(c, err)
	}
	
	imageURL := fmt.Sprintf("/images/member/%s", image)

	if err = database.DB.Model(&umodels.Member{}).Where(&umodels.Member{ID: id}).Update("image_path", imageURL).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"image added"})
}

// PUT, admin, /:memberId
func EditMember(c *fiber.Ctx) error{
	// get id form params
	id := utils.GetIDFromParams(c, "memberId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the memberId didn't send"})
	}

	// check member is exist
	var member umodels.Member
	if err:= database.DB.First(&member, &umodels.Member{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"member not found"})
		}
		return utils.ServerError(c, err)
	}

	// get member from body
	var mb umodels.IMember
	if err:= c.BodyParser(&mb); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check member validation
	if err:=mb.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	} 

	// fill the member
	member.Fill(&mb)
	member.TimeModified = time.Now()

	// modify member in database
	if err:= database.DB.Updates(&member).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfully modified"})
}


// GET, admin, /:id
func GetMember(c *fiber.Ctx) error{
	
	id := utils.GetIDFromParams(c, "memberId")
	if id == 0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}
	
	var member umodels.Member
	if err:= database.DB.First(&member, &umodels.Member{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"member not found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":member})
	
}

// DELETE, admin, /:id
func DeleteMember(c *fiber.Ctx) error{
	id := utils.GetIDFromParams(c, "memberId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}

	var member umodels.Member
	if err:= database.DB.First(&member, &umodels.Member{ID: id}).Delete(&member).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"member not found"})
		}
		return utils.ServerError(c, err)
	}
	if strings.Contains(member.ImagePath, "*"){
		return utils.ServerError(c, errors.New("one star is exist in image path. maybe hacker do this"))
	}
	if member.ImagePath != ""{
		err := os.Remove(fmt.Sprintf(".%s", member.ImagePath))
		if err!=nil{
			return utils.ServerError(c, err)
		}
	}
	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly deleted"})
}