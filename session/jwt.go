package session

import "os"

func CreateToken(userid string) (token string, err error){
	err = os.Setenv("ACCESS_SECRET", "Our-web-app")

	return token, err
}