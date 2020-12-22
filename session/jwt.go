package session

import (
	"github.com/dgrijalva/jwt-go"
	"os"
)

func CreateToken(userid string) (token string, err error){
	err = os.Setenv("ACCESS_SECRET", "Our-web-app")
	atClaim := jwt.MapClaims{}
	atClaim["authorized"] = true

	return token, err
}