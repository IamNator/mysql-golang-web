package controllers

import (
	"database/sql"
	"fmt"
	//"golang.org/x/net/html"
	"encoding/json"
	"github.com/IamNator/mysql-golang-web/models"
	"log"
	"net/http"
	//"strconv"
)

type DBData struct {
	DBType, User, Password, Host, DBName string
	Session                              *sql.DB
}

func (db DBData) OpenDB() (*sql.DB, error) {
	return sql.Open(db.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s", db.User, db.Password, db.Host, db.DBName))
}

func (db DBData) CloseDB() string {
	err := db.Session.Close()
	if err != nil {
		return fmt.Sprintf("%v", err)
	} else {
		return fmt.Sprintln("Data base closed")
	}
}

func (db *DBData) Fetch(w http.ResponseWriter, req *http.Request) {

	// db, err := sql.Open("mysql", "root:299792458m/s@tcp(127.0.0.1:3306)/test")
	// check(err)

	db.Session.Ping()

	rows, err := db.Session.Query(`SELECT fname, lname, phone_number, id FROM phonenumber`)
	check(err)

	var user models.User
	var users []models.User

	for rows.Next() {
		err = rows.Scan(&user.Fname, &user.Lname, &user.Phone_number, &user.ID)
		check(err)

		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users) //Sends an array of user information
	//	db.Close()
}

func (db *DBData) Delete(writer http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm() err: %v", err)
		return
	}
	var user models.User
	json.NewDecoder(req.Body).Decode(&user)


	stmt, err := db.Session.Prepare(`DELETE FROM phonenumber WHERE id = ? ;`)

	_, err = stmt.Exec(user.ID)
	check(err)
	//db.Close() //#######################

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode("deleted")

}

func (db *DBData) Update(w http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	var user models.User
	json.NewDecoder(req.Body).Decode(&user)

	if user.Fname != "" && user.Lname != "" && user.Phone_number != "" && string(user.ID) != "" {

		stmt, err := db.Session.Prepare(`INSERT INTO phonenumber (fname,lname,phone_number,id)
	VALUES (?,?,?,?)`)

		_, err = stmt.Exec(user.Fname, user.Lname, user.Phone_number, user.ID)
		check(err)
		//	db.Close() //#######################

		if err != nil {
			fmt.Fprintln(w, err)
		} else {
			fmt.Println("\nData Successfully Added")
		}

		fmt.Fprintf(w, `Successful`)
	} else {
		fmt.Fprintf(w, `Please fill in all the fields `)
	}

}

func check(err error) {
	if err != nil {
		log.Printf("%+v\n", err)
	}
}
