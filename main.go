package main

import (
	"fmt"
	"github.com/IamNator/mysql-golang-web/controllers"
	"github.com/IamNator/mysql-golang-web/database/migrations"
	"github.com/IamNator/mysql-golang-web/database/seeders"
	"github.com/IamNator/mysql-golang-web/models"
	user "github.com/IamNator/mysql-golang-web/session"
	"github.com/IamNator/mysql-golang-web/views"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

/*
	package injection needs more simplicity,
	Migrations seems to be error prone, this needs to be simplified
	User database is not sufficient, We can make use of Google Register/Login API

	Next Modification : Use Google Register/Login API
*/

func main() {
	//dbGeneral := models.DBData{
	//	DBType:       "mysql",                       //Type
	//	User:         "b7e0a0a81fef1f",              //User
	//	Password:     "2e02951d",                    //Password
	//	Host:         "eu-cdbr-west-03.cleardb.net", //Host 3306
	//	DBName:       "heroku_31043c4e11d34ce",      //DBName
	//	Session:      nil,                           //Session for db
	//	SessionToken: make(map[string]models.UserCredentials),    // map[string]struct [token]userDetails
	//}

	//dbGeneral := controllers.DBData{
	//	DBType:   "mysql",          //Type
	//	User:     "root",    	    //User
	//	Password: "299792458m/s",   //Password
	//	Host:     "localhost:3306", //Host 3306
	//	DBName:   "app",  			//DBName
	//	Session:  nil,              //Session
	//	SessionIDs:	make(map[string]string),	//map[string]string
	//	SessionUsers: make(map[string]string),	// map[string]string
	//}
	DB := controllers.Controllersdb(dbGeneral)
	db, _ := DB.OpenDB()
	defer DB.CloseDB()

	dbGeneral.Session = db
	dbUser := user.Sessiondb(dbGeneral)//session

	if !DB.DbExists() {
		CreateAndFillDb(&dbGeneral)
		fmt.Println("Database created and updated")
	}

	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/home", views.Home).Methods("GET")
	myRouter.HandleFunc("/insert", views.Insert).Methods("GET")
	myRouter.HandleFunc("/", views.Index).Methods("GET")
	myRouter.HandleFunc("/login", views.Login).Methods("GET")
	myRouter.HandleFunc("/register", views.Register).Methods("GET")

	myRouter.HandleFunc("/api/fetch", DB.Fetch).Methods("GET")        //use dbData.Fetch_t to test
	myRouter.HandleFunc("/api/update", DB.Update).Methods("POST")     //use dbData.Update_t to test
	myRouter.HandleFunc("/api/delete", DB.Delete).Methods("DELETE")   //use dbData.Delete_t to test

	myRouter.HandleFunc("/api/register", dbUser.Register).Methods("POST") //use dbData.Register_t to test
	myRouter.HandleFunc("/api/login", dbUser.Login).Methods("POST")       //use dbData.Login_t to test

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	go fmt.Printf("server running...@localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, myRouter))

}

func CreateAndFillDb(dbGeneral *models.DBData) {
	dbMigration := migrations.Migrationdb(*dbGeneral)
	dbSeeders := seeders.Seeddb(*dbGeneral)
	dbMigration.CreateUserDb()
	dbMigration.CreatePhoneBookDb()
	dbSeeders.FillUserDb()
	dbSeeders.FillDb()
}
