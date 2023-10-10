package models


type SettingsDB struct{
	Key string `gorm:"unique,not null"`
	Value string `gorm:"not null"`
}
