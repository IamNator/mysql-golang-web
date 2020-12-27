package session

import "github.com/IamNator/mysql-golang-web/models"

//
type LoginCredentials struct {
	Email    string `json:"email"`
	PassWord string `json:"password"`
}

//Sessiondb is a http.Handler
type Sessiondb models.DBData



type MyStdResp struct {
	Status bool
	Message interface{}
}