package controllers

import (
	"encoding/json"
	"github.com/IamNator/mysql-golang-web/models"
	"github.com/IamNator/mysql-golang-web/session"
	"net/http"
)

// takes req.Body = { "token": "ere-dfd-f3432", "id": "42cv"}
//
//returns w.Body = { "status": "true", "message": [ phone book contacts ] }
func (db *Controllersdb) Fetch(w http.ResponseWriter, req *http.Request) {
	var reqBody struct {
		Token string `json:"token"`
		ID string    `json:"id"`
	}
	json.NewDecoder(req.Body).Decode(&reqBody)

	if _, ok := db.SessionToken[reqBody.Token]; !ok { //Check if user is logged in (id exists in the MAP)
		session.JsonError(&w, "Unauthorized access", http.StatusUnauthorized)
		return
	}

	_ = db.Session.Ping()
	rows, err := db.Session.Query(`SELECT id, FirstName, LastName, PhoneNumber FROM phoneBook WHERE userID=` + db.SessionToken[reqBody.Token].ID)
	Check(err)

	var user models.PhoneBookContact
	var users []models.PhoneBookContact

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.PhoneNumber)
		Check(err)

		users = append(users, user)
	}


	resp := session.MyStdResp{
		Status: true,
		Message: users,
	}

	err = json.NewEncoder(w).Encode(resp) //Sends an array of user information
	if err != nil {
		session.JsonError(&w, err.Error(), http.StatusInternalServerError)
	}

}
