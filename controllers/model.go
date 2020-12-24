package controllers

import (
	"database/sql"
	"github.com/IamNator/mysql-golang-web/models"
)

type DBData struct {
	DBType, User, Password, Host, DBName string
	Session                              *sql.DB
	SessionToken                         map[string]models.User
}
