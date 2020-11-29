package main


import (
	"fmt"
	"github.com/IamNator/mysql-golang-web/controllers"
	"github.com/IamNator/mysql-golang-web/user"
	"github.com/IamNator/mysql-golang-web/views"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	dbData := controllers.DBData{
		DBType: "mysql",          //Type
		User: "admin",            //User
		Password: "",             //Password
		Host: "",				  //Host
		DBName: "test",          //DBName
		Session: nil,            //Session
	}

	//db, _ := dbData.OpenDB()
	//dbData.Session = db
	dbUser := user.DBData(dbData)

	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", views.Index).Methods("GET")

	myRouter.HandleFunc("/api/fetch", dbData.Fetch_t).Methods("GET")      //use dbData.Fetch_t to test
	myRouter.HandleFunc("/api/update", dbData.Update_t).Methods("POST")   //use dbData.Update_t to test
	myRouter.HandleFunc("/api/delete", dbData.Delete_t).Methods("DELETE") //use dbData.Delete_t to test
	myRouter.HandleFunc("/api/register", dbUser.Register).Methods("DELETE") //use dbData.Register_t to test
	myRouter.HandleFunc("/api/login", dbUser.Login).Methods("DELETE") //use dbData.Login_t to test

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Printf("server running...@localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, myRouter))

	// defer dbData.CloseDB()

}



//"database-1.cakv5tpw09ys.eu-west-2.rds.amazonaws.com:3306",
//"3XeaektyhNmPoUqJsifH",
