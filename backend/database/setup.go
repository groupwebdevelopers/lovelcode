package database

import (
	"os"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"

	"lovelcode/models"
)

var DB *gorm.DB
var Settings map[string]string

func Setup() error{
	var err error
	var db *gorm.DB
	if os.Getenv("dev") == "true"{
		dsn := "mohammadamin:M@85mmohammadamin@tcp(127.0.0.1:3306)/lovelcode?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		
	}else{
		dsn := "host=database user=root password=2CEezrHZLl3SP5VnNhu4kdto dbname=lovelcode port=5432 sslmode=disable TimeZone=Asia/Tehran"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	}
	
		if err!=nil{
		return err
	}
	DB = db
	db.AutoMigrate(&models.User{}, &models.SettingsDB{}, &models.ProjectDoingRequest{})

	err = SetupSettings()

	return err
}


func SetupSettings() error{
	var st []models.SettingsDB
	if err:=DB.Find(&st).Error; err!=nil{
		return err
	}
	
	Settings = make(map[string]string)
	for _, s := range st{
		Settings[s.Key] = s.Value
	}

	return nil
}