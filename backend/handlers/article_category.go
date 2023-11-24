package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"lovelcode/database"
	"lovelcode/models"
	"lovelcode/utils"
)



//////////////////  public  //////////////////////////////


func GetAllArticleCategories(c *fiber.Ctx) error{
	var categories []models.OArticleCategory
	if err:= database.DB.Model(&models.ArticleCategory{}).Scan(&categories).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no category found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":categories})
}


//////////////////  admin  //////////////////////////////


// POST, auth required, admin required
func CreateArticleCategory(c *fiber.Ctx) error{
	
	var al models.IArticleCategory
	if err:= c.BodyParser(&al); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}
	
	// check check validation
	if err:=al.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
		// check article category is exist
		if err:=database.DB.First(&models.ArticleCategory{}, &models.ArticleCategory{Name: al.Name}).Error;err!=nil{
			if err!=gorm.ErrRecordNotFound{
				return utils.ServerError(c, err)
			}
		}else{
			return utils.JSONResponse(c, 400, fiber.Map{"error":"the article title already exist"})
		}
	
	// create Article category and fill it
	var category models.ArticleCategory
	category.Fill(&al)
	
	if err:= database.DB.Create(&category).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 201, fiber.Map{"msg":"successfully created"})
}

// PUT, Auth Required, Admin Required, /:categoryId
func EditArticleCategory(c *fiber.Ctx) error{
	// get id form params
	id := utils.GetIDFromParams(c, "articleCategoryId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the articleCategoryId didn't send or invalid"})
	}

	// check category is exist
	var category models.ArticleCategory
	if err:= database.DB.First(&category, &models.ArticleCategory{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Article Category not found"})
		}
		return utils.ServerError(c, err)
	}

	
	// get article category from body
	var al models.IArticleCategory
	if err:= c.BodyParser(&al); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}
	
	// check article category is exist
	if err:=database.DB.First(&models.ArticleCategory{}, &models.ArticleCategory{Name: al.Name}).Error;err!=nil{
		if err!=gorm.ErrRecordNotFound{
			return utils.ServerError(c, err)
		}
	}else{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the article category english field already exist"})
	}
	
	// check Article validation
	if err:=al.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	} 

	// fill the Article
	category.Fill(&al)

	// modify Article in database
	if err:= database.DB.Updates(&category).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfully modified"})
}


// DELETE, /:id
func DeleteArticleCategory(c *fiber.Ctx) error{
	
	id := utils.GetIDFromParams(c, "articleCategoryId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the articleCategoryId didn't send or invalid"})
	}

	
	if err:= database.DB.Delete(&models.ArticleCategory{}, id).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Article not found"})
		}
		return utils.ServerError(c, err)
	}
	
	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly deleted"})
}



func convertArticleStringTimesForOutput(st []models.OArticleTitle)  {
	for i:= range st{

		st[i].TimeCreated =strings.Split(utils.ConvertStringTimeToPersianStringTime(st[i].TimeCreated), " ")[0]
		st[i].TimeModified =strings.Split(utils.ConvertStringTimeToPersianStringTime(st[i].TimeModified), " ")[0]
		
	}
}

