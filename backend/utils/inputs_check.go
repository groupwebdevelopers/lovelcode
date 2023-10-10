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

func IsJustLetter(s string, allows string) error{

	valid := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, c := range allows{
		valid += string(c)
	}
	for _, c := range s{
		if !strings.Contains(valid, string(c)){
			return errors.New("invalid character: "+string(c))
		}
	}
	return nil
}

