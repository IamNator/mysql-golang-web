package user

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
	"net/http"
)


func (db * DBData) Login(w http.ResponseWriter, req * http.Request){
	var user LoginUser
	var userdb RegisterUser
	json.NewDecoder(req.Body).Decode(&user)

	id := uuid.NewV4().String()
	fmt.Println(id)

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

	http.SetCookie(w, &http.Cookie{
		Name: "session_id",
		Value: id+userdb.Password.String(),
	})
	w.Write([]byte(userdb.userName + "logged in successfully"))
}