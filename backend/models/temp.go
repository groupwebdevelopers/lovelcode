package models

import (
	"time"
)

type Temp struct{
	ID uint64 `gorm:"primaryKey"`
	String1f string `gorm:"size:200"`
	String2f string `gorm:"size:200"`

	EndTime time.Time
}