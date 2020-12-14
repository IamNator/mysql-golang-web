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
	var userDb RegisterUser
	var id string
	err := json.NewDecoder(req.Body).Decode(&user)
	check(err)

	fmt.Println(user.UserName)

	err = db.Session.QueryRow("SELECT id, username, password FROM users WHERE username=?", user.UserName).Scan(&id, &userDb.UserName, &userDb.Password)
	if err != nil {
		fmt.Printf("dbQuery Error %v \n", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(user.Password))
	if err != nil {
		fmt.Printf("CompareHashPassword Error %v \n", err)
		http.Error(w, "Password incorrect", http.StatusNotFound)
		return
	}

	http.SetCookie(w, LoginCookie(id, user.UserName, db))
	loginSuccess := struct{
		Login string `json:"login"`
	}{
		"Successful",
	}
	err = json.NewEncoder(w).Encode(loginSuccess)
	check(err)
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
