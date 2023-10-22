package models

import (
	"errors"
	"lovelcode/utils"
	"time"
)

type WorkSample struct{
	ID uint64 `gorm:"primaryKey"`
	Title string `gorm:"not null"`
	ImagePath string `gorm:"size:400"`
	SiteUrl string `gorm:"size:200"`
	Description string `gorm:"size:800,not null"`

	DoneTime time.Time
}


type IWorkSample struct{
	Title string `json:"title"`
	Description string `json:"description"`
	SiteUrl string `json:"siteUrl"`
	DoneTime time.Time `json:"time"`
}

// output article for user
type OWorkSample struct{
	Title string `json:"title"`
	Description string `json:"description"`
	ImagePath string `json:"imagePath"`
	SiteUrl string `json:"siteUrl"`
	DoneTime time.Time `json:"time"`
}

func (w *WorkSample) FillWithIWorkSample(i IWorkSample){
	w.Title = i.Title
	w.Description = i.Description
	w.SiteUrl = i.SiteUrl
	w.DoneTime = i.DoneTime
}

func (o *OWorkSample) FillWithWorkSample(a WorkSample){
	o.Title = a.Title
	o.Description = a.Description
	o.ImagePath = a.ImagePath
	o.SiteUrl = a.SiteUrl
	o.ImagePath = a.ImagePath
	o.DoneTime = a.DoneTime
}

func (a *IWorkSample) Check() error{
	if err:= utils.IsNotInvalidCharacter(a.Title, "/"); err!=nil{
		return errors.New("invalid titile:" + err.Error())
	}
	
	return nil
}