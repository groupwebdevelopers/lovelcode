package customer

import (
	"errors"
	"time"
	"lovelcode/utils"
)

type Customer struct{
	ID uint64 `gorm:"primaryKey"`
	Name string `gorm:"size:100,not null"`
	SiteUrl string `gorm:"size:100,not null"`
	ImagePath string `gorm:"size:200"`
	IsFeatured bool

	TimeCreated time.Time
	TimeModified time.Time
}


type ICustomer struct{
	Name string `json:"name"`
	SiteUrl string `json:"siteUrl"`
	IsFeatured bool `json:"isFeatured"`
}


type OCustomer struct{
	ID uint64 `json:"id"`
	Name string `json:"name"`
	SiteUrl string `json:"siteUrl"`
	ImagePath string `json:"imagePath"`
	IsFeatured bool `json:"isFeatured"`
}


func (i *ICustomer) Check() error{
	if err:= utils.IsNotInvalidCharacter(i.Name); err!=nil{
		return errors.New("invalid name: " + err.Error())
	}
	if err:= utils.IsNotInvalidCharacter(i.SiteUrl); err!=nil{
		return errors.New("invalid section: " + err.Error())
	}
	if len(i.Name) > 100{
		return errors.New("too long name")
	}
	if len(i.SiteUrl) > 100{
		return errors.New("too long siteUrl")
	}

	return nil
}

func (m *Customer) Fill(i *ICustomer){
	m.Name = i.Name
	m.SiteUrl = i.SiteUrl
	m.IsFeatured = i.IsFeatured
}
