package models

import (
	"time"
	"errors"

	"lovelcode/utils"
)

type Plan struct{
	ID uint64 `gorm:"primaryKey"`
	Name string `gorm:"not null"`
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
	Price uint32
	IsHave bool `gorm:"not null"` // the plan is have this feature

	TimeCreated time.Time `gorm:"not null"`
	TimeModified time.Time `gorm:"not null"`
}

// create and edit feature
type IFeature struct{
	// PlanID uint64 `json:"planID"`
	Name string `json:"name"`
	Value string `json:"value"`
	Price int `json:"price"`
	IsHave bool `json:"isHave"`
}

type IPlan struct{
	Name string `json:"name"`
	Price uint32 `json:"price"`
	Type string `json:"type"`
}

type OPlan struct{
	Name string `json:"name"`
	Price uint32 `json:"price"`
	ImagePath string `json:"imagePath"`
	Type string `json:"type"`
}

type OFeature struct{
	PlanID uint64 `json:"planID"`
	Name string `json:"name"`
	Value string `json:"value"`
	Price uint32 `json:"price"`
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
	if f.Price < 0{
		return errors.New("invalid price")
	}
	// if f.PlanID == 0{
	// 	return errors.New("invalid planID")
	// }
	return nil
}

func (f *Feature) FillWithCEFeature(i IFeature){
	// f.PlanID = ce.PlanID
	f.Name = i.Name
	f.Value = i.Value
	f.Price = uint32(i.Price)
	f.IsHave = i.IsHave
} 

func (p *IPlan) Check() error{
	if err:=utils.IsNotInvalidCharacter(p.Name); err!=nil{
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
	p.Name = ce.Name
	p.Price = ce.Price
	p.Type = ce.Type
}

func (o *OPlan) FillWithPlan(p Plan){
	o.Name = p.Name
	o.Price = p.Price
	o.ImagePath = p.ImagePath
	o.Type = p.Type
}

func (o *OFeature) FillWithFeature(f Feature){
	o.Name = f.Name
	o.IsHave = f.IsHave
	o.PlanID = f.PlanID
	o.Price = f.Price
	o.Value = f.Value
}