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

	myRouter := mux.NewRouter()
	fileServer := http.FileServer(http.Dir("./"))

	fmt.Println("server running...@localhost:9080")
	myRouter.HandleFunc("/index", views.Index).Methods("GET")
	myRouter.HandleFunc("/insert", views.Insert).Methods("GET")
	myRouter.Handle("/css/bootstrap.min.css", fileServer)
	myRouter.Handle("/js/bootstrap.min.js", fileServer)

	myRouter.HandleFunc("/api/fetch", controllers.Fetch).Methods("GET")
	myRouter.HandleFunc("/api/update", controllers.Update).Methods("POST")
	myRouter.HandleFunc("/api/delete", controllers.Delete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9080", myRouter))

}
