package session

import (
	//"github.com/IamNator/mysql-golang-web/models"
	"github.com/IamNator/mysql-golang-web/models"
	uuid "github.com/satori/go.uuid"
)



func CreateToken(db *Sessiondb, user UserCredentials) (token string){
	token = uuid.NewV1().String()
	userr := models.UserCredentials{
		user.FirstName,
		user.LastName,
		user.Email,
		user.PassWord,
	}
	db.SessionToken[token] = userr
	return token
}
