package mainpage

import (
	"errors"

	"lovelcode/utils"

)

type MainpageText struct{
	ID uint64 `gorm:"primaryKey"`
	PageName string `gorm:"size:100,not null"`
	Section string `gorm:"size:300"`
	Body string `gorm:"size:5000,not null"`
	ImagePath string `gorm:"size:300"`
	OrderT int `gorm:"not null"`
}

type OMainpageText struct{
	PageName string `json:"pageName"`
	Section string `json:"section"`
	Body string `json:"body"`
	ImagePath string `json:"imagePath"`
	OrderT int `json:"order"`
}

type IMainpageText struct{
	PageName string `json:"pageName"`
	Section string `json:"section"`
	Body string `json:"body"`
	Order int `json:"order"`
}


func (i *IMainpageText) Check() error{
	if err:= utils.IsNotInvalidCharacter(i.PageName); err!=nil{
		return errors.New("invalid pageName: " + err.Error())
	}
	if err:= utils.IsNotInvalidCharacter(i.Section); err!=nil{
		return errors.New("invalid section: " + err.Error())
	}
	if len(i.PageName) > 100{
		return errors.New("too long pageName")
	}
	if len(i.Section) > 100{
		return errors.New("too long section")
	}
	if len(i.Body) > 5000{
		return errors.New("too long body")
	}

	return nil
}

func (m *MainpageText) Fill(i *IMainpageText){
	m.PageName = i.PageName
	m.Section = i.Section
	m.Body = i.Body
	m.OrderT = i.Order
}
