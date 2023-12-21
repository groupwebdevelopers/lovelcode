package article

import (
	"errors"
	"time"
	"strconv"

	umodels "lovelcode/models/user"
	"lovelcode/utils"
)

type Article struct{
	ID uint64 `gorm:"primaryKey"`
	UserID uint64
	User umodels.User
	ArticleCategoryID uint64 
	ArticleCategory ArticleCategory
	Title string `gorm:"not null,size:200"`
	TitleUrl string `gorm:"not null,size:1000,unique"`
	Body string `gorm:"not null,size:10000"`
	Tags string `gorm:"size:200"` // split with |
	ShortDesc string `gorm:"not null,size:200"` // short description
	ImagePath string `gorm:"size:200"`
	Views uint64
	IsFeatured bool
	Likes uint64

	TimeCreated time.Time
	TimeModified time.Time
}

type IArticle struct{
	Title string `json:"title"`
	Body string `json:"body"`
	Tags string `json:"tags"`
	ShortDesc string `json:"shortDesc"`
	IsFeatured bool `json:"isFeatured"`
	CategoryID uint64 `json:"categoryId"`
}

// output article for user
type OArticle struct{
	ID uint64 `json:"i"`
	Title string `json:"title"`
	Body string `json:"body"`
	Tags string `json:"tags"`
	ShortDesc string `json:"shortDesc"`
	ImagePath string `json:"imagePath"`
	TimeCreated time.Time `json:"timeCreated"`
	TimeModified time.Time `json:"timeModified"`
	Views uint64
	Likes uint64

	UserName string `json:"userName"`
	UserFamily string `json:"userFamily"`
	UserEmail string `json:"userEmail"`

	CategoryName string `json:"categoryName"`
	CategoryTranslatedName string `json:"categoryTranslatedName"`
}

type OArticleTitle struct{
	Title string `json:"title"`
	TitleUrl string `json:"titleUrl"`
	ShortDesc string `json:"shortDesc"`
	ImagePath string `json:"imagePath"`
	TimeCreated string `json:"timeCreated"`
	TimeModified string `json:"timeModified"`
	Views uint64
	Likes uint64

	UserName string `json:"userName"`
	UserFamily string `json:"userFamily"`
}

func (a *Article) Fill(i IArticle){
	a.Title = i.Title
	a.Body = i.Body
	a.Tags = i.Tags
	a.ShortDesc = i.ShortDesc
	a.IsFeatured = i.IsFeatured
	a.ArticleCategoryID = i.CategoryID
}

func (a *IArticle) Check() error{
	if err:= utils.IsNotInvalidCharacter(a.Title, "/"); err!=nil{
		return errors.New("invalid titile:" + err.Error())
	}
	// body and short desctiption must converted to base64
	if err:= utils.IsNotInvalidCharacter(a.Tags, "/!@#$%^&*()_+-=][{}:?><.,] "); err!=nil{
		return errors.New("invalid tags:"+err.Error())
	}

	if len(a.Body) > 10000{
		return errors.New("too long body")
	}
	if len(a.Tags) > 200{
		return errors.New("too long tags")
	}
	if len(a.ShortDesc) > 200{
		return errors.New("too long shertDesc ("+ strconv.Itoa(len(a.ShortDesc))+">120)")
	}
	
	return nil
}
