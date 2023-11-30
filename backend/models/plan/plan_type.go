package plan

import (
	"time"
	"errors"

	"lovelcode/utils"
)




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

