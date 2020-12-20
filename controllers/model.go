package controllers

import "database/sql"

type DBData struct {
	DBType, User, Password, Host, DBName string
	Session                              *sql.DB
	SessionIDs                           map[string]string
	SessionUsers                         map[string]string
}
