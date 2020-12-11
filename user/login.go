package user

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

func (db *DBData) Login(w http.ResponseWriter, req *http.Request) {
	var user LoginUser
	var userdb RegisterUser
	var id string
	err := json.NewDecoder(req.Body).Decode(&user)
	check(err)

	fmt.Println(user.UserName)

	err = db.Session.QueryRow("SELECT id, username, password FROM users WHERE username=?", user.UserName).Scan(&id, &userdb.UserName, &userdb.Password)
	if err != nil {
		fmt.Println(err)
		//http.Redirect(w, req, "/", 301)
		//http.Write(http.StatusNotFound)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userdb.Password), []byte(user.Password))
	if err != nil {
		//http.Redirect(w, req, "/", 301)
		http.Error(w, "Password incorrect", http.StatusNotFound)
		return
	}

	http.SetCookie(w, LoginCookie(id, user.UserName, db))
	w.Write([]byte(userdb.UserName + "logged in successfully"))
}

func LoginCookie(ID, username string, db *DBData) *http.Cookie {
	expire := time.Now().AddDate(0, 0, 1)
	id := uuid.NewV4().String()
	db.SessionIDs[id] = username
	db.SessionUsers[username] = ID
	fmt.Println(id)
	return &http.Cookie{
		Name:    "sessionID",
		Value:   id,
		Expires: expire,
		//Secure: true,
	}
}

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}
