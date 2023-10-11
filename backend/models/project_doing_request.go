package models

import (
	"time"
)

type ProjectDoingRequest struct{
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
	SuggestedPrice int `json:"suggestedprice`
}

func (pdr *ProjectDoingRequest) FillWithCEPDR(ce CEPDR){
	pdr.Title = ce.Title
	pdr.Description = ce.Description
	if ce.SuggestedPrice >= 0{
		pdr.SuggestedPrice = uint(ce.SuggestedPrice)
	}
}

func (c CEPDR) Check() error{
	if err:=utils.IsJustLetter(c.Title, "-!$%&*()_+= .,?");err!=nil{
		return err
	}
	if c.Description == ""{
		return errors.New("empty desciption")
	}
	if c. SuggestedPrice < 0{
		return errors.New("negetive suggested price")
	}
	return nil
}