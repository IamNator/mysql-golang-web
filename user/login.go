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
	if user.UserName == "" || user.Password == "" {
		JsonLoginError(&w, "please Fill in fields", http.StatusBadRequest)
		return
	}

	err = db.Session.QueryRow("SELECT id, username, password FROM users WHERE username=?", user.UserName).Scan(&id, &userDb.UserName, &userDb.Password)
	if err != nil {
		fmt.Printf("dbQuery Error %v \n", err)
		JsonLoginError(&w,"User Not Found", http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(user.Password))
	if err != nil {
		fmt.Printf("CompareHashPassword Error %v \n", err)
		JsonLoginError(&w,"Password Incorrect", http.StatusNotFound)
		return
	}

	http.SetCookie(w, LoginCookie(id, user.UserName, db))
	//loginSuccess := struct{
	//	Login string `json:"login"`
	//}{
	//	"Successful",
	//}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user.UserName)
	check(err)
}


func (db *DBData) Login_new(w http.ResponseWriter, req *http.Request) {
	var user User
	var userDb User
	var id string
	err := json.NewDecoder(req.Body).Decode(&user)
	check(err)

	fmt.Println(user.FirstName)
	if user.Email == "" || user.PassWord == "" {
		JsonLoginError(&w, "please Fill in fields", http.StatusBadRequest)
		return
	}

	err = db.Session.QueryRow("SELECT id, firstname, lastname, email, password FROM users WHERE email=?", user.Email).Scan(&id, &userDb.FirstName, &userDb.LastName, &userDb.Email, &userDb.PassWord)
	if err != nil {
		fmt.Printf("dbQuery Error %v \n", err)
		JsonLoginError(&w,"User Not Found", http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDb.PassWord), []byte(user.PassWord))
	if err != nil {
		fmt.Printf("CompareHashPassword Error %v \n", err)
		JsonLoginError(&w,"Password Incorrect", http.StatusNotFound)
		return
	}

	http.SetCookie(w, LoginCookie(id, user.Email, db))
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(userDb)
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

func JsonLoginError(w * http.ResponseWriter, ErrorMessage string, ErrorCode int) {
	(*w).WriteHeader(ErrorCode)
	json.NewEncoder(*w).Encode(struct{
		Login string `json:"login"`
	}{
		ErrorMessage,
	})
}

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}
