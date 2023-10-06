package utils


import (
	"os"
	"time"
	"errors"

	"github.com/golang-jwt/jwt"

	"lovelcode/models"
)


func CreateToken(user models.User, tokenExpHours uint8) (string, error){



	// create token
	secret := os.Getenv("secret") + "se"
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA,
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

// todo: send token with expared time
// verify token
func VerifyToken(tokenString string) (models.User, error){
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