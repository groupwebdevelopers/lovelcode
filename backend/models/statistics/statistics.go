package statistics

import (
	"lovelcode/utils"
	"errors"	
)

type Statistic struct{
	ID uint64 `gorm:"primaryKey"`
	Name string `gorm:"size:100,not null,unique"`
	Name2 string `gorm:"size:100,not null"`
	Number float64 `gorm:"not null"`
	IsPublic bool

}

type OStatistic struct{
	Name string `json:"name"`
	Name2 string `json:"name2"`
	Number float64 `json:"number"`
}

type IStatistic struct{
	Name string `json:"name"`
	Name2 string `json:"name2"`
	Number float64 `json:"number"`
}

func (i *IStatistic) Check() error{
	if err:= utils.IsJustLetter(i.Name, "-"); err!=nil{
		return errors.New("invalid name:"+err.Error())
	}
	if err:= utils.IsNotInvalidCharacter(i.Name2);err!=nil{
		return errors.New("invalid name2:"+err.Error())
	}

	return nil
} 

func (s *Statistic) Fill(i *IStatistic){
	s.Name = i.Name
	s.Name2 = i.Name2
	s.Number = i.Number
}