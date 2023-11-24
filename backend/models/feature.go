package models

import (
	"time"
	"errors"

	"lovelcode/utils"
)




type Feature struct{
	ID uint64 `gorm:"primaryKey"`
	PlanID uint64
	Plan Plan
	Name string `gorm:"not null,size:200"`
	Value string `gorm:"size:100"`
	// Price uint32
	IsHave bool `gorm:"not null"` // the plan is have this feature
	IsFeatured bool

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
	IsFeatured bool `json:"isFeatured"`
}


type OFeature struct{
	FID uint64 `json:"id"`
	Name string `json:"name"`
	Value string `json:"value"`
	// Price uint32 `json:"price"`
	IsHave bool `json:"isHave"`
	IsFeatured bool `json:"isFeatured"`
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
	if len(f.Name) > 200{
		return errors.New("too long name")
	}
	if len(f.Value) > 100{
		return errors.New("too long value")
	}
	// if f.Price < 0{
		// return errors.New("invalid price")
		// }
		return nil
}

func (f *Feature) Fill(i *IFeature){
	// f.ID = i.ID
	f.Name = i.Name
	f.Value = i.Value
	// f.Price = uint32(i.Price)
	f.IsHave = i.IsHave
	f.IsFeatured = i.IsFeatured
} 

