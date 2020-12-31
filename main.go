// Package classification mysql-golang-web API
//
// Documentation for mysql-golang-web API
//
// schemes: http
// BasePath: /
// Version: 1.0.0
// Contact: natverior1@gmail.com
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package main

import (
	"fmt"
	"github.com/IamNator/mysql-golang-web/controllers"
	"github.com/IamNator/mysql-golang-web/database/migrations"
	"github.com/IamNator/mysql-golang-web/database/seeders"
	"github.com/IamNator/mysql-golang-web/models"
	user "github.com/IamNator/mysql-golang-web/session"
	"github.com/IamNator/mysql-golang-web/views"
	"github.com/go-openapi/runtime/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
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
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dbGeneral := models.DBData{
		DBType:       config.DBType,                                 //Type
		User:         config.User,                        //User
		Password:     config.Password,                              //Password
		Host:         config.Host,           //Host 3306
		DBName:       config.DBName,                //DBName
		Session:      nil,                                     //Session for db
		SessionToken: make(map[string]models.UserCredentials), // map[string]struct [token]userDetails
	}

	//dbGeneral := models.DBData{
	//	DBType:   "mysql",          //Type
	//	User:     "root",    	    //User
	//	Password: "299792458m/s",   //Password
	//	Host:     "localhost:3306", //Host 3306
	//	DBName:   "app",  			//DBName
	//	Session:  nil,              //Session
	//	SessionToken: make(map[string]models.UserCredentials),	// map[string]string
	//}

	DB := controllers.Controllersdb(dbGeneral)
	db, _ := DB.OpenDB()
	DB.Session = db
	defer DB.CloseDB()

	dbUser := user.Sessiondb(DB) //session

	if !DB.DbExists() {
		CreateAndFillDb(DB)
		fmt.Println("Database created and updated")
	}

	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/home", views.Home).Methods("GET")
	myRouter.HandleFunc("/insert", views.Insert).Methods("GET")
	myRouter.HandleFunc("/index", views.Index).Methods("GET")
	myRouter.HandleFunc("/", views.Login).Methods("GET")
	myRouter.HandleFunc("/register", views.Register).Methods("GET")

	myRouter.HandleFunc("/api/fetch", DB.Fetch).Methods("POST")      //use dbData.Fetch_t to test
	myRouter.HandleFunc("/api/update", DB.Update).Methods("PUT")   //use dbData.Update_t to test
	myRouter.HandleFunc("/api/delete", DB.Delete).Methods("DELETE") //use dbData.Delete_t to test

	myRouter.HandleFunc("/api/register", dbUser.Register).Methods("POST") //use dbData.Register_t to test
	myRouter.HandleFunc("/api/login", dbUser.Login).Methods("POST")       //use dbData.Login_t to test
	myRouter.HandleFunc("/api/logout", dbUser.Logout).Methods("POST")     //use dbData.Login_t to test

	// "/auth/google/callback"

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	myRouter.Handle("/docs", sh)
	myRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	go fmt.Printf("server running...@localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, myRouter))

}

//Creates and fills up database tables if they don't exist
func CreateAndFillDb(db controllers.Controllersdb) {

	//package (mysql) injection happens here
	dbMigration := migrations.Migrationdb(db)
	dbSeeders := seeders.Seeddb(db)

	//Create database tables
	dbMigration.CreateUserDb()
	dbMigration.CreatePhoneBookDb()

	//fill up database with dummy data
	dbSeeders.FillUserDb()
	dbSeeders.FillDb()
}

//extracts important information from app.env file { I should gitignore this
func LoadConfig(path string) (config models.DBData, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)

	return
}