package models

import (
	"errors"
	"lovelcode/utils"
	"strconv"
	
)


type SettingsDB struct{
	ID uint64 `gorm:"primaryKey" json:"id"`
	Key string `gorm:"not null,size:100" json:"key"`
	Value string `gorm:"not null,size:800" json:"value"`
}

type ISettingsDB struct{
	Key string `json:"key"`
	Value string `json:"value"`
}

type Settings struct{
	TokenExpHours uint64
	PageLength int
	ImageSaveUrl string
	SocialMedias string
}

func SetupSettings(st []SettingsDB) (Settings, error){
	var settings Settings
	
	// set default
	settings.TokenExpHours = 72
	settings.PageLength = 20
	settings.ImageSaveUrl = "/../frontend/dist/images/"

	for _, s := range st{
		if s.Value != ""{
			
			switch s.Key{
			
			case "tokenExpHours":
				i, err := strconv.Atoi(s.Value)
				if err!=nil{
					return settings,errors.New("invalid tokenExpHours in database")
				}
				settings.TokenExpHours = uint64(i)
			
			case "pageLength":
				i, err := strconv.Atoi(s.Value)
				if err!=nil{
					return settings,errors.New("invalid pageLength in database")
				}
				settings.PageLength = i
			case "imageSaveUrl":
				settings.ImageSaveUrl = s.Value
			case "socialMedias":
				settings.SocialMedias = s.Value

			default:
				return settings, errors.New("unhandled setting: "+ s.Key+" value:" +s.Value)
			}
		}else{
			return settings, errors.New("empty value key: "+ s.Key)	
		}
	}

	return settings, nil
}

func (st *ISettingsDB) Check() error{
	if err:=utils.IsJustLetter(st.Key);err!=nil{
		return errors.New("invalid setting key: "+err.Error())
	}
	if err:=utils.IsNotInvalidCharacter(st.Key);err!=nil{
		return errors.New("invalid setting value: "+err.Error())
	}
	if len(st.Key) > 100{
		return errors.New("too long key")
	}
	if len(st.Value) > 800{
		return errors.New("too long value")
	}

	return nil
}

func (st *SettingsDB) Fill(i *ISettingsDB) {
	st.Key = i.Key
	st.Value = i.Value
}