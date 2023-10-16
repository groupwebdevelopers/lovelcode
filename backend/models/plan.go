package models

import (
	"time"
	"errors"

)

type Plan struct{
	ID uint64 `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Price uint32 `gorm:"not null"`
	ImagePath string

	TimeCreated time.Time `gorm:"not null"`
	TimeModified time.Time `gorm:"not null"`
}

type Feature struct{
	ID uint64 `gorm:"primaryKey"`
	PlanID uint64
	Plan Plan
	Name string `gorm:"not null"`
	Description string
	Price uint32
	IsHave bool `gorm:"not null"` // the plan is have this feature

	TimeCreated time.Time `gorm:"not null"`
	TimeModified time.Time `gorm:"not null"`
}

// create and edit feature
type CEFeature struct{
	// PlanID uint64 `json:"planID"`
	Name string `json:"name"`
	// Description string `json:"description"`
	Price uint32 `json:"price"`
	IsHave bool `json:"isHave"`
}

type CEPlan struct{
	Name string `json:"name"`
	Price uint32 `json:"price"`
	
}

type CEPlanAndFeatures struct{
	Plan CEPlan `json:"plan"`
	Features []CEFeature `json:"features`
}

func (f *CEFeature) Check() error{
	if f.Name == ""{
		return errors.New("empty name")
	}
	// if f.PlanID == 0{
	// 	return errors.New("invalid planID")
	// }
	return nil
}

func (f *Feature) FillWithCEFeature(ce CEFeature){
	// f.PlanID = ce.PlanID
	f.Name = ce.Name
	// f.Description = ce.Description
	f.Price = ce.Price
	f.IsHave = ce.IsHave
} 

func (p *CEPlan) Check() error{
	if p.Name == ""{
		return errors.New("empty name")
	}
	if p.Price == 0{
		return errors.New("invalid price")
	}
	return nil
}

func (p *Plan) FillWithCEPlan(ce CEPlan) {
	p.Name = ce.Name
	p.Price = ce.Price
}