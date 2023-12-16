package utils

import (
	"errors"
	"net/mail"
	"strings"
	// "encoding/hex"
)


func CheckEmail(e string) error{
	_, err := mail.ParseAddress(e)
	return err
}

// valid character is english characters
func IsJustLetter(s string, allows... string) error{

	if s == "" || s == " "{
		return errors.New("empty field")
	}

	valid := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if len(allows) >0{
		valid += allows[0]
	}
	
	for _, c := range s{
		if !strings.Contains(valid, string(c)){
			return errors.New("invalid character: "+string(c))
		}
	}
	return nil
}

// support persion character
// invalid characters is \;'""
func IsNotInvalidCharacter(s string, disallows... string) error{
	
	if s == "" || s == " "{
		return errors.New("empty field")
	}
	
	invalid := "\\;'\"<>"
	
	if len(disallows) > 0{
		invalid += disallows[0]
	}
	
	for _, c := range s{
		if strings.Contains(invalid, string(c)){
			return errors.New("invalid character: "+string(c))
		}
	}

	return nil
}

// func ConvertToUrl(s string) string{
// 	replaced :=  strings.Replace(s, " ", "-", -1)
// 	hexed := hex.EncodeToString([]byte(replaced))
// 	result := "%"
// 	for i, v := range hexed{
// 		if (i+1)%2==0{
// 			result += "%"
// 		}
// 		if v!= '-'
// 		result += string(v)
// 	}

// 	return result
// }