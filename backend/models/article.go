package models

import (
	"errors"
	"lovelcode/utils"
	"time"
)

type Article struct{
	ID uint64 `gorm:"primaryKey"`
	UserID uint64
	User User
	ArticleCategoryID uint64 
	ArticleCategory ArticleCategory
	Title string `gorm:"not null,size:100"`
	TitleUrl string `gorm:"not null,size:100,unique"`
	Body string `gorm:"not null,size:10000"`
	Tags string `gorm:"size:200"` // split with |
	ShortDesc string `gorm:"not null,size:100"` // short description
	ImagePath string `gorm:"size:200"`
	Views uint64
	IsFeatured bool

	TimeCreated time.Time
	TimeModified time.Time
}

type IArticle struct{
	Title string `json:"title"`
	Body string `json:"body"`
	Tags string `json:"tags"`
	ShortDesc string `json:"shortDesc"`
	IsFeatured bool `json:"isFeatured"`
}

// output article for user
type OArticle struct{
	Title string `json:"title"`
	Body string `json:"body"`
	Tags string `json:"tags"`
	ShortDesc string `json:"shortDesc"`
	ImagePath string `json:"imagePath"`
	TimeCreated time.Time `json:"timeCreated"`
	TimeModified time.Time `json:"timeModified"`
	Views uint64

	UserName string `json:"userName"`
	UserFamily string `json:"userFamily"`
	UserEmail string `json:"userEmail"`

	CategoryName string `json:"categoryName"`
	CategoryTranslatedName string `json:"categoryTranslatedName"`
}

type ArticleCategory struct{
	ID uint64 `gorm:"primaryKey"`
	MainCategory string `gorm:"size:100,not null"`
	TranslatedMainCategory string `json:"size:100"`
	Name string `gorm:"size:100,not null,unique"`
	TranslatedName string `gorm:"size:100"`
	Description string `gorm:"size:400"`
	MainOrder int //`gorm:"not null"`
	Order int `gorm:"not null"`
}

// example
// test articles (MainOrder=1)
//		test 1 (Order=1)
// 		test 2 (Order=2)
// SEO articles (MainOrder=2)
//		seo1 (Order=1)
//		seo2 (Order=2)

type IArticleCategory struct{
	MainCategory string `json:"mainCategory"`
	TranslatedMainCategory string `json:"translatedMainCategory"`
	Name string `json:"name"`
	TranslatedName string `json:"translatedName"`
	Description string `json:"description"`
	MainOrder int `json:"mainOrder"`
	Order int `json:"order"`
}

type OArticleCategory struct{
	ID uint64 `json:"id"`
	MainCategory string `json:"mainCategory"`
	TranslatedMainCategory string `json:"translatedMainCategory"`
	Name string `json:"name"`
	TranslatedName string `json:"translatedName"`
	Description string `json:"description"`
	MainOrder int `json:"mainOrder"`
	Order int `json:"order"`
}

// output article for admin
// type AArticle struct{
// 	Title string `json:"title"`
// 	Body string `json:"body"`
// 	Tags string `json:"tags"`
// 	ShortDesc string `json:"shortDesc"`
// }

func (a *Article) Fill(i IArticle){
	a.Title = i.Title
	a.Body = i.Body
	a.Tags = i.Tags
	a.ShortDesc = i.ShortDesc
	a.IsFeatured = i.IsFeatured
}

func (a *IArticle) Check() error{
	if err:= utils.IsNotInvalidCharacter(a.Title, "/"); err!=nil{
		return errors.New("invalid titile:" + err.Error())
	}
	// body and short desctiption must converted to base64
	if err:= utils.IsNotInvalidCharacter(a.Tags, "/!@#$%^&*()_+-=][{}:?><.,]"); err!=nil{
		return errors.New("invalid tags:"+err.Error())
	}

	if len(a.Body) > 10000{
		return errors.New("too long body")
	}
	if len(a.Tags) > 200{
		return errors.New("too long tags")
	}
	if len(a.ShortDesc) > 100{
		return errors.New("too long shertDesc")
	}
	
	return nil
}

func (c *ArticleCategory) Fill(i *IArticleCategory){
	c.MainCategory = i.MainCategory
	c.TranslatedMainCategory = i.TranslatedMainCategory
	c.Name = i.Name
	c.TranslatedName = i.TranslatedName
	c.Description = i.Description
	c.Order = i.Order
	c.MainOrder = i.MainOrder

}

func (i *IArticleCategory) Check() error{
	if err:=utils.IsJustLetter(i.MainCategory, " "); err!=nil{
		return errors.New("invalid mainCategory field: "+err.Error())
	}
	if err:=utils.IsNotInvalidCharacter(i.TranslatedMainCategory); err!=nil{
		return errors.New("invalid translatedMainCategory field: "+err.Error())
	}
	if err:=utils.IsJustLetter(i.Name, " "); err!=nil{
		return errors.New("invalid name field: "+err.Error())
	}
	if err:= utils.IsNotInvalidCharacter(i.TranslatedName);err!=nil{
		return errors.New("invalid persian field: "+err.Error())
	}
	if err:= utils.IsNotInvalidCharacter(i.Description);err!=nil{
		return errors.New("invalid desctiption field: "+err.Error())
	}

	return nil
}