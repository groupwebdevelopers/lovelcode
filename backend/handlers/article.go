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
	utilstoken "lovelcode/utils/token"
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

		Name string
		Family string
	}
	page := utils.GetIntFromParams(c, "page")
	if page==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid page"})
	}
	var articles []Ar
	if err:= database.DB.Model(&models.Article{}).Select("articles.title, articles.title_url, articles.image_path, articles.short_desc, users.name, users.family").Joins("INNER JOIN users ON articles.user_id=users.id").Offset((int(page)-1)*database.Settings.PageLength).Limit(database.Settings.PageLength).Scan(&articles).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no Article found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"articles":articles}) //user
}


// GET /:page
func GetFeaturedArticlesTitle(c *fiber.Ctx) error{
	type Ar struct{
		Title string
		TitleUrl string
		ImagePath string
		ShortDesc string
		Views uint64

		Name string
		Family string
	}
	page := utils.GetIntFromParams(c, "page")
	if page==0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid page"})
	}
	var articles []Ar
	if err:= database.DB.Model(&models.Article{}).Select("articles.title, articles.title_url, articles.image_path, articles.short_desc, articles.views, users.name, users.family").Joins("INNER JOIN users ON articles.user_id=users.id").Where(&models.Article{IsFeatured: true}).Scan(&articles).Offset((int(page)-1)*database.Settings.PageLength).Limit(database.Settings.PageLength).Error; err!=nil{
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
	
	var article models.OArticle
	if err:= database.DB.Model(&models.Article{}).Select("articles.title, articles.body, articles.tags, articles.short_desc, articles.image_path, articles.time_created, articles.time_modified, articles.views, users.name AS user_name, users.family AS user_family, users.email as user_email").Where(&models.Article{TitleUrl: titleUrl}).Joins("INNER JOIN users ON articles.user_id=users.id").Scan(&article).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"Article not found"})
		}
		return utils.ServerError(c, err)
	}

	if article.Title == ""{
		return utils.JSONResponse(c, 404, fiber.Map{"error":"Article not found"})
	}
	
	// have view token
	viewToken := c.Cookies("Vtoken", "")
	if viewToken == ""{
		// add view token
		viewToken = utilstoken.CreateRandomToken()
		viewToken = hash(viewToken)
		c.Cookie(&fiber.Cookie{
			Name: "Vtoken",
			Value: viewToken,
			Expires: time.Now().Add(24*time.Hour),
		})
	}
	go AddViewArticle(titleUrl, viewToken, article.Views)

	return utils.JSONResponse(c, 200, fiber.Map{"data":article})
	
}

// DELETE, /:articleTitleUrl
func DeleteArticle(c *fiber.Ctx) error{
	
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

func AddViewArticle(titleUrl string, vtoken string, views uint64){
	if err:= database.DB.Model(&models.Temp{}).First(&models.Temp{}, &models.Temp{String1f: titleUrl, String2f: vtoken}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			// add view to article
			views++
			if err2:= database.DB.Model(&models.Article{}).Update("views", views).Error;err!=nil{
				utils.LogError(err2)
				return
			}
			// add titleUrl and vtoken to temp table
			ctemp := models.Temp{String1f: titleUrl, String2f: vtoken, EndTime: time.Now().Add(24*time.Hour)}
			if err2:= database.DB.Create(&ctemp).Error;err!=nil{
				utils.LogError(err2)
				return
			}

			// all done
			return
		}
		
		// server error
		utils.LogError(err)
		return
	}	
}