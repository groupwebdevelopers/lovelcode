package utils

import (
	"net/mail"
	"strings"
	"errors"
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
	for _, c := range allows[0]{
		valid += string(c)
	}
	for _, c := range s{
		if !strings.Contains(valid, string(c)){
			return errors.New("invalid character: "+string(c))
		}
	}
	return nil
}

// support persion character
// invalid characters is \/
func IsNotInvalidCharacter(s string, disallows... string) error{

	if s == "" || s == " "{
		return errors.New("empty field")
	}

	invalid := "\\/"

	invalid += disallows[0]

	for _, c := range s{
		if strings.Contains(invalid, string(c)){
			return errors.New("invalid character: "+string(c))
		}
	}
	return nil
}