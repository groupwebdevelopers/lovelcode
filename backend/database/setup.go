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
	if os.Getenv("deploy") == "true"{
		dsn := "host=database user=root password=2CEezrHZLl3SP5VnNhu4kdto dbname=lovelcode port=5432 sslmode=disable TimeZone=Asia/Tehran"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		
	}else{
		dsn := "mohammadamin:M@85mmohammadamin@tcp(127.0.0.1:3306)/lovelcode?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	}
	
		if err!=nil{
		return err
	}
	DB = db
	err = db.AutoMigrate(
		&models.User{},
		&models.SettingsDB{},
		&models.ProjectDoingRequest{},
		&models.Plan{},
		&models.Feature{},
	)

	if err!=nil{
		return err
	}

	err = SetupSettings()

	// create Owner
	// db.Create(models.User{
	// 	Name: "Owner",
	// 	Family: "Owner",
	// 	Username: "1000000",
	// 	Email: "theowner@localhost.lh",

	// })

	return err
}


func SetupSettings() error{
	var st []models.SettingsDB
	if err:=DB.Find(&st).Error; err!=nil{
		return err
	}
	
	Settings = make(map[string]string)
	for _, s := range st{
		if s.Value != ""{
			Settings[s.Key] = s.Value
		}else {
			Settings[s.Key] = "72"
		}
	}

	return nil
}