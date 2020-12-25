package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/IamNator/mysql-golang-web/session"
	"log"

	//"fmt"
	"net/http"
)

//body takes ("token": "342-342s-fsd-343cv", "id": "23" }
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


	stmt, err := db.Session.Prepare(`DELETE FROM phoneBook WHERE id = ? AND userID = ? ;`)
	res, err := stmt.Exec(user.ID, masterID)
	log.Println(res)
	Check(err)

	writer.Header().Set("Content-Type", "application/json")
	resp := session.MyStdResp{
		Status: true,
		Message: "User Deleted",
	}

	err = json.NewEncoder(writer).Encode(resp)
	if err != nil {
		session.JsonError(&writer, err.Error(), http.StatusInternalServerError)
	}

}
