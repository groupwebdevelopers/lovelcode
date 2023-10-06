package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	
	"lovelcode/models"
)

var DB *gorm.DB

func Setup() error{
	dsn := "mohammadamin:M@85mmohammadamin@tcp(127.0.0.1:3306)/lovelcode?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		return err
	}
	DB = db
	db.AutoMigrate(&models.User{})

	return nil
}