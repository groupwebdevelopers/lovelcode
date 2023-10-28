package models

import (
	"errors"
	"lovelcode/utils"
	"time"
)

type ContactUs struct{
	ID uint64 `gorm:"primaryKey"`
	Title string `gorm:"size:100,not null"`
	Body string `gorm:"size:400,not null"`
	Email string `gorm:"size:400,not null"`
	Number uint32

	IsSeen bool
	
	TimeCreated time.Time
	// TimeModified time.Time
}

type IContactUs struct{
	Title string `json:"title"`
	Body string `json:"body"`
	Email string `json:"email"`
	Number uint32 `json:"number"`
}


type OContactUs struct{
	ID uint64 `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	TitleUrl string `json:"titleUrl"`
	Email string `json:"email"`
	Number uint64 `json:"number"`
	IsSeen bool `json:"isSeen"`
	TimeCreated time.Time
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
	c.Email = i.Email
	c.Number = i.Number
}