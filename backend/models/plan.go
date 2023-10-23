package models

import (
	"time"
	"errors"

	"lovelcode/utils"
)

type Plan struct{
	ID uint64 `gorm:"primaryKey"`
	Title string `gorm:"not null"`
	Price uint32 `gorm:"not null"`
	ImagePath string `gorm:"size:200"`
	Type string
	TimeCreated time.Time `gorm:"not null"`
	TimeModified time.Time `gorm:"not null"`
}

type Feature struct{
	ID uint64 `gorm:"primaryKey"`
	PlanID uint64
	Plan Plan
	Name string `gorm:"not null"`
	Value string
	// Price uint32
	IsHave bool `gorm:"not null"` // the plan is have this feature

	TimeCreated time.Time `gorm:"not null"`
	TimeModified time.Time `gorm:"not null"`
}

// create and edit feature
type IFeature struct{
	// PlanID uint64 `json:"planID"`
	Name string `json:"name"`
	Value string `json:"value"`
	// Price int `json:"price"`
	IsHave bool `json:"isHave"`
}

type IPlan struct{
	Title string `json:"title"`
	Price uint32 `json:"price"`
	Type string `json:"type"`
}

type OPlan struct{
	Title string `json:"title"`
	Price uint32 `json:"price"`
	ImagePath string `json:"imagePath"`
	Type string `json:"type"`
}

type OFeature struct{
	Name string `json:"name"`
	Value string `json:"value"`
	// Price uint32 `json:"price"`
	IsHave bool `json:"isHave"`
}

func (f *IFeature) Check() error{
	if err:=utils.IsNotInvalidCharacter(f.Name); err!=nil{
		return errors.New("invalid feature name:"+err.Error())
	}
	if f.Value != ""{
	if err:=utils.IsNotInvalidCharacter(f.Value); err!=nil{
		return errors.New("invalid feature value:"+err.Error())
	}
}
	// if f.Price < 0{
		// return errors.New("invalid price")
	// }
	// if f.PlanID == 0{
	// 	return errors.New("invalid planID")
	// }
	return nil
}

func (f *Feature) FillWithIFeature(i IFeature){
	// f.ID = i.ID
	f.Name = i.Name
	f.Value = i.Value
	// f.Price = uint32(i.Price)
	f.IsHave = i.IsHave
} 

func (p *IPlan) Check() error{
	if err:=utils.IsNotInvalidCharacter(p.Title); err!=nil{
		return errors.New("invalid plan name:"+err.Error())
	}
	if err:=utils.IsNotInvalidCharacter(p.Type); err!=nil{
		return errors.New("invalid plan type:"+err.Error())
	}
	if p.Price <= 0{
		return errors.New("invalid price")
	}
	return nil
}

func (p *Plan) FillWithIPlan(ce IPlan) {
	p.Title = ce.Title
	p.Price = ce.Price
	p.Type = ce.Type
}
