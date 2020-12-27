// Package classification delete API
//
//Documentation for delete API
//
// schemes: http
// BasePath: /
// Version: 1.0.0
// Contact: natverior1@gmail.com
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/IamNator/mysql-golang-web/session"
	"log"

	//"fmt"
	"net/http"
)

// swagger:response deleteResponse
type deleteResponseWrapper struct {
	// in: body
	Body MyStdResp
}

// swagger:model
type MyStdResp struct {
	// successful / not successful
	Status bool `json:"status"`
	// delete message / error message when delete is unsuccessful
	Message string `json:"message"`
}

// swagger:route POST /api/delete controllers delete
// delete contact from phone book
// responses:
// 200: deleteResponse
func (db *Controllersdb) Delete(writer http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		session.JsonError(&writer, fmt.Sprintf("ParseForm()  err : %v", err), http.StatusBadRequest)
		//fmt.Fprintf(writer, "ParseForm() err: %v", err)
		return
	}

	var user struct {
		Token string `json:"token"`
		ID    string `json:"id"` //id to be deleted
	}

	_ = json.NewDecoder(req.Body).Decode(&user)

	if _, ok := db.SessionToken[user.Token]; !ok {
		session.JsonError(&writer, "Unauthorized access please login", http.StatusUnauthorized)
		return
	}

	masterID := db.SessionToken[user.Token] //The person authorizing the delete

	stmt, err := db.Session.Prepare(`DELETE FROM phoneBook WHERE id = ? AND userID = ? ;`)
	res, err := stmt.Exec(user.ID, masterID)
	log.Println(res)
	Check(err)

	writer.Header().Set("Content-Type", "application/json")
	resp := MyStdResp{
		Status:  true,
		Message: "User Deleted",
	}

	err = json.NewEncoder(writer).Encode(resp)
	if err != nil {
		session.JsonError(&writer, err.Error(), http.StatusInternalServerError)
	}

}
