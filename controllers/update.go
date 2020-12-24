package controllers

import (
	"encoding/json"
	"github.com/IamNator/mysql-golang-web/models"
	"github.com/IamNator/mysql-golang-web/session"
	validate "github.com/go-playground/validator/v10"
	"net/http"
)


// takes in req.Body = { "token": "42442-343-3432n-34mv", {contact details} }
//
//contact details = { firstname, lastname, phone_number }
//
func (db *Controllersdb) Update(w http.ResponseWriter, req *http.Request) {

	var reqBody struct {
	 	Token string `validate: "required"`
	 	models.PhoneBookContact `validate: "required"`
	}
	json.NewDecoder(req.Body).Decode(&reqBody)

	validator := validate.New()
	err := validator.Struct(reqBody)
	if err != nil {
		session.JsonError(&w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := db.SessionToken[reqBody.Token]; !ok {
		session.JsonError(&w, "Unauthorized Access, Please login", http.StatusUnauthorized)
		return
	}


		stmt, err := db.Session.Prepare(`INSERT INTO phoneBook (userID, FirstName,LastName,phoneNumber)
	VALUES (?,?,?,?)`)

		_, err = stmt.Exec(db.SessionToken[reqBody.Token].ID, reqBody.FirstName, reqBody.LastName, reqBody.PhoneNumber)
		if err != nil {
			session.JsonError(&w, "Unable to create user Database Error", http.StatusInternalServerError)
		} else {

		}


}
