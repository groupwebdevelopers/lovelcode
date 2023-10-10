package models

import (
	"time"
)

type ProjectDoingRequest struct{
	ID uint64
	UserID uint64
	User User
	Title string
	Description string
	SuggestedPrice uint
	TimeCreated time.Time
	TimeModified time.Time
}