package models

import (
	"time"
	"errors"

	"lovelcode/utils"
)

type Plan struct{
	ID uint64 `gorm:"primaryKey"`
	Title string `gorm:"not null,size:100,unique"`
	Price uint32 `gorm:"not null"`
	ImagePath string `gorm:"size:200"`
	Type string `gorm:"size:50"`
	IsFeatured bool
	IsCompare bool

	TimeCreated time.Time `gorm:"not null"`
	TimeModified time.Time `gorm:"not null"`
}

type IPlan struct{
	Title string `json:"title"`
	Price uint32 `json:"price"`
	Type string `json:"type"`
	IsFeatured bool `json:"isFeatured"`
	IsCompare bool `json:"isCompare"`
}

type OPlan struct{
	ID uint64 `json:"id"`
	Title string `json:"title"`
	Price uint32 `json:"price"`
	ImagePath string `json:"imagePath"`
	Type string `json:"type"`
	IsFeatured bool `json:"isFeatured"`
	IsCompare bool `json:"isCompare"`
}


func (p *IPlan) Check() error{
	if err:=utils.IsNotInvalidCharacter(p.Title); err!=nil{
		return errors.New("invalid plan name:"+err.Error())
	}
	if err:=utils.IsNotInvalidCharacter(p.Type); err!=nil{
		return errors.New("invalid plan type:"+err.Error())
	}
	
	if len(p.Title) > 100{
		return errors.New("too long title")
	}
	if len(p.Type) > 50{
		return errors.New("too long type")
	}

	if p.Price <= 0{
		return errors.New("invalid price")
	}
	return nil
}


func (p *Plan) Fill(ce *IPlan) {
	p.Title = ce.Title
	p.Price = ce.Price
	p.Type = ce.Type
	p.IsFeatured = ce.IsFeatured
	p.IsCompare = ce.IsCompare
}