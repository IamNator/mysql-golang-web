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
		"root",           //User
		"299792458m/s",   //Password
		"127.0.0.1:3306", //Host
		"test",           //DBName
		nil,              //Session
	}

	// db, _ := dbData.OpenDB()
	// dbData.Session = db

	myRouter := mux.NewRouter()
	go fmt.Println("server running...@localhost:9080")
	myRouter.HandleFunc("/", views.Index).Methods("GET")
	myRouter.HandleFunc("/insert", views.Insert).Methods("GET")


	myRouter.HandleFunc("/api/fetch", dbData.Fetch_t).Methods("GET")
	myRouter.HandleFunc("/api/update", dbData.Update_t).Methods("POST")
	myRouter.HandleFunc("/api/delete", dbData.Delete_t).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9080", myRouter))

	//defer dbData.CloseDB()

}
