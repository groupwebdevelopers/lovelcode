package token


import (
	"os"
	"time"
	"errors"
	"math/rand"

	"github.com/golang-jwt/jwt"

	"lovelcode/models"
)

func Setup(){
	rand.Seed(time.Now().UnixNano())
}


func CreateJWTToken(user models.User, tokenExpHours uint16) (string, error){



	// create token
	var secret = []byte( os.Getenv("secret") + "se")
	// var secret = []byte("testestte")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
				"exp" : time.Now().Add(time.Duration(tokenExpHours) * time.Hour).Unix(),
				"iss": "localhost", // todo:must changed
			
				"email": user.Email,
				"name" : user.Name,
				"family": user.Family,
		})
	stoken, err := token.SignedString(secret)
	return stoken, err


}


func CreateToken() string{
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result:= ""
	for i:=0;i<256;i++{
		result +=string(chars[ rand.Intn(len(chars))])
	}
	return result
}

// todo: send token with expared time
// verify token
func VerifyJWTToken(tokenString string) (models.User, error){
	secret := os.Getenv("secret") + "se"

	
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		if _, ok:= token.Method.(*jwt.SigningMethodECDSA); !ok{
			return nil, errors.New("invalid signing method")
		}
		return secret, nil
	})

	if err!=nil{
		return models.User{}, err
	}

	claims, ok:=token.Claims.(jwt.MapClaims)
	if ok && token.Valid{
		user := models.User{
			Email: claims["email"].(string),
			Name: claims["name"].(string),
			Family: claims["family"].(string),
		}
		return user, nil
	}
	return models.User{}, errors.New("invalid token")

}

func CreateRandomToken() string{
	chars := "abcdefghijklmanopqrstuvwxryzABCDEFhGHIJKLMNOPQRSTUVWXYZ-_=)(*&^%$#@!1234567890-=-[]|}{;d;vdvk:/.,.,/})"
	result := ""
	for i:=0;i<128;i++{
		result += string(chars[rand.Intn(len(chars))])
	}
	return result
}