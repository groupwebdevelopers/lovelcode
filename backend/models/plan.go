package models

import (
	"time"
)

type Plan struct{
	ID uint64 `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Price uint32 `gorm:"not null"`
	// Features string `gorm:"not null"`
	ImagePath string
}

type Feature struct{
	ID uint64 `gorm:"primaryKey"`
	PlanID uint64
	Plan Plan
	Name string `gorm:"not null"`
	Description string
	Price uint32
	IsHave bool `gorm:"not null"` // the plan is have this feature
}

// create and edit feature
type CEFeature struct{
	PlanID uint64 `json:"planID"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price uint32 `json:"price"`
	IsHave bool `json:"isHave"`
}

type CEPlan struct{
	Name string `json:"name"`
	Price uint32 `json:"price"`
	// image must sent
}