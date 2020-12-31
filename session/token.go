package session

import (
	"github.com/IamNator/mysql-golang-web/models"
	uuid "github.com/satori/go.uuid"
	"sync"
)

var Mutex sync.Mutex

//creates a
func CreateToken(db *Sessiondb, user models.UserCredentials) (token string) {
	token = uuid.NewV1().String()
	Mutex.Lock()
	db.SessionToken[token] = user
	Mutex.Unlock()
	return token
}
