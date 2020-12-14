package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func (db *DBData) Register(w http.ResponseWriter, req *http.Request) {
	var user RegisterUser
	json.NewDecoder(req.Body).Decode(&user)
	err := db.Session.QueryRow("Select username From users WHERE username=?", user.UserName).Scan(&user.UserName)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "server Error, unable to create your account (hash problem)", 500)
			log.Fatal(err)
			return
		}

		_, err = db.Session.Exec("INSERT INTO users(username, password) VALUES(?,?)", user.UserName, hashedPassword)
		if err != nil {
			http.Error(w, "Server error, unable to create your account (db problem)", 500)
			log.Fatal(err)
			return
		}
		//http.SetCookie(w, LoginCookie(user.userName, db))
		w.Write([]byte("User Created!"))
		return
	case err != nil:
		http.Error(w, "server error, Unable to create your account", 500)
		fmt.Print(err)
		return
	default:
		http.Redirect(w, req, "/", 301)
	}

}
