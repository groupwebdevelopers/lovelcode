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
	CommentAnswerID uint64

	Body string `gorm:"size:100"`

	TimeCreated time.Time
	TimeModified time.Time
}

type IComment struct{
	Body string `json:"body"`
	CommentAnswerID uint64 `json:"commentAnswerID"`
}

type OComment struct{
	Body string `json:"body"`
	Name string  `json:"name"`
	Family string `json:"family"`
	CommentAnswerID uint64 `json:"commentAnswerID"`
}

func (c *Comment) Fill(i *IComment){
	c.Body = i.Body
	c.CommentAnswerID = i.CommentAnswerID
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