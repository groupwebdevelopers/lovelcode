package models

import (
	"errors"
	"lovelcode/utils"
	"strconv"
	"strings"
)


type SettingsDB struct{
	Key string `gorm:"unique,not null" json:"key"`
	Value string `gorm:"not null,size:800" json:"value"`
}

type Settings struct{
	TokenExpHours uint64
	SiteFeatures []SiteFeature
}

type SiteFeature struct{
	ImagePath string `json:"imagePath"`
	Name string `json:"name"`
}


func SetupSettings(st []SettingsDB) (Settings, error){
	var settings Settings
	
	// set default
	settings.TokenExpHours = 72


	for _, s := range st{
		if s.Value != ""{
			
			switch s.Key{
			
			case "tokenExpHours":
				i, err := strconv.Atoi(s.Value)
				if err!=nil{
					return settings,errors.New("invalid tokenExpHours in database")
				}
				settings.TokenExpHours = uint64(i)
			
			case "siteFeature":
				splited := strings.Split(s.Value, "|||")
				settings.SiteFeatures = append(settings.SiteFeatures, SiteFeature{ImagePath: splited[1], Name: splited[0]})
			default:
				return settings, errors.New("unhandled setting: "+ s.Key+" value:" +s.Value)
			}
		}else{
			return settings, errors.New("empty value key: "+ s.Key)	
		}
	}

	return settings, nil
}

func (st *SettingsDB) Check() error{
	if err:=utils.IsJustLetter(st.Key);err!=nil{
		return errors.New("invalid setting key: "+err.Error())
	}
	if err:=utils.IsNotInvalidCharacter(st.Key);err!=nil{
		return errors.New("invalid setting value: "+err.Error())
	}

	return nil
}