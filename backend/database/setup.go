package database

import (
	"gorm.io/gorm"
	"gorm.io/dirver/mysql"
)

var DB *gorm.DB

func Setup() error{
	dsn := ""
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		return err
	}
	DB = &db
	db.AutoMigrate(&User{})

	return nil
}