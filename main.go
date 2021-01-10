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
	"context"
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
	"os/signal"
	"time"
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
		DBType:       config.DBType,                           //Type
		User:         config.User,                             //User
		Password:     config.Password,                         //Password
		Host:         config.Host,                             //Host 3306
		DBName:       config.DBName,                           //DBName
		Session:      nil,                                     //Session for db
		SessionToken: make(map[string]models.UserCredentials), // map[string]struct [token]userDetails
	}

	DB := controllers.Controllersdb(dbGeneral)
	db, _ := DB.OpenDB()
	DB.Session = db
	defer DB.CloseDB()

	dbUser := user.Sessiondb(DB) //session

	// checks if DB table exist
	if !DB.DbExists() {
		CreateAndFillDb(DB)
		fmt.Println("Database created and updated")
	}

	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/home", views.Home).Methods("GET")
	myRouter.HandleFunc("/insert", views.Insert).Methods("GET")
	myRouter.HandleFunc("/index", views.PhoneBook).Methods("GET")
	myRouter.HandleFunc("/", views.Login).Methods("GET")
	myRouter.HandleFunc("/register", views.Register).Methods("GET")

	myRouter.HandleFunc("/api/contacts", DB.Fetch).Methods("GET")      //fetches contact informations
	myRouter.HandleFunc("/api/contacts", DB.Update).Methods("PUT")    // adds new contact informations
	myRouter.HandleFunc("/api/contacts", DB.Delete).Methods("DELETE") // delete a contact

	myRouter.HandleFunc("/user/register", dbUser.Register).Methods("POST") //use dbData.Register_t to test
	myRouter.HandleFunc("/user/login", dbUser.Login).Methods("POST")       //use dbData.Login_t to test
	myRouter.HandleFunc("/user/logout", dbUser.Logout).Methods("POST")     //use dbData.Login_t to test

	// "/auth/google/callback"
	//reads swagger.yaml doc file for APIs
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	myRouter.Handle("/docs", sh)
	myRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	//Seems heroku has .env file { need to add DB credentials to this file too }
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}
	//go fmt.Printf("Number of CPU : %d \n", runtime.NumCPU())
	//go fmt.Printf("Number of Goroutine : %d \n", runtime.NumGoroutine())

	myserver := http.Server{
		Addr: ":"+port,
		Handler: myRouter,
		IdleTimeout: 20*time.Second, //increase this esp when running as a microservice
		ReadTimeout: 20*time.Second,
		WriteTimeout: 45*time.Second,
	}

	go func(){
		go fmt.Printf("server running...@localhost:%s\n", port)
		log.Fatal(myserver.ListenAndServe())
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig :=<- sigChan
	log.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	myserver.Shutdown(tc)
	close(sigChan)

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

//extracts important information from app.env file { I should gitignore this file}
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
