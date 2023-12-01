package article

import (

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"lovelcode/database"
	amodels "lovelcode/models/article"
	"lovelcode/utils"
)



//////////////////  public  //////////////////////////////


// GET
func GetAllArticleCategories(c *fiber.Ctx) error{
	var categories []amodels.OArticleCategory
	if err:= database.DB.Model(&amodels.ArticleCategory{}).Scan(&categories).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no category found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":categories})
}


//////////////////  admin  //////////////////////////////


// POST, Admin Required
func CreateArticleCategory(c *fiber.Ctx) error{
	
	var al amodels.IArticleCategory
	if err:= c.BodyParser(&al); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}
	
	// check check validation
	if err:=al.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
		// check article category is exist
		if err:=database.DB.First(&amodels.ArticleCategory{}, &amodels.ArticleCategory{Name: al.Name}).Error;err!=nil{
			if err!=gorm.ErrRecordNotFound{
				return utils.ServerError(c, err)
			}
		}else{
			return utils.JSONResponse(c, 400, fiber.Map{"error":"the article title already exist"})
		}
	
	// create Article category and fill it
	var category amodels.ArticleCategory
	category.Fill(&al)
	
	if err:= database.DB.Create(&category).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 201, fiber.Map{"msg":"successfully created"})
}


// PUT, Admin Required, /:id
func EditArticleCategory(c *fiber.Ctx) error{
	// get id form params
	id := utils.GetIDFromParams(c, "articleCategoryId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the articleCategoryId didn't send or invalid"})
	}

	// check category is exist
	var category amodels.ArticleCategory
	if err:= database.DB.First(&category, &amodels.ArticleCategory{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Article Category not found"})
		}
		return utils.ServerError(c, err)
	}

	
	// get article category from body
	var al amodels.IArticleCategory
	if err:= c.BodyParser(&al); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}
	
	// check article category is exist
	if err:=database.DB.First(&amodels.ArticleCategory{}, &amodels.ArticleCategory{Name: al.Name}).Error;err!=nil{
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

	
	if err:= database.DB.Delete(&amodels.ArticleCategory{}, id).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Article not found"})
		}
		return utils.ServerError(c, err)
	}
	
	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly deleted"})
}




