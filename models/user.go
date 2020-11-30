package models

import (
	"database/sql"
)

type User struct {
	FirstName        string `json:"fname"`
	LastName         string `json:"lname"`
	PhoneNumber      string `json:"phone_number"`
	ID               string `json:"id"`
}

type DBData struct {
	DBType, User, Password, Host, DBName string
	Session                              *sql.DB
}

//type DBDATA interface {
//	//From controllers
//	Fetch(w http.ResponseWriter, req *http.Request)
//	Delete(writer http.ResponseWriter, req *http.Request)
//	Update(w http.ResponseWriter, req *http.Request)
//	//From migrations
//	FillDb()
//	CreateUserDb()
//	CreatePhoneBookDb()
//	//From users
//	Register(w http.ResponseWriter, req *http.Request)
//	Login(w http.ResponseWriter, req * http.Request)
//}
