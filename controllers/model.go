package controllers

import "database/sql"

type DBData struct {
	DBType string
	User string
	Password string
	Host	string
	DBName	string
	Session                              *sql.DB
	SessionIDs                           map[string]string
	SessionToken                         map[string]string
}
