package controllers

import (
	"encoding/json"
	"github.com/IamNator/mysql-golang-web/models"
	"github.com/IamNator/mysql-golang-web/session"
	validate "github.com/go-playground/validator/v10"
	"net/http"
)

// swagger:response updateResponse
type updateResponseWrapper struct {
	// in: body
	Body struct {
		Status  bool        `json:"status"`
		Message string `json:"message"`
	}
}

// swagger:route POST /api/update controllers update
// adds new contacts to a phoneBook
// responses:
// 200: updateResponse
func (db *Controllersdb) Update(w http.ResponseWriter, req *http.Request) {

	var reqBody struct {
		Token   string                  `json: "token" validate: "required"`
		Details models.PhoneBookContact `json:"details" validate: "required" `
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

	_, err = stmt.Exec(db.SessionToken[reqBody.Token].ID, reqBody.Details.FirstName, reqBody.Details.LastName, reqBody.Details.PhoneNumber)
	if err != nil {
		session.JsonError(&w, "Unable to create user Database Error", http.StatusInternalServerError)
	} else {
		resp := struct {
			Status  bool        `json:"status"`
			Message string `json:"message"`
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
