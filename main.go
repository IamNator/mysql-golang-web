package main

import (
	"fmt"
	"github.com/IamNator/mysql-golang-web/controllers"
	"github.com/IamNator/mysql-golang-web/views"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	dbData := controllers.DBData{
		DBType: "mysql",          //Type
		User: "admin",           //User
		Password: "", //"3XeaektyhNmPoUqJsifH",   //Password
		Host: "",//"database-1.cakv5tpw09ys.eu-west-2.rds.amazonaws.com:3306", //Host
		DBName: "test",           //DBName
		Session: nil,              //Session
	}

	 db, _ := dbData.OpenDB()
	 dbData.Session = db

	myRouter := mux.NewRouter()
	go fmt.Println("server running...@localhost:9080")
	myRouter.HandleFunc("/", views.Index).Methods("GET")


	myRouter.HandleFunc("/api/fetch", dbData.Fetch_t).Methods("GET") //use dbData.Fetch_t to test
	myRouter.HandleFunc("/api/update", dbData.Update_t).Methods("POST") //use dbData.Update_t to test
	myRouter.HandleFunc("/api/delete", dbData.Delete_t).Methods("DELETE")//use dbData.Delete_t to test
	log.Fatal(http.ListenAndServe(":9080", myRouter))

	defer dbData.CloseDB()

}
