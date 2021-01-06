package controllers

import (
	"encoding/json"
	"github.com/IamNator/mysql-golang-web/models"
	"github.com/IamNator/mysql-golang-web/session"
	validate "github.com/go-playground/validator/v10"
	"net/http"
)

// swagger:parameters update
type updateRequestWrapper struct {
	// in: body
	Body updateRequest
}

// swagger:model
type updateRequest struct {
	Token   string                  `json:"token" validate:"required"`
	Details models.PhoneBookContact `json:"details" validate:"required" `
}

// update successful
// swagger:response updateResponse
type updateResponseWrapper struct {
	// in: body
	Body struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}
}

// token not valid, login to get a new one
// swagger:response updateUnauthorized
type updateUnauthorizedWrapper struct {
	// in: body
	Body MyStdResp
}

// unable to respond to request
// swagger:response updateInternalError
type updateInternalErrorWrapper struct {
	// in: body
	Body MyStdResp
}

// swagger:route PUT /api/update controllers update
// adds new contacts to a phoneBook
// responses:
// 200: updateResponse
// 401: updateUnauthorized
// 500: updateInternalError
func (db *Controllersdb) Update(w http.ResponseWriter, req *http.Request) {

	var reqBody updateRequest
	json.NewDecoder(req.Body).Decode(&reqBody)

	validator := validate.New()
	err := validator.Struct(reqBody)
	if err != nil {
		session.JsonError(&w, err.Error(), http.StatusBadRequest)
		return
	}

	Mutex.Lock()
	id, ok := db.SessionToken[reqBody.Token]
	Mutex.Unlock()

	if !ok {
		session.JsonError(&w, "Unauthorized Access, Please login", http.StatusUnauthorized)
		return
	}

	var ph_no string
	err = db.Session.QueryRow("SELECT phonenumber FROM phonebook WHERE phonenumber=?", reqBody.Details.PhoneNumber).Scan(&ph_no)
	if err == nil {
		session.JsonError(&w, "phone number already exists", http.StatusConflict)
		return
	}

	stmt, err := db.Session.Prepare(`INSERT INTO phoneBook (userID, FirstName,LastName,phoneNumber)
	VALUES (?,?,?,?)`)

	_, err = stmt.Exec(id.ID, reqBody.Details.FirstName, reqBody.Details.LastName, reqBody.Details.PhoneNumber)
	if err != nil {
		session.JsonError(&w, "Unable to create user Database Error", http.StatusInternalServerError)
	} else {
		rows, err := db.Session.Query(`SELECT id FROM phoneBook WHERE phoneNumber=` + reqBody.Details.PhoneNumber)
			Check(err)

			var userID string

		err = rows.Scan(&userID)
			Check(err)

		resp := struct {
			Status  bool   `json:"status"`
			Message string `json:"message"`
		}{
			true,
			userID,
		}
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			session.JsonError(&w, err.Error(), http.StatusInternalServerError)
		}
	}
}
