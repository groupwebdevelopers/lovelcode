package database


type SettingsDB struct{
	Key string `gorm:"unique,not null"`
	Value string `gorm:"not null"`
}

var Settings map[string]string

func SetupSettings() error{
	var st []SettingsDB
	if err:=DB.Find(&st).Error; err!=nil{
		return err
	}
	
	Settings = make(map[string]string)
	for _, s := range st{
		Settings[s.Key] = s.Value
	}

	return nil
}