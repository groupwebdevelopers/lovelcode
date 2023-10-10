package models

// when AdminPermisions is "" means it not get from database
// for normal user AdminPermisions is "0"
type User struct {
	ID uint32 `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Family string 
	Username string `gorm:"not null,unique"`
	Email string `gorm:"not null,unique`
	Password string `gomr:"not null`
	AdminPermisions string `gorm:"not null"`
	IsDeleted bool `gorm:"not null"`
	IsBanned bool
}