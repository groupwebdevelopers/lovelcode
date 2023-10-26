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
	JobTitle string `gorm:"not null,size:50"`
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
	UserID uint64 `json:"userID"`
	JobTitle string `json:"jobTitle"`
	WorkExp int `json:"workExp"`
	Contact string `json:"contact"`
	ImagePath string `json:"imagePath"`
	
	Name string `json:"name"`
	Family string `json:"family"`
	Email string `json:"email"`
}

func (m *Member) Fill(im *IMember) {
	m.JobTitle = im.JobTitle
	m.WorkExp = im.WorkExp
	m.Contact = im.Contact
}

func (im *IMember) Check() error{
	if err:=utils.IsNotInvalidCharacter(im.JobTitle); err!=nil{
		return errors.New("invalid jobTitle:"+ err.Error())
	}
	if err:=utils.IsNotInvalidCharacter(im.Contact); err!=nil{
		return errors.New("invalid Contact:"+ err.Error())
	}

	if len(im.JobTitle) > 50{
		return errors.New("too long jobTitle")
	}
	if im.WorkExp > 100{
		return errors.New("too long workExp")
	}
	if len(im.Contact) > 200{
		return errors.New("too long contact")
	}

	return nil
}