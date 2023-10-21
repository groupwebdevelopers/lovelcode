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
func CreateArticle(c *fiber.Ctx) error{
	
	var al models.IArticle
	if err:= c.BodyParser(&al); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}
	
	// check check validation
	if err:=al.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
		// check article is exist
		if err:=database.DB.First(&models.Article{}, &models.Article{TitleUrl: utils.ConvertToUrl(al.Title)}).Error;err!=nil{
			if err!=gorm.ErrRecordNotFound{
				return utils.ServerError(c, err)
			}
		}else{
			return utils.JSONResponse(c, 400, fiber.Map{"error":"the article title already exist"})
		}
	
	// create Article and fill it
	var article models.Article
	article.FillWithIArticle(al)
	article.UserID = c.Locals("user").(models.User).ID
	article.TimeCreated = time.Now()
	article.TimeModified = time.Now()
	article.TitleUrl = utils.ConvertToUrl(article.Title)

	if err:= database.DB.Create(&article).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 201, fiber.Map{"msg":"successfully created"})
}

// POST, Auth Required, Admin Required, /:id
// function getting article id and a image
func UploadArticleImage(c *fiber.Ctx) error{

	id := utils.GetIntFromParams(c, "articleId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the articleId didn't send"})
	}
	
	
	// check Article is exist
	var article models.Article
	if err:=database.DB.First(&article, &models.Article{ID: id}).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Article not found"})
		}
		return utils.ServerError(c, err)
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

	if err = database.DB.Model(&models.Article{}).Where(&models.Article{ID: id}).Update("image_path", imageURL).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"image added"})
}

// PUT, Auth Required, Admin Required, /:articleId
func EditArticle(c *fiber.Ctx) error{
	// // get id form params
	// id := utils.GetIntFromParams(c, "articleId")
	// if id==0{
	// 	return utils.JSONResponse(c, 400, fiber.Map{"error":"the articleId didn't send"})
	// }

	// // check article is exist
	// var article models.Article
	// if err:= database.DB.First(&article, &models.Article{ID: id}).Error; err!=nil{
	// 	if err==gorm.ErrRecordNotFound{
	// 		return utils.JSONResponse(c, 404, fiber.Map{"error":"Article not found"})
	// 	}
	// 	return utils.ServerError(c, err)
	// }

	article := c.Locals("article").(models.Article)

	
	// get article from body
	var al models.IArticle
	if err:= c.BodyParser(&al); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}
	
	// check article is exist
	if err:=database.DB.First(&models.Article{}, &models.Article{TitleUrl: utils.ConvertToUrl(al.Title)}).Error;err!=nil{
		if err!=gorm.ErrRecordNotFound{
			return utils.ServerError(c, err)
		}
	}else{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"the article title already exist"})
	}
	
	// check Article validation
	if err:=al.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	} 

	// fill the Article
	article.FillWithIArticle(al)
	article.TimeModified = time.Now()

	// modify Article in database
	if err:= database.DB.Updates(&article).Error; err!=nil{
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfully modified"})
}

// GET /:page
func GetAllArticlesTitles(c *fiber.Ctx) error{
	type Ar struct{
		Title string
		TitleUrl string
		ImagePath string
		ShortDesc string
	}
	page := utils.GetIntFromParams(c, "page")
	if page==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid page"})
	}
	var articles []Ar
	if err:= database.DB.Model(&models.Article{}).Select("title, title_url, image_path, short_desc").Find(&articles).Offset(int((page-1)*20)).Limit(20).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no Article found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"articles":articles}) //user
}

// GET, /:articleTitleUrl
func GetArticle(c *fiber.Ctx) error{
	
	titleUrl := c.Params("articleTitleUrl")
	
	var article models.Article
	if err:= database.DB.First(&article, &models.Article{TitleUrl: titleUrl}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Article not found"})
		}
		return utils.ServerError(c, err)
	}

	var user models.User
	if err:= database.DB.First(&user, &models.User{ID: article.UserID}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"user that created article not found"})
		}
		return utils.ServerError(c, err)
	}

	var oArticle models.OArticle
	oArticle.FillWithArticle(article)
	oArticle.OUser.FillWithUser(user)

	return utils.JSONResponse(c, 200, fiber.Map{"data":oArticle})
	
}

// DELETE, /:articleId
func DeleteArticle(c *fiber.Ctx) error{
	id := utils.GetIntFromParams(c, "articleId")
	if id==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}
	article := c.Locals("article").(models.Article)
	
	if err:= database.DB.Delete(&article).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Article not found"})
		}
		return utils.ServerError(c, err)
	}
	if strings.Contains(article.ImagePath, "*"){
		return utils.ServerError(c, errors.New("one star is exist in image path. maybe hacker do this"))
	}
	if article.ImagePath != ""{
		err := os.Remove(fmt.Sprintf(".%s", article.ImagePath))
		if err!=nil{
			return utils.ServerError(c, err)
		}
	}
	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly deleted"})
}