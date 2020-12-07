package models

import (
	"database/sql"
)

type User struct {
	FirstName   string `json:"fname"`
	LastName    string `json:"lname"`
	PhoneNumber string `json:"phone_number"`
	ID          string `json:"id"`
}

type DBData struct {
	DBType, User, Password, Host, DBName string
	Session                              *sql.DB
	SessionIDs                           map[string]string
}
