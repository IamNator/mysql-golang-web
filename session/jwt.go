package session

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
	"fmt"
)

func CreateToken(userid string) (token string, err error){

	atClaim := jwt.MapClaims{}
	atClaim["authorized"] = true
	atClaim["user_id"] = userid
	atClaim["exp"] = time.Now().Add(time.Minute * 30).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaim)
	///token, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	token, err = at.SignedString([]byte(getEnv()))
	if err != nil {
		return "", err
	}

	return token, nil
}

func getEnv() string {

		as := os.Getenv("ACCESS_SECRET")
		//fmt.Println(string(as))
		if as != "Our-web-app" {
			err := os.Setenv("ACCESS_SECRET", "Our-web-app")
			if err != nil {
				fmt.Println("failed to set ACCESS_SECRET")
			}
			fmt.Println("ACCESS_SECRET Set!")
		} else {
			fmt.Println("ACCESS_SECRET Already set")
		}
	return as
}