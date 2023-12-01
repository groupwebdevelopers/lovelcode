package hutils

import (
	"strings"
	
	amodels "lovelcode/models/article"
	pfmodels "lovelcode/models/portfolio"
	"lovelcode/utils"
)

// removte time hours and minutes and seconds from time
func ConvertArticleStringTimesForOutput(st []amodels.OArticleTitle) error {
	for i:= range st{

		t, err := utils.ConvertStringTimeToPersianStringTime(st[i].TimeCreated)
		if err!=nil{
			return err
		}
		st[i].TimeCreated =strings.Split(t, " ")[0]
		t, err = utils.ConvertStringTimeToPersianStringTime(st[i].TimeModified)

		if err!=nil{
			return err
		}
		st[i].TimeModified =strings.Split(t, " ")[0]
		
	}

	return nil
}

// removte time hours and minutes and seconds from time
func ConvertPortfolioStringTimesForOutput(st []pfmodels.OPortfolio) error {
	for i:= range st{

		t, err := utils.ConvertStringTimeToPersianStringTime(st[i].DoneTime)
		if err!=nil{
			return err
		}
		st[i].DoneTime =strings.Split(t, " ")[0]		
	}

	return nil
}