package session

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
	var user LoginCredentials
	var userDb Credentials
	var id string
	err := json.NewDecoder(req.Body).Decode(&user)
	check(err)

	fmt.Println(user.Email)
	if user.Email == "" || user.PassWord == "" {
		JsonError(&w, "please Fill in fields", http.StatusBadRequest)
		return
	}

	err = db.Session.QueryRow("SELECT id, firstname, lastname, email, password FROM users WHERE email=?", user.Email).Scan(&id, &userDb.FirstName, &userDb.LastName, &userDb.Email, &userDb.PassWord)
	if err != nil {
		fmt.Printf("dbQuery Error %v \n", err)
		JsonError(&w, "User Not Found", http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDb.PassWord), []byte(user.PassWord))
	if err != nil {
		fmt.Printf("CompareHashPassword Error %v \n", err)
		JsonError(&w, "Password Incorrect", http.StatusNotFound)
		return
	}

	token, _ := CreateToken(id)
	http.SetCookie(w, LoginCookie(token, user.Email, db))
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(userDb)
	check(err)
}

func LoginCookie(token, username string, db *DBData) *http.Cookie {
	expire := time.Now().AddDate(0, 0, 1)
	id := uuid.NewV4().String()
	db.SessionIDs[id] = username
	db.SessionToken[username] = token
	fmt.Println(token)
	return &http.Cookie{
		Name:    "sessionID",
		Value:   token,
		Expires: expire,
		HttpOnly: true,
		//Secure: true,
	}
}

func JsonError(w *http.ResponseWriter, ErrorMessage string, ErrorCode int) {
	(*w).WriteHeader(ErrorCode)
	json.NewEncoder(*w).Encode(struct {
		Error string `json:"error"`
	}{
		Error: ErrorMessage,
	})
}

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}
