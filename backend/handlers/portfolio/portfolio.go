package portfolio

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
	pfmodels "lovelcode/models/portfolio"
	"lovelcode/utils"
	"lovelcode/utils/s3"
	)	

///////////////////   public   /////////////////////////////

// GET /:page
func GetAllPortfolios(c *fiber.Ctx) error{

	page, pageLimit, err := utils.GetPageAndPageLimitFromMap(c.Queries())
	if err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	var Portfolios []pfmodels.OPortfolio
	if err:= database.DB.Model(&pfmodels.Portfolio{}).Order("id DESC").Offset((int(page)-1)*pageLimit).Limit(pageLimit).Find(&Portfolios).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no Portfolio found"})
		}
		return utils.ServerError(c, err)
	}

	for i := range Portfolios{
		date := strings.Split(Portfolios[i].DoneTime, "T")[0]
		Portfolios[i].DoneTime = strings.Split(utils.ConvertToPersianTime(utils.ConvertStringToTime(date, time.UTC)).String(), " ")[0]
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":Portfolios}) //user
}


// GET
func GetFeaturedPortfolios(c *fiber.Ctx) error{

	var Portfolios []pfmodels.OPortfolio
	if err:= database.DB.Model(&pfmodels.Portfolio{}).Where(&pfmodels.Portfolio{IsFeatured: true}).Find(&Portfolios).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no Portfolio found"})
		}
		return utils.ServerError(c, err)
	}

	
	for i := range Portfolios{
		date := strings.Split(Portfolios[i].DoneTime, "T")[0]
		Portfolios[i].DoneTime = strings.Split(utils.ConvertToPersianTime(utils.ConvertStringToTime(date, time.UTC)).String(), " ")[0]
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":Portfolios}) //user
}

//////////////////   admin   ///////////////////////////

	// POST, auth required, admin required
func CreatePortfolio(c *fiber.Ctx) error{
	
	var al pfmodels.IPortfolio
	if err:= c.BodyParser(&al); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}	
	
	// check check validation
	if err:=al.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}	
	
	// create Portfolio and fill it
	var Portfolio pfmodels.Portfolio
	Portfolio.Fill(&al)

	if err:= database.DB.Create(&Portfolio).Error; err!=nil{
		return utils.ServerError(c, err)
	}	
	return utils.JSONResponse(c, 201, fiber.Map{"msg":"successfully created"})
}	

// POST, Auth Required, Admin Required, /:id
// function getting Portfolio id and a image
func UploadPortfolioImage(c *fiber.Ctx) error{

	id := utils.GetIDFromParams(c, "portfolioId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the PortfolioId didn't send"})
	}	
	
	
	// check Portfolio is exist
	var Portfolio pfmodels.Portfolio
	if err:=database.DB.First(&Portfolio, &pfmodels.Portfolio{ID: id}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Portfolio not found"})
		}	
		return utils.ServerError(c, err)
	}	


	// delete last image if exist
	if Portfolio.ImagePath != ""{
		if strings.Contains(Portfolio.ImagePath, "*"){
			return utils.ServerError(c, errors.New("one star is exist in image path. maybe hacker do this"))
		}	
		if Portfolio.ImagePath != ""{
			err := os.Remove(fmt.Sprintf(".%s", Portfolio.ImagePath))
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

	err = s3.PutObject(fl, fmt.Sprintf("/images/portfolio/%s", image))
	// err = c.SaveFile(file, database.Settings.ImageSaveUrl+image)

	if err!=nil{
		return utils.ServerError(c, err)
	}	
	
	imageURL := fmt.Sprintf("/images/portfolio/%s", image)

	if err = database.DB.Model(&pfmodels.Portfolio{}).Where(&pfmodels.Portfolio{ID: id}).Update("image_path", imageURL).Error; err!=nil{
		return utils.ServerError(c, err)
	}	

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"image added"})
}	

// PUT, Auth Required, Admin Required, /:PortfolioId
func EditPortfolio(c *fiber.Ctx) error{
	// get id form params
	id := utils.GetIDFromParams(c, "portfolioId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the PortfolioId didn't send"})
	}	

	// check Portfolio is exist
	var Portfolio pfmodels.Portfolio
	if err:= database.DB.First(&Portfolio, &pfmodels.Portfolio{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Portfolio not found"})
		}	
		return utils.ServerError(c, err)
	}	
	
	// get Portfolio from body
	var al pfmodels.IPortfolio
	if err:= c.BodyParser(&al); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}	
	
	// check Portfolio validation
	if err:=al.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	} 	

	// fill the Portfolio
	Portfolio.Fill(&al)

	// modify Portfolio in database
	if err:= database.DB.Updates(&Portfolio).Error; err!=nil{
		return utils.ServerError(c, err)
	}	

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfully modified"})
}	


// GET, admin, /:PortfolioId
func GetPortfolio(c *fiber.Ctx) error{

	// get id form params
	id := utils.GetIDFromParams(c, "portfolioId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the PortfolioId didn't send"})
	}

	
	var Portfolio pfmodels.Portfolio
	if err:= database.DB.First(&Portfolio, &pfmodels.Portfolio{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Portfolio not found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":Portfolio})
	
}

// DELETE, /:PortfolioId
func DeletePortfolio(c *fiber.Ctx) error{
	id := utils.GetIDFromParams(c, "portfolioId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}

	var Portfolio pfmodels.Portfolio
	
	if err:= database.DB.First(&Portfolio, &pfmodels.Portfolio{ID: id}).Delete(&pfmodels.Portfolio{}, id).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Portfolio not found"})
		}
		return utils.ServerError(c, err)
	}
	if strings.Contains(Portfolio.ImagePath, "*"){
		return utils.ServerError(c, errors.New("one star is exist in image path. maybe hacker do this"))
	}
	if Portfolio.ImagePath != ""{
		err := os.Remove(fmt.Sprintf(".%s", Portfolio.ImagePath))
		if err!=nil{
			return utils.ServerError(c, err)
		}
	}
	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly deleted"})
}