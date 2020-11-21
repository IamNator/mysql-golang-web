package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/IamNator/mysql-golang-web/models"
	"log"
	"net/http"
	"os"
	//"sync"
)

// //const GlobalDB := "mysql","user:password@tcp(127.0.0.1:3306)/hello"
// type DBData struct {
// 	DBType, User, Password, Host, DBName string
// 	Session                              *sql.DB
// }

// func (db DBData) OpenDB() (*sql.DB, error) {
// 	return sql.Open(db.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s", db.User, db.Password, db.Host, db.DBName))
// }

// func (db DBData) CloseDB() string {
// 	err := db.Session.Close()
// 	if err != nil {
// 		return fmt.Sprintf("%v", err)
// 	} else {
// 		return fmt.Sprintln("Data base closed")
// 	}
// }

func (db *DBData) Fetch_t(w http.ResponseWriter, req *http.Request) {

	// db, err := sql.Open("mysql", "root:299792458m/s@tcp(127.0.0.1:3306)/test")
	// check(err)
	file, _ := os.Open("data.json")
	defer file.Close()

	//var user models.User
	var users []models.User

	json.NewDecoder(file).Decode(&users)

	// for rows.Next() {
	// 	err = rows.Scan(&user.Fname, &user.Lname, &user.Phone_number, &user.ID)
	// 	check(err)

	// 	users = append(users, user)
	// }
	json.NewEncoder(w).Encode(&users) //Sends an array of user information
	log.Println("Data fetched")
	//	db.Close()
}

func (db *DBData) Delete_t(writer http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm() err: %v", err)
		return
	}

	//del_id := req.FormValue("id")

	writer.Header().Set("Content-Type", "application/json")
	s := "{\"deleted\":\"successfully\"}"
	json.NewEncoder(writer).Encode(s)

}

func (db *DBData) Update_t(w http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	file, _ := os.OpenFile("data.json", os.O_WRITE, os.ModePerm)
	defer file.Close()

	var user models.User
	var users []models.User

	json.NewDecoder(req.Body).Decode(&user)
	json.NewDecoder(file).Decode(&users)

	users = append(users, user)

	if user.Fname != "" && user.Lname != "" && user.Phone_number != "" && string(user.ID) != "" {
		json.NewEncoder(file).Encode(&users)
		fmt.Println("\nData Successfully Added")
		fmt.Fprintf(w, `Successful`)
	} else {
		fmt.Fprintf(w, `Please fill in all the fields `)
	}

}
