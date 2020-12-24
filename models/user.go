package models

import (
	"database/sql"
)

type PhoneBookContact struct {
	FirstName   string `json:"fname" validate:"required"`
	LastName    string `json:"lname" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	ID          string `json:"id"`
}

//Type 
type DBData struct {
	DBType, User, Password, Host, DBName string
	Session                              *sql.DB
	SessionToken                           map[string]UserCredentials
}


type UserCredentials struct {
	ID          string `json:"id"`
	FirstName   string `json:"firstname" validate:"required"`
	LastName    string `json:"lastname" validate:"required"`
	Email 		string `json:"email" validate:"required"`
	PassWord    string `json:"password" validate:"required"`
}
