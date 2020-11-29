package user

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)


func (db * DBData) Login(w http.ResponseWriter, req * http.Request){
	var user LoginUser
	var userdb RegisterUser
	json.NewDecoder(req.Body).Decode(&user)

	err:= db.Session.QueryRow("SELECT username, password FROM users WHERE username=?", user.userName).Scan(&userdb.userName, &userdb.Password)
	if err != nil {
		http.Redirect(w, req, "/register", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userdb.Password), []byte(user.Password))
	if err != nil {
		http.Redirect(w, req, "/register", 301)
		return
	}

	w.Write([]byte(userdb.userName + "logged in successfully"))
}