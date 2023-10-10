package models

import (
	"time"
)

type ProjectDoingRequest struct{
	ID uint64 `gorm:"primaryKey"`
	UserID uint64 `gorm:"not null"`
	User User `gorm:"not null"`
	Title string `gorm:"not null"`
	Description string `gorm:"not null"`
	SuggestedPrice uint 
	TimeCreated time.Time `gorm:"not null"`
	TimeModified time.Time `gorm:"not null"`
}