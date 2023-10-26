package models

import (
	"errors"
	"lovelcode/utils"
	"time"
)

type Comment struct{
	ID uint64 `gorm:"primaryKey"`
	UserID uint64
	User User
	ArticleID uint64
	Article Article

	Body string `gorm:"size:100"`

	TimeCreated time.Time
	TimeModified time.Time
}

type IComment struct{
	Body string `json:"body"`
}

type OComment struct{
	Body string `json:"body"`
	Name string  `json:"name"`
	Family string `json:"family"`
}

func (c *Comment) Fill(i *IComment){
	c.Body = i.Body
}

func (i *IComment) Check() error{
	if err:= utils.IsNotInvalidCharacter(i.Body); err!=nil{
		return errors.New("invalid body: "+ err.Error())
	}

	if len(i.Body) > 100{
		return errors.New("too long body")
	}
	return nil
}