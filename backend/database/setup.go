package database

import (
	"os"
	"runtime"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"

	"lovelcode/models"
)

var DB *gorm.DB
var Settings models.Settings

func Setup() error{
	var err error
	var db *gorm.DB
	if os.Getenv("dev") == "true" || runtime.GOOS == "windows"{
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
	err = db.AutoMigrate(
		&models.User{},
		&models.SettingsDB{},
		&models.ProjectDoingRequest{},
		&models.Plan{},
		&models.Feature{},
		&models.Member{},
		&models.Article{},
		&models.SettingsDB{},
	)

	if err!=nil{
		return err
	}

	var st []models.SettingsDB
	if err:=DB.Find(&st).Error; err!=nil{
		return err
	}
	
	Settings, err = models.SetupSettings(st)

	// create Owner
	// db.Create(models.User{
	// 	Name: "Owner",
	// 	Family: "Owner",
	// 	Username: "1000000",
	// 	Email: "theowner@localhost.lh",

	// })

	return err
}

