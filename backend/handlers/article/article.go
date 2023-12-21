package article

import (
	"errors"
	"fmt"
	"path"
	"slices"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	main_session "github.com/gofiber/fiber/v2/middleware/session"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"lovelcode/database"
	amodels "lovelcode/models/article"

	// umodels "lovelcode/models/user"
	// tmodels "lovelcode/models/temp"
	"lovelcode/session"
	"lovelcode/utils"
	"lovelcode/utils/hutils"
	"lovelcode/utils/s3"
	// utilstoken "lovelcode/utils/token"
)

///////////////////////   public   ///////////////////////////

// GET /?page=x&pageLimit=x
func GetAllArticlesTitles(c *fiber.Ctx) error {

	// get queries from url
	page, pageLimit, err := utils.GetPageAndPageLimitFromMap(c.Queries())
	if err != nil {
		return utils.JSONResponse(c, 400, fiber.Map{"error": err.Error()})
	}

	// fetch data from database
	var articles []amodels.OArticleTitle
	if err := database.DB.Model(&amodels.Article{}).Select("articles.title, articles.title_url, articles.image_path, articles.short_desc, articles.time_created, articles.time_modified, articles.views, articles.likes, users.name AS user_name, users.family AS user_family").Joins("INNER JOIN users ON articles.user_id=users.id").Offset((int(page) - 1) * pageLimit).Limit(pageLimit).Scan(&articles).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.JSONResponse(c, 404, fiber.Map{"error": "no Article found"})
		}
		return utils.ServerError(c, err)
	}

	// fix time for output
	hutils.ConvertArticleStringTimesForOutput(articles)

	return utils.JSONResponse(c, 200, fiber.Map{"data": articles}) //user
}

// GET /?page=x&pageLimit=x
func GetFeaturedArticlesTitle(c *fiber.Ctx) error {

	// get queries from url
	page, pageLimit, err := utils.GetPageAndPageLimitFromMap(c.Queries())
	if err != nil {
		return utils.JSONResponse(c, 400, fiber.Map{"error": err.Error()})
	}

	// fetch from database
	var articles []amodels.OArticleTitle
	if err := database.DB.Model(&amodels.Article{}).Select("articles.title, articles.title_url, articles.image_path, articles.short_desc, articles.views,articles.time_created, articles.time_modified, articles.views, articles.likes, users.name AS user_name, users.family AS user_name").Joins("INNER JOIN users ON articles.user_id=users.id").Where(&amodels.Article{IsFeatured: true}).Scan(&articles).Offset((int(page) - 1) * pageLimit).Limit(pageLimit).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.JSONResponse(c, 404, fiber.Map{"error": "no Article found"})
		}
		return utils.ServerError(c, err)
	}

	hutils.ConvertArticleStringTimesForOutput(articles)

	return utils.JSONResponse(c, 200, fiber.Map{"data": articles}) //user
}

// GET, /:articleTitleUrl
func GetArticle(c *fiber.Ctx) error {

	titleUrl := c.Params("articleTitleUrl")
	var article amodels.OArticle
	
	if err := database.DB.Model(&amodels.Article{}).Select("articles.id, articles.title, articles.body, articles.tags, articles.short_desc, articles.image_path, articles.time_created, articles.time_modified, articles.views, users.name AS user_name, users.family AS user_family, users.email as user_email, article_categories.name AS category_name, article_categories.translated_name AS category_translated_name").Where(&amodels.Article{TitleUrl: titleUrl}).Joins("INNER JOIN users ON articles.user_id=users.id").Joins("INNER JOIN article_categories ON articles.article_category_id=article_categories.id").Scan(&article).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.JSONResponse(c, 404, fiber.Map{"error": "Article not found"})
		}
		return utils.ServerError(c, err)
	}

	if article.Title == "" {
		return utils.JSONResponse(c, 404, fiber.Map{"error": "Article not found"})
	}

	// have view token
	// viewToken := c.Cookies("Vtoken", "")
	// if viewToken == "" {
	// 	// add view token
	// 	viewToken = utilstoken.CreateRandomToken()
	// 	viewToken = utils.Hash(viewToken)
	// 	c.Cookie(&fiber.Cookie{
	// 		Name:    "Vtoken",
	// 		Value:   viewToken,
	// 		Expires: time.Now().Add(24 * time.Hour),
	// 	})
	// }
	sess, err := session.GlobalSession.Get(c)
	if err==nil{
			AddViewArticle(article.ID, article.Views, sess)
	}

	article.ID = 0

	return utils.JSONResponse(c, 200, fiber.Map{"data": article})

}

// GET, /?title=x or /?tags=x|y
func SearchArticle(c *fiber.Ctx) error {
	q := c.Queries()

	var isReturnData bool
	var articles []amodels.OArticleTitle

	// search with title
	title, ok := q["title"]
	if ok {
		// check title
		if err := utils.IsNotInvalidCharacter(title); err != nil {
			return utils.JSONResponse(c, 400, fiber.Map{"error": err})
		}
		// search by title
		if err := database.DB.Model(&amodels.Article{}).Where("title like ?", "%"+title+"%").Select("articles.title, articles.title_url, articles.image_path, articles.short_desc, users.name AS user_name, users.family AS user_family").Joins("INNER JOIN users ON articles.user_id=users.id").Scan(&articles).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return utils.JSONResponse(c, 404, fiber.Map{"error": "no Article found"})
			}
			return utils.ServerError(c, err)
		}

		isReturnData = true

	}

	// search with tag
	tags, ok := q["tags"]
	if ok {
		// check tags
		if err := utils.IsNotInvalidCharacter(tags); err != nil {
			return utils.JSONResponse(c, 400, fiber.Map{"error": err.Error()})
		}
		// search by tags
		stags := strings.Split(tags, "|")

		// have two tag
		if len(stags) >= 2 {

			if err := database.DB.Model(&amodels.Article{}).Where("tags like ? and tags like %?%", "%"+stags[0]+"%", "%"+stags[1]+"%").Select("articles.title, articles.title_url, articles.image_path, articles.short_desc, users.name AS user_name, users.family AS user_family").Joins("INNER JOIN users ON articles.user_id=users.id").Scan(&articles).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return utils.JSONResponse(c, 404, fiber.Map{"error": "no Article found"})
				}
				return utils.ServerError(c, err)
			}

		}
		if len(articles) == 0 && len(stags) == 1 {
			if err := database.DB.Model(&amodels.Article{}).Where("tags=?", "%"+stags[0]+"%").Select("articles.title, articles.title_url, articles.image_path, articles.short_desc, users.name AS user_name, users.family AS user_family").Joins("INNER JOIN users ON articles.user_id=users.id").Scan(&articles).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return utils.JSONResponse(c, 404, fiber.Map{"error": "no Article found"})
				}
				return utils.ServerError(c, err)
			}
		}

		isReturnData = true
	}

	if isReturnData {
		if len(articles) == 0 {
			return utils.JSONResponse(c, 404, fiber.Map{"error": "no article found"})
		}

		err := hutils.ConvertArticleStringTimesForOutput(articles)
		if err != nil {
			return utils.ServerError(c, errors.New("cant convert article database times to time: "+err.Error()))

		}
		return utils.JSONResponse(c, 200, fiber.Map{"data": articles})

	}

	return utils.JSONResponse(c, 400, fiber.Map{"error": "invalid query"})
}

//////////////////  admin  //////////////////////////////

// POST, Admin Required, /:articleTitleUrl
func CreateArticle(c *fiber.Ctx) error {
	
	// get article from request body
	var al amodels.IArticle
	if err := c.BodyParser(&al); err != nil {
		return utils.JSONResponse(c, 400, fiber.Map{"error": "invalid json"})
	}

	// check check validation
	if err := al.Check(); err != nil {
		return utils.JSONResponse(c, 400, fiber.Map{"error": err.Error()})
	}

	titleUrl := c.Params("articleTitleUrl")
	if titleUrl==""{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid titleUrl"})
	}

	// check article is exist
	if err := database.DB.First(&amodels.Article{}, &amodels.Article{TitleUrl: titleUrl}).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return utils.ServerError(c, err)
		}
	} else {
		return utils.JSONResponse(c, 400, fiber.Map{"error": "the article title already exist"})
	}

	user, err := session.GetUserFromSession(c)
	if err != nil {
		return utils.JSONResponse(c, 401, fiber.Map{"error": "Auth Required"})
	}
	

	// create Article and fill it
	var article amodels.Article
	article.Fill(al)
	article.UserID = user.ID
	article.TimeCreated = time.Now()
	article.TimeModified = time.Now()
	article.TitleUrl = titleUrl //utils.ConvertToUrl(article.Title)

	if err := database.DB.Create(&article).Error; err != nil {
		if err.Error()[:18] == "Error 1216 (23000)" {
			return utils.JSONResponse(c, 400, fiber.Map{"error": "Category not found"})
		}
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 201, fiber.Map{"msg": "successfully created"})
}

// POST, Auth Required, Admin Required, /:id
// function getting article id and a image
func UploadArticleImage(c *fiber.Ctx) error {

	// id := utils.GetIDFromParams(c, "articleTitleUrl")
	// if id == 0 {
	// 	return utils.JSONResponse(c, 400, fiber.Map{"error": "the articleId didn't send"})
	// }

	// // check Article is exist
	// var article amodels.Article
	// if err := database.DB.First(&article, &amodels.Article{ID: id}).Error; err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		return utils.JSONResponse(c, 404, fiber.Map{"error": "Article not found"})
	// 	}
	// 	return utils.ServerError(c, err)
	// }

		article := c.Locals("article").(amodels.Article)

	file, err := c.FormFile("i")
	if err != nil {
		return utils.ServerError(c, err)
	}

	// delete last image if exist
	if article.ImagePath != ""{
		if strings.Contains(article.ImagePath, "*") {
			return utils.ServerError(c, errors.New("one star is exist in image path. maybe hacker do this"))
		}
		if article.ImagePath != "" {
			err := s3.DeleteObject(fmt.Sprintf(".%s", article.ImagePath))
			if err != nil {
				return utils.ServerError(c, err)
			}
		}
	}

	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	image := fmt.Sprintf("%s.%s", filename, fileExt)

	fl, err := file.Open()
	if err != nil {
		return utils.ServerError(c, err)
	}
	defer fl.Close()

	imageURL := path.Join(database.Settings.ImageSaveUrl, fmt.Sprintf("/article/%s", image))

	err = s3.PutObject(fl, imageURL)
	// err = c.SaveFile(file, fmt.Sprintf("../frontend/dist/images/%s", image))

	if err != nil {
		return utils.ServerError(c, err)
	}

	imageURL = path.Join(database.Settings.ImageUrlSubdomain, imageURL)


	if err = database.DB.Model(&amodels.Article{}).Where(&amodels.Article{ID: article.ID}).Update("image_path", imageURL).Error; err != nil {
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg": "image added"})
}

// PUT, Auth Required, Admin Required, /:articleTitleUrl
func EditArticle(c *fiber.Ctx) error {
	// // get id form params
	// id := utils.GetIntFromParams(c, "articleId")
	// if id==0{
	// 	return utils.JSONResponse(c, 400, fiber.Map{"error":"the articleId didn't send"})
	// }

	// // check article is exist
	// var article amodels.Article
	// if err:= database.DB.First(&article, &amodels.Article{ID: id}).Error; err!=nil{
	// 	if err==gorm.ErrRecordNotFound{
	// 		return utils.JSONResponse(c, 404, fiber.Map{"error":"Article not found"})
	// 	}
	// 	return utils.ServerError(c, err)
	// }

	article := c.Locals("article").(amodels.Article)

	// get article from body
	var al amodels.IArticle
	if err := c.BodyParser(&al); err != nil {
		return utils.JSONResponse(c, 400, fiber.Map{"error": "invalid json"})
	}

	// // check article is exist
	// if err := database.DB.First(&amodels.Article{}, &amodels.Article{TitleUrl: article.TitleUrl}).Error; err != nil {
	// 	if err != gorm.ErrRecordNotFound {
	// 		return utils.ServerError(c, err)
	// 	}
	// } else {
	// 	return utils.JSONResponse(c, 400, fiber.Map{"error": "the article title already exist"})
	// }

	// check Article validation
	if err := al.Check(); err != nil {
		return utils.JSONResponse(c, 400, fiber.Map{"error": err.Error()})
	}

	// fill the Article
	article.Fill(al)
	article.TimeModified = time.Now()

	// modify Article in database
	if err := database.DB.Updates(&article).Error; err != nil {
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"msg": "successfully modified"})
}

// DELETE, /:articleTitleUrl
func DeleteArticle(c *fiber.Ctx) error {

	article := c.Locals("article").(amodels.Article)

	if err := database.DB.Delete(&article).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.JSONResponse(c, 404, fiber.Map{"error": "Article not found"})
		}
		return utils.ServerError(c, err)
	}
	if strings.Contains(article.ImagePath, "*") {
		return utils.ServerError(c, errors.New("one star is exist in image path. maybe hacker do this"))
	}
	if article.ImagePath != "" {
		err := s3.DeleteObject(fmt.Sprintf(".%s", article.ImagePath))
		if err != nil {
			return utils.ServerError(c, err)
		}
	}
	return utils.JSONResponse(c, 200, fiber.Map{"msg": "successfuly deleted"})
}

func AddViewArticle(articleID uint64, views uint64, sess *main_session.Session) {
	// if err:= database.DB.Model(&tmodels.Temp{}).First(&tmodels.Temp{}, &tmodels.Temp{String1f: titleUrl, String2f: vtoken}).Error; err!=nil{
	// 	if err==gorm.ErrRecordNotFound{
	// 		// add view to article
	// 		views++
	// 		if err2:= database.DB.Model(&amodels.Article{}).Where(&amodels.Article{TitleUrl: titleUrl}).Update("views", views).Error;err!=nil{
	// 			utils.LogError(err2)
	// 			return
	// 		}
	// 		// add titleUrl and vtoken to temp table
	// 		ctemp := tmodels.Temp{String1f: titleUrl, String2f: vtoken, EndTime: time.Now().Add(24*time.Hour)}
	// 		if err2:= database.DB.Create(&ctemp).Error;err!=nil{
	// 			utils.LogError(err2)
	// 			return
	// 		}

	// 		// all done
	// 		return
	// 	}

	// server error
	// utils.LogError(err)
	// return
	// }

	defer sess.Save()


	mustIncrease := false
	ar := sess.Get("articleIDs")
	if ar == nil {
		// user don't view until now
		sess.Set("articleIDs", []uint64{articleID})
		mustIncrease = true

	} else {
		arr := ar.([]uint64)

		if slices.Contains(arr, articleID) {
			return
		} else {
			arr = append(arr, articleID)
			sess.Set("articleIDs", arr)
			mustIncrease = true
		}

	}

	if mustIncrease {
		views++
		go func(){
			if err := database.DB.Model(&amodels.Article{}).Where(&amodels.Article{ID: articleID}).Update("views", views).Error; err != nil {
			utils.LogError(err)
			return
		}}()

	}
}


