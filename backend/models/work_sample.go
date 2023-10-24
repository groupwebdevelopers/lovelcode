package models

import (
	"errors"
	"lovelcode/utils"
	"time"
	"strings"
	"strconv"

)

type WorkSample struct{
	ID uint64 `gorm:"primaryKey"`
	Title string `gorm:"not null"`
	ImagePath string `gorm:"size:400"`
	SiteUrl string `gorm:"size:200"`
	Description string `gorm:"size:800,not null"`
	IsFeatured bool

	DoneTime time.Time
}


type IWorkSample struct{
	Title string `json:"title"`
	Description string `json:"description"`
	SiteUrl string `json:"siteUrl"`
	DoneTime string `json:"doneTime"`
	IsFeatured bool `json:"isFeatured"`
}

// output article for user
type OWorkSample struct{
	Title string `json:"title"`
	Description string `json:"description"`
	ImagePath string `json:"imagePath"`
	SiteUrl string `json:"siteUrl"`
	DoneTime string `json:"doneTime"`
	IsFeatured bool `json:"isFeatured"`
}

func (w *WorkSample) FillWithIWorkSample(i IWorkSample){
	w.Title = i.Title
	w.Description = i.Description
	w.SiteUrl = i.SiteUrl
	w.IsFeatured = i.IsFeatured
	w.DoneTime = utils.ConvertToMiladiTime(utils.ConvertStringToTime(i.DoneTime, time.FixedZone("Tehran", 3.5 * 60 *60)))
}

func (a *IWorkSample) Check() error{
	if err:= utils.IsNotInvalidCharacter(a.Title, "/"); err!=nil{
		return errors.New("invalid titile:" + err.Error())
	}
	
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

	return nil
}
