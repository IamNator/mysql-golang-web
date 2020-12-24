package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/IamNator/mysql-golang-web/session"

	//"fmt"
	"net/http"
)

//body takes ("token": "342-342s-fsd-343cv", "id": "23", "id": "dfd-434-cvd"}
//
//returns in body ("status": "true", "message": "deleted id, first and last name"
//
func (db *Controllersdb) Delete(writer http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		session.JsonError(&writer, fmt.Sprintf("ParseForm()  err : %v",err), http.StatusBadRequest )
		//fmt.Fprintf(writer, "ParseForm() err: %v", err)
		return
	}

	var user struct {
		Token string `json:"token"`
		ID string 	 `json:"id"` //id to be deleted
	}

	_ = json.NewDecoder(req.Body).Decode(&user)

	if _, ok := db.SessionToken[user.Token]; !ok {
		session.JsonError(&writer, "Unauthorized access please login", http.StatusUnauthorized)
		return
	}

	masterID := db.SessionToken[user.Token]  //The person authorizing the delete

	stmt, err := db.Session.Prepare(`DELETE FROM phoneBook WHERE id = ?, userID = ? ;`)
	_, err = stmt.Exec(user.ID, masterID)
	Check(err)

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode("deleted")

}
