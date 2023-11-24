package models

import (
	"time"
	"errors"

	"lovelcode/utils"
)




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