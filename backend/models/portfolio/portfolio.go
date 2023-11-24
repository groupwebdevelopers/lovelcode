package models

import (
	"errors"
	"lovelcode/utils"
	"time"
	"strings"
	"strconv"

)

type Portfolio struct{
	ID uint64 `gorm:"primaryKey"`
	Title string `gorm:"not null"`
	ImagePath string `gorm:"size:400"`
	LogoPath string `gorm:"size:400"`
	SiteUrl string `gorm:"size:200"`
	Description string `gorm:"size:800,not null"`
	IsFeatured bool
	Type string

	DoneTime time.Time
}


type IPortfolio struct{
	Title string `json:"title"`
	Description string `json:"description"`
	SiteUrl string `json:"siteUrl"`
	DoneTime string `json:"doneTime"`
	IsFeatured bool `json:"isFeatured"`
	Type string `json:"type"`
}

// output article for user
type OPortfolio struct{
	ID uint64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	ImagePath string `json:"imagePath"`
	LogoPath string `json:"logoPath"`
	SiteUrl string `json:"siteUrl"`
	DoneTime string `json:"doneTime"`
	IsFeatured bool `json:"isFeatured"`
	Type string `json:"type"`
}

func (w *Portfolio) Fill(i *IPortfolio){
	w.Title = i.Title
	w.Description = i.Description
	w.SiteUrl = i.SiteUrl
	w.IsFeatured = i.IsFeatured
	w.DoneTime = utils.ConvertToMiladiTime(utils.ConvertStringToTime(i.DoneTime, time.FixedZone("Tehran", 3.5 * 60 *60)))
	w.Type = i.Type
}

func (a *IPortfolio) Check() error{
	if err:= utils.IsNotInvalidCharacter(a.Title, "/"); err!=nil{
		return errors.New("invalid titile:" + err.Error())
	}

	if len(a.Description) > 800{
		return errors.New("too long desctiption")
	}
	if len(a.SiteUrl) > 200{
		return errors.New("too long body")
	}
	
	if a.DoneTime != ""{
	// doneTime format year-month-day
	splited := strings.Split(a.DoneTime, "-")
	if len(splited) < 3{
		return errors.New("invalid doneTime")
	}
	for _, s := range splited{
		_, err := strconv.Atoi(s)
		if err!=nil{
			return errors.New("invalid doneTime")
		}
	}
	}else{
		return errors.New("empty doneTime")
	}

	return nil
}
