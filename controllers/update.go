package controllers

import (
	"encoding/json"
	"github.com/IamNator/mysql-golang-web/models"
	"github.com/IamNator/mysql-golang-web/session"
	validate "github.com/go-playground/validator/v10"
	"net/http"
)


// takes in req.Body = { "token": "42442-343-3432n-34mv", "detail": {contact details} }
//
//contact details = { firstname, lastname, phone_number }
//
//return w.Body =  {  "status": true, "message": "contact added to phone book" }
func (db *Controllersdb) Update(w http.ResponseWriter, req *http.Request) {

	var reqBody struct {
	 	Token string `validate: "required"`
	 	Details models.PhoneBookContact `validate: "required" json:"details"`
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


		stmt, err := db.Session.Prepare(`INSERT INTO phoneBook (FirstName,LastName,phoneNumber)
	VALUES (?,?,?)`)

		_, err = stmt.Exec(reqBody.Details.FirstName, reqBody.Details.LastName, reqBody.Details.PhoneNumber)
		if err != nil {
			session.JsonError(&w, "Unable to create user Database Error", http.StatusInternalServerError)
		} else {
			resp := struct {
				Status bool `json:"status"`
				Message interface{} `json:"message"`
			}{
				true,
				"New contact Added to Phone Book",
			}
			err = json.NewEncoder(w).Encode(resp)
			if err != nil {
				session.JsonError(&w, err.Error(), http.StatusInternalServerError)
			}
		}

}
