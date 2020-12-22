package session

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func CreateToken(userid string) (token string, err error){
	atClaim := jwt.MapClaims{}
	atClaim["authorized"] = true
	atClaim["user_id"] = userid
	atClaim["exp"] = time.Now().Add(time.Minute * 30).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaim)
	token, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}