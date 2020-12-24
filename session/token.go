package session

import (
	"github.com/IamNator/mysql-golang-web/models"
	uuid "github.com/satori/go.uuid"
)

func CreateToken(db *Sessiondb, user models.UserCredentials) {
	token := uuid.NewV1().String()
	db.SessionToken[token] = user
}
