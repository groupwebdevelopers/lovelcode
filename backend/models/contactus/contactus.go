package contactus

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
	Number uint64

	IsSeen bool
	
	TimeCreated time.Time
	// TimeModified time.Time
}

type IContactUs struct{
	Title string `json:"title"`
	Body string `json:"body"`
	Email string `json:"email"`
	Number uint64 `json:"number"`
}


type OContactUs struct{
	ID uint64 `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Email string `json:"email"`
	Number uint64 `json:"number"`
	IsSeen bool `json:"isSeen"`
	TimeCreated time.Time
}

func (i *IContactUs) Check() error{
	if err:=utils.IsNotInvalidCharacter(i.Title, "/"); err!=nil{
		return errors.New("invalid title:" +err.Error())
	}
	if err:=utils.IsNotInvalidCharacter(i.Body); err!=nil{
		return errors.New("invalid body:" +err.Error())
	}

	if len(i.Title) > 100{
		return errors.New("too long title")
	}
	if len(i.Body) > 400{
		return errors.New("too long body")
	}
	if i.Number < 9000000000 || i.Number > 9999999999{
		return errors.New("invalid number")
	}

	return nil
}

func (c *ContactUs) Fill(i *IContactUs){
	c.Title = i.Title
	c.Body = i.Body
	c.Email = i.Email
	c.Number = i.Number
}