package models

import (
	
)

type Statistics struct{
	ID uint64 `gorm:"primaryKey"`
	Name string `gorm:"size:100,not null"`
	Number uint64 `gorm:"not null"`

}

type OStatistics struct{
	Name string `json:"name"`
	Number string `json:"number"`
}