package models

import (
	"time"
	"errors"

	"gorm.io/gorm"

	"lovelcode/utils"
)

type ProjectDoingRequest struct{
	gorm.Model
	ID uint64 `gorm:"primaryKey"`
	UserID uint64 `gorm:"not null"`
	User User `gorm:"not null"`
	Title string `gorm:"not null"`
	Description string `gorm:"not null"`
	SuggestedPrice uint 
	TimeCreated time.Time `gorm:"not null"`
	TimeModified time.Time `gorm:"not null"`
}

// create or edit project doing request
type CEPDR struct{
	Title string `json:"title"`
	Description string `json:"description"`
	SuggestedPrice int `json:"suggestedPrice"`
}

func (pdr *ProjectDoingRequest) FillWithCEPDR(ce CEPDR){
	pdr.Title = ce.Title
	pdr.Description = ce.Description
	if ce.SuggestedPrice >= 0{
		pdr.SuggestedPrice = uint(ce.SuggestedPrice)
	}
}

func (c CEPDR) Check() error{
	if err:=utils.IsNotInvalidCharacter(c.Title);err!=nil{
		return errors.New("invalid title"+err.Error())
	}
	if err:= utils.IsNotInvalidCharacter(c.Description); err!=nil{
		return errors.New("invalid desciption:"+err.Error())
	}
	if c. SuggestedPrice < 0{
		return errors.New("negetive suggested price")
	}
	return nil
}