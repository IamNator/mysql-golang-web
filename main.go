package main

import (
	"fmt"
	"github.com/IamNator/mysql-golang-web/controllers"
	"github.com/IamNator/mysql-golang-web/database/migrations"
	"github.com/IamNator/mysql-golang-web/database/seeders"
	"github.com/IamNator/mysql-golang-web/user"
	"github.com/IamNator/mysql-golang-web/views"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)


//"mysql", "root:299792458m/s@tcp(127.0.0.1:3306)/test"

func main() {
	dbGeneral := controllers.DBData{
		DBType:   "mysql",          //Type
		User:     "b7e0a0a81fef1f", //User
		Password: "2e02951d",  //Password
		Host:     "mysql://b7e0a0a81fef1f:2e02951d@eu-cdbr-west-03.cleardb.net/heroku_31043c4e11d34ce?reconnect=true:3306", //Host 3306
		DBName:   "app",           //DBName
		Session:  nil,              //Session
	}

	db, _ := dbGeneral.OpenDB()
	dbGeneral.Session = db

	dbData := controllers.DBData(dbGeneral)
	dbUser := user.DBData(dbGeneral)

    if !dbGeneral.DbExists() {
		CreateAndFillDb(&dbGeneral)
    	fmt.Print("Database created and updated")
	}

	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", views.Index).Methods("GET")

	myRouter.HandleFunc("/api/fetch", dbData.Fetch).Methods("GET")          //use dbData.Fetch_t to test
	myRouter.HandleFunc("/api/update", dbData.Update).Methods("POST")       //use dbData.Update_t to test
	myRouter.HandleFunc("/api/delete", dbData.Delete).Methods("DELETE")     //use dbData.Delete_t to test
	myRouter.HandleFunc("/api/register", dbUser.Register).Methods("DELETE") //use dbData.Register_t to test
	myRouter.HandleFunc("/api/login", dbUser.Login).Methods("DELETE")       //use dbData.Login_t to test

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Printf("server running...@localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, myRouter))

	defer dbData.CloseDB()

}

func CreateAndFillDb(dbGeneral * controllers.DBData){
	dbMigration := migrations.DBData(*dbGeneral)
	dbSeeders := seeders.DBData(*dbGeneral)
	dbMigration.CreateUserDb()
	dbMigration.CreatePhoneBookDb()
	dbSeeders.FillDb()
}

//"database-1.cakv5tpw09ys.eu-west-2.rds.amazonaws.com:3306",
//"3XeaektyhNmPoUqJsifH",
