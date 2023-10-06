package models


type User struct {
	ID uint32
	Username string
	Email string
	password string
	AdminPermisions string
	IsDeleted bool

}