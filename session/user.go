package session

import "github.com/IamNator/mysql-golang-web/models"

//For registering new users
type Credentials struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	PassWord  string `json:"password"`
}

type LoginCredentials struct {
	Email    string `json:"email"`
	PassWord string `json:"password"`
}

type Sessiondb models.DBData

type MyStdResp struct {
	Status bool
	Message interface{}
}