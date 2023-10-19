package models

import (
	"errors"
	"lovelcode/utils"
	"time"
)

type Member struct{
	ID uint64 `gorm:"primaryKey"`
	UserID uint64
	User User
	JobTitle string `gorm:"not null"`
	ImagePath string `gorm:"size:100"`
	WorkExp int `gorm:"not null"`  // woek experience
	Contact string `gorm:"size:200"` // splited with |

	TimeCreated time.Time
	TimeModified time.Time
}

type IMember struct{
	JobTitle string `json:"jobTitle"`
	WorkExp int `json:"workExp"`
	Contact string `json:"contact"`
}

type OMember struct{
	JobTitle string
	WorkExp int
	Contact string
	OUser OUser
}

func (m *Member) FillWithIMember(im IMember) {
	m.JobTitle = im.JobTitle
	m.WorkExp = im.WorkExp
	m.Contact = im.Contact
}

func (o *OMember) FillWithMember(m Member) {
	o.JobTitle = m.JobTitle
	o.WorkExp = m.WorkExp
	o.Contact = m.Contact
	o.OUser.FillWithUser(m.User)

}

func (im *IMember) Check() error{
	if err:=utils.IsNotInvalidCharacter(im.JobTitle); err!=nil{
		return errors.New("invalid jobTitle:"+ err.Error())
	}
	if err:=utils.IsNotInvalidCharacter(im.Contact); err!=nil{
		return errors.New("invalid Contact:"+ err.Error())
	}

	return nil
}