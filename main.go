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

// mt.Sprintf("%s:%s@tcp(%s)/%s", db.User, db.Password, db.Host, db.DBName))

// db, err := sql.Open("mysql", "root:299792458m/s@tcp(127.0.0.1:3306)/test")
// check(err)

func main() {
	dbData := controllers.DBData{
		"mysql",          //Type
		"admin",           //User
		"3XeaektyhNmPoUqJsifH",   //Password
		"database-1.cakv5tpw09ys.eu-west-2.rds.amazonaws.com:3306", //Host
		"test",           //DBName
		nil,              //Session
	}

	// db, _ := dbData.OpenDB()
	// dbData.Session = db

	myRouter := mux.NewRouter()
	go fmt.Println("server running...@localhost:9080")
<<<<<<< HEAD
	myRouter.HandleFunc("/", views.Index).Methods("GET")
	myRouter.HandleFunc("/insert", views.Insert).Methods("GET")
=======
	myRouter.HandleFunc("/index", views.Index).Methods("GET")
>>>>>>> 9cf9367f51d5c2900d3ce236a6101e1238149c87


	myRouter.HandleFunc("/api/fetch", dbData.Fetch).Methods("GET") //use dbData.Fetch_t to test
	myRouter.HandleFunc("/api/update", dbData.Update).Methods("POST") //use dbData.Update_t to test
	myRouter.HandleFunc("/api/delete", dbData.Delete).Methods("DELETE")//use dbData.Delete_t to test
	log.Fatal(http.ListenAndServe(":9080", myRouter))

	//defer dbData.CloseDB()

}
