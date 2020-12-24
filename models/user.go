package models

import (
	"database/sql"
)

type User struct {
	FirstName   string `json:"fname" validate:"required"`
	LastName    string `json:"lname" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	ID          string `json:"id"`
}

//Type 
type DBData struct {
	DBType, User, Password, Host, DBName string
	Session                              *sql.DB
	SessionToken                           map[string]User
}
