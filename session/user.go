package session

import "github.com/IamNator/mysql-golang-web/models"

// swagger:model
type LoginCredentials struct {
	//user's email address
	Email    string `json:"email"`
	//user's password
	PassWord string `json:"password"`
}

//Sessiondb is a http.Handler
type Sessiondb models.DBData

// swagger:model
type MyStdResp struct {
	// successful / not successful
	Status bool	   `json:"status`
	// string
	Message string `json:"message"`
}
