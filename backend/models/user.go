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
	Name string `gorm:"not null"`
	Family string 
	// Number int `gorm:"not null,unique"`
	Email string `gorm:"not null,unique`
	Password string `gomr:"not null`
	AdminPermisions string `gorm:"not null"`
	IsDeleted bool `gorm:"not null"`
	IsBanned bool
	Token string `gorm:"unique"`
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

func (u *SignupUser) Check() error{
	if err := utils.IsNotInvalidCharacter(u.Name, "!@#$%^&*()_+-={}|[]\\:\";<>?./'"); err!=nil{
		return errors.New("invalid name: "+ err.Error())
	}
	if err := utils.IsNotInvalidCharacter(u.Family, "!@#$%^&*()_+-={}|[]\\:\";<>?./'"); err!=nil{
		return errors.New("invalid name: "+ err.Error())
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

