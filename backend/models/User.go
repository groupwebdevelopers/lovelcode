package models

// when AdminPermisions is "" means it not get from database
// for normal user AdminPermisions is "0"
type User struct {
	ID uint32
	Name string
	Family string
	Email string
	Password string
	AdminPermisions string
	IsDeleted bool

}