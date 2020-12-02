package user

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
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

	http.SetCookie(w, LoginCookie(user.userName, db))
	w.Write([]byte(userdb.userName + "logged in successfully"))
}

func LoginCookie(username string, db * DBData) * http.Cookie {
	expire := time.Now().AddDate(0,0,1)
	id := uuid.NewV4().String()
	db.SessionIDs[id] = username
	db.SessionUsers[username] = "loggedIN"
	fmt.Println(id)
	return &http.Cookie{
		Name: "sessionID",
		Value: id,
		Expires: expire,
		//Secure: true,
	}
}