package database

import (
	"log"
	"os"
	"runtime"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	amodels "lovelcode/models/article"
	comodels "lovelcode/models/contactus"
	cumodels "lovelcode/models/customer"
	mmodels "lovelcode/models/mainpage"
	pmodels "lovelcode/models/plan"
	pfmodels "lovelcode/models/portfolio"
	smodels "lovelcode/models/settings"
	ssmodels "lovelcode/models/statistics"
	umodels "lovelcode/models/user"
	tmodels "lovelcode/models/temp"
	// "lovelcode/models"
)

var DB *gorm.DB
var Settings smodels.Settings
var MainpagesTexts []mmodels.OMainpageText

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
		&umodels.User{},
		&smodels.SettingsDB{},
		&pmodels.Plan{},
		&pmodels.Feature{},
		&umodels.Member{},
		&amodels.Article{},
		&smodels.SettingsDB{},
		&tmodels.Temp{},
		&pfmodels.Portfolio{},
		&amodels.Comment{},
		&comodels.ContactUs{},
		&cumodels.Customer{},
		&mmodels.MainpageText{},
		&ssmodels.Statistic{},
		&amodels.ArticleCategory{},
		&pmodels.PlanType{},
		&pmodels.PlanOrder{},
	)

	if err!=nil{
		return err
	}

	var st []smodels.SettingsDB
	if err:=DB.Find(&st).Error; err!=nil{
		return err
	}
	
	Settings, err = smodels.SetupSettings(st)

	if Settings.MainpageInMemory{
		// get mainpages from database and set into array
		if err:= DB.Model(&mmodels.MainpageText{}).Find(&MainpagesTexts).Error; err!=nil{
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
	var st []smodels.SettingsDB
	if err=DB.Find(&st).Error; err!=nil{
		log.Fatal(err)
	}
	Settings, err = smodels.SetupSettings(st)
	if err!=nil{
		log.Fatal(err)
	}
}