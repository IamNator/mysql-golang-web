package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/IamNator/mysql-golang-web/models"
	"net/http"
)

func (db *DBData) Fetch(w http.ResponseWriter, req *http.Request) {
	var SessionUserID string
	cookie, err := req.Cookie("sessionID")
	if err == http.ErrNoCookie {
		//http.Redirect(w, req, "/", 301)
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode("Cookie not found")
		fmt.Println("Cookie not found")
		return

	}
	userName := db.SessionIDs[cookie.Value]      //returns the username
	if id, ok := db.SessionUsers[userName]; ok { //Check if user is logged in (id exists in the MAP)
		SessionUserID = id
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode("Cookie Value invalid, please login")
		fmt.Println("Cookie.Value does not match userNAme")
		return
	}

	_ = db.Session.Ping()
	rows, err := db.Session.Query(`SELECT id, FirstName, LastName, PhoneNumber FROM phoneBook WHERE userID=` + SessionUserID)
	Check(err)

	var user models.User
	var users []models.User

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.PhoneNumber)
		Check(err)

		users = append(users, user)
	}
	_ = json.NewEncoder(w).Encode(users) //Sends an array of user information

}
