package database

import (
	"log"
	"os"
	"runtime"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"lovelcode/models"
	_ "lovelcode/models/user"
	_ "lovelcode/models/plan"
)

var DB *gorm.DB
var Settings models.Settings
var MainpagesTexts []models.OMainpageText

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
		&models.Temp{},
		&models.Portfolio{},
		&models.Comment{},
		&models.ContactUs{},
		&models.Customer{},
		&models.MainpageText{},
		&models.Statistic{},
		&models.ArticleCategory{},
		&models.PlanType{},
		&models.OrderPlan{},
	)

	if err!=nil{
		return err
	}

	var st []models.SettingsDB
	if err:=DB.Find(&st).Error; err!=nil{
		return err
	}
	
	Settings, err = models.SetupSettings(st)

	if Settings.MainpageInMemory{
		// get mainpages from database and set into array
		if err:= DB.Model(&models.MainpageText{}).Find(&MainpagesTexts).Error; err!=nil{
			return err
		}
	}

	// create Owner
	// db.Create(models.User{
	// 	Name: "Owner",
	// 	Family: "Owner",
	// 	Username: "1000000",
	// 	Email: "theowner@localhost.lh",

	// })

	return err
}

func RegetSettings(){
	var err error
	var st []models.SettingsDB
	if err=DB.Find(&st).Error; err!=nil{
		log.Fatal(err)
	}
	Settings, err = models.SetupSettings(st)
	if err!=nil{
		log.Fatal(err)
	}
}