package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/IamNator/mysql-golang-web/models"
	"github.com/IamNator/mysql-golang-web/session"
	"net/http"
	validate "github.com/go-playground/validator/v10"
)


// takes in req.Body = { "token": "42442-343-3432n-34mv", {contact details} }
//
//contact details = { firstname, lastname, phone_number }
//
func (db *Controllersdb) Update(w http.ResponseWriter, req *http.Request) {

	var reqBody struct {
	 	Token string
	 	models.PhoneBookContact
	}
	json.NewDecoder(req.Body).Decode(&user)

	if user

	validator := validate.New()
	err := validator.Struct(user)

	if err == nil {
		var userid string

		userid = db.SessionToken[token]
		stmt, err := db.Session.Prepare(`INSERT INTO phoneBook (userID, FirstName,LastName,phoneNumber)
	VALUES (?,?,?,?)`)

		_, err = stmt.Exec(userid, user.FirstName, user.LastName, user.PhoneNumber)
		Check(err)

		if err != nil {
			session.JsonError(&w, "Unable to create user Database Error", http.StatusInternalServerError)
		} else {
			fmt.Println("Data Successfully Added")
		}
		fmt.Fprintf(w, `Successful\n`)
	} else {
		fmt.Fprintf(w, `Please fill in all the fields\n`)
	}

}
