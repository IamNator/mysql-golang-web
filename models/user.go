package models

import (
	"database/sql"
)

type PhoneBookContact struct {
	FirstName   string `json:"firstname" validate:"required"`
	LastName    string `json:"lastname" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

//Type 
type DBData struct {
	DBType, User, Password, Host, DBName string
	Session                              *sql.DB
	SessionToken                           map[string]UserCredentials
}

//For registering new users
type UserCredentials struct {
	FirstName   string `json:"firstname" validate:"required"`
	LastName    string `json:"lastname" validate:"required"`
	Email 		string `json:"email" validate:"required"`
	PassWord    string `json:"password" validate:"required"`
}
