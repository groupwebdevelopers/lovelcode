package models

import (
	"errors"
	"lovelcode/utils"
	"time"
)

type ContactUs struct{
	ID uint64 `gorm:"primaryKey"`
	UserID uint64
	User User
	Title string `gorm:"size:100,not null"`
	TitleUrl string `gorm:"size:100,not null,unique"`
	Body string `gorm:"size:400,not null"`
	TimeCreated time.Time
	TimeModified time.Time
}

type IContactUs struct{
	Title string `json:"title"`
	Body string `json:"body"`
}

type OContactUs struct{
	Title string `json:"title"`
	Body string `json:"body"`
	TitleUrl string `json:"titleUrl"`
}

func (i *IContactUs) Check() error{
	if err:=utils.IsNotInvalidCharacter(i.Title, "/"); err!=nil{
		return errors.New("invalid title:" +i.Check().Error())
	}
	if err:=utils.IsNotInvalidCharacter(i.Body); err!=nil{
		return errors.New("invalid body:" +i.Check().Error())
	}

	if len(i.Title) > 100{
		return errors.New("too long title")
	}
	if len(i.Body) > 400{
		return errors.New("too long body")
	}

	return nil
}

func (c *ContactUs) Fill(i *IContactUs){
	c.Title = i.Title
	c.Body = i.Body
}