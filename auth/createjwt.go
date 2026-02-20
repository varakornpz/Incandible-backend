package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func getJWT(uuid uuid.UUID , jwtSecret string) (string , error) {
	claims := jwt.MapClaims{
		"uuid" : uuid ,
		"iat": time.Now().Unix(),
		//Expired next 48hr
		"exp" : time.Now().Add(time.Hour * 48).Unix() ,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256 , claims)

	t , tokenErr := token.SignedString([]byte(jwtSecret))
	if tokenErr != nil {
		return  "" , errors.New("Cant sign token.")
	}
	return t , nil
}