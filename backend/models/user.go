package models

import (
	"time"
	"errors"
	"lovelcode/utils"

)

// when AdminPermisions is "" means it not get from database
// for normal user AdminPermisions is "0"
type User struct {
	ID uint64 `gorm:"primaryKey"`
	Name string `gorm:"not null,size:50"`
	Family string `gorm:"size:50"`
	// Number int `gorm:"not null,unique"`
	Email string `gorm:"not null,unique,size:70"`
	Password string `gomr:"not null"`
	AdminPermisions string `gorm:"not null"`
	IsDeleted bool `gorm:"not null"`
	IsBanned bool `gorm:"not null"`
	Token string `gorm:"unique,size:64"`
	TokenExp time.Time
}

type SigninUser struct{
	Email string `json:"email"`
	// Number int `json:"number"`
	Password string `json:"password"`
}

type SignupUser struct{
	Email string `json:"email"`
	Password string `json:"password"`
	Name string `json:"name"`
	Family string `json:"family"`
	// Number int `json:"number"`
}

type OUser struct{
	Name string `json:"name"`
	Family string `json:"family"`
	Email string `json:"email"`
}

func (u *SignupUser) Check() error{
	if len(u.Name) > 50{
		return errors.New("too big name")
	}
	if len(u.Family) > 50{
		return errors.New("too big family")
	}
	if len(u.Email) > 70{
		return errors.New("too big email")
	}
	if err := utils.IsNotInvalidCharacter(u.Name, "!@#$%^&*()_+-={}|[]\\:\";<>?./'"); err!=nil{
		return errors.New("invalid name: "+ err.Error())
	}
	if u.Family != "" {
		if err := utils.IsNotInvalidCharacter(u.Family, "!@#$%^&*()_+-={}|[]\\:\";<>?./'"); err!=nil{
			return errors.New("invalid family: "+ err.Error())
		}
	}
	// if u.Number < 9000000000{
		// return errors.New("small number")
	// }
	if u.Email != ""{
		if err := utils.CheckEmail(u.Email); err!=nil{
			return errors.New("invalid email")
		}
	}else{
		return errors.New("empty email")
	}
	if len(u.Password) < 8{
		return errors.New("small password (<8)")
	}
	
	return nil
}

func (u *SigninUser) Check() error{
	// if u.Number < 9000000000{
		// return errors.New("small number")
	// }
	if len(u.Email) > 70{
		return errors.New("too big email")
	}
	if u.Email != ""{
		if err := utils.CheckEmail(u.Email); err!=nil{
			return errors.New("invalid email")
		}
	}else{
		return errors.New("empty email")
	}
	if len(u.Password) < 8{
		return errors.New("small password (<8)")
	}
	

	return nil
}
