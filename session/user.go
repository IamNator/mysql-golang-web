package session

import "github.com/IamNator/mysql-golang-web/models"

type LoginCredentials struct {
	Email    string `json:"email"`
	PassWord string `json:"password"`
}

type Sessiondb models.DBData

type MyStdResp struct {
	Status bool
	Message interface{}
}