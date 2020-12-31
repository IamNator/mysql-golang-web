package models

import (
	"database/sql"
)

//
type PhoneBookContact struct {
	FirstName   string `json:"firstname" validate:"required"`
	LastName    string `json:"lastname" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	ID          string `json:"id"`
}

//For database connection
type DBData struct {
	DBType string   `mapstructure:"DB_TYPE"`
	User string     `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	Host string     `mapstructure:"HOST"`
	DBName string	`mapstructure:"DB_NAME"`
	Session                              *sql.DB
	SessionToken                         map[string]UserCredentials
}


//For registering new users
type UserCredentials struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required"`
	PassWord  string `json:"password" validate:"required"`
	ID        string `json:"id"`
}
