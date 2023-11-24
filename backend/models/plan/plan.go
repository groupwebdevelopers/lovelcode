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

type OFeature struct{
	FID uint64 `json:"id"`
	Name string `json:"name"`
	Value string `json:"value"`
	// Price uint32 `json:"price"`
	IsHave bool `json:"isHave"`
	IsFeatured bool `json:"isFeatured"`
}


type PlanType struct{
	ID uint64 `gorm:"primaryKey"`
	Name string `gorm:"not null,size:100,unique"`
	TranslatedName string
	// ImagePath string `gorm:"size:200"`
	Order int `gorm:"not null`

	TimeCreated time.Time `gorm:"not null"`
	TimeModified time.Time `gorm:"not null"`
}

type IPlanType struct{
	Name string `json:"name"`
	TranslatedName string `json:"translatedName"`
	Order int `json:"order"`
}

type OPlanType struct{
	ID uint64 `json:"id"`
	Name string `json:"name"`
	TranslatedName string `json:"translatedName"`
	Order int `json:"order"`
}

type OrderPlan struct{
	ID uint64 `gorm:"primaryKey"`
	Name string `gorm:"not null, size:100"`
	Phone uint64 `gorm:"not null"`
	Email string `gorm:"not null"`
	Type string `gorm:"not null"`
	Storage int64 `gorm:"not null"`
	Meet string `gorm:"size:200"`
	Desc string `gorm:"size:500"`

	TimeCreated time.Time
	TimeModified time.Time
}


type IOrderPlan struct{
	Name string `json:"name"`
	Phone uint64 `json:"phone"`
	Email string `json:"email"`
	Type string `json:"typeOfWebSite"`
	Storage int64 `json:"storage"`
	Meet string `json:"meet"`
	Desc string `json:"desc"`
}


type OOrderPlan struct{
	ID uint64 `json:"id"`
	Name string `json:"name"`
	Phone uint64 `json:"phone"`
	Email string `json:"email"`
	Type string `json:"typeOfWebSite"`
	Storage int64 `json:"storage"`
	Meet string `json:"meet"`
	Desc string `json:"desc"`

	TimeCreated time.Time `json:"timeCreated"`
	TimeModified time.Time `json:"timeModified"`
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

func (i *IPlanType) Check() error{
	if err:= utils.IsNotInvalidCharacter(i.Name); err!=nil{
		return errors.New("invalid name: "+err.Error())
	}
	if err:=utils.IsNotInvalidCharacter(i.TranslatedName); err!=nil{
		return errors.New("invalid translated name:" +err.Error())
	}

	return nil
}

func (t *PlanType) Fill(i *IPlanType){
	t.Name = i.Name
	t.TranslatedName = i.TranslatedName
	t.Order = i.Order
}

func (i *IOrderPlan) Check() error{
	if err:= utils.IsNotInvalidCharacter(i.Name); err!=nil{
		return errors.New("invalid name: "+err.Error())
	}
	if i.Phone < 9000000000{
		return errors.New("small phone")
	}
	if err:= utils.IsNotInvalidCharacter(i.Type); err!=nil{
		return errors.New("invalid type: "+err.Error())
	}
	if err:= utils.IsNotInvalidCharacter(i.Meet); err!=nil{
		return errors.New("invalid meet: "+err.Error())
	}
	if err:= utils.IsNotInvalidCharacter(i.Desc); err!=nil{
		return errors.New("invalid description: "+err.Error())
	}
	if err:=utils.CheckEmail(i.Email); err!=nil{
		return errors.New("invalid email")
	}
	if i.Storage <= 0 {
		return errors.New("invalid storage")
	}

	return nil
}

func (o *OrderPlan) Fill(i *IOrderPlan) {
	o.Name = i.Name
	o.Email = i.Email
	o.Phone = i.Phone
	o.Meet = i.Meet
	o.Desc = i.Desc
	o.Type = i.Type
	o.Storage = i.Storage
}