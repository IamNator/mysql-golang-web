package controllers

import (
	"database/sql"
	"fmt"
	//"golang.org/x/net/html"
	"encoding/json"
	"github.com/IamNator/mysql-golang-web/models"
	"log"
	"net/http"
	"strconv"
)

func Fetch(w http.ResponseWriter, req *http.Request) {

	db, err := sql.Open("mysql", "root:299792458m/s@tcp(127.0.0.1:3306)/test")
	check(err)

	err = db.Ping()
	check(err)

	rows, err := db.Query(`SELECT fname, lname, phone_number, id FROM phonenumber`)
	check(err)

	var user models.User
	var users []models.User

	for rows.Next() {
		err = rows.Scan(&user.Fname, &user.Lname, &user.Phone_number, &user.ID)
		check(err)

		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users) //Sends an array of user information
	db.Close()
}

func Delete(writer http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm() err: %v", err)
		return
	}

	del_id := req.FormValue("Del_id")

	db, err := sql.Open("mysql", "root:299792458m/s@tcp(127.0.0.1:3306)/test") //##################
	check(err)

	stmt, err := db.Prepare(`DELETE FROM phonenumber WHERE id = ? ;`)

	_, err = stmt.Exec(del_id)
	check(err)
	db.Close() //#######################

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode("\"delete\":\"successful\"")

}

func Update(w http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	var user models.User
	user.Fname = req.FormValue("fname")
	user.Lname = req.FormValue("lname")
	user.Phone_number = req.FormValue("phone_number")
	user.ID, _ = strconv.Atoi(req.FormValue("id"))

	if req.FormValue("fname") != "" && req.FormValue("lname") != "" && req.FormValue("phone_number") != "" && req.FormValue("id") != "" {

		db, err := sql.Open("mysql", "root:299792458m/s@tcp(127.0.0.1:3306)/test") //##################
		check(err)

		stmt, err := db.Prepare(`INSERT INTO phonenumber (fname,lname,phone_number,id)
	VALUES (?,?,?,?)`)

		_, err = stmt.Exec(user.Fname, user.Lname, user.Phone_number, user.ID)
		check(err)
		db.Close() //#######################

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

func Insert(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(w, `Successful`)
}

func check(err error) {
	if err != nil {
		log.Printf("%+v\n", err)
	}
}
