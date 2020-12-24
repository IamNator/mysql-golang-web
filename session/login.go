package session

import (
	"encoding/json"
	"fmt"
	"github.com/IamNator/mysql-golang-web/models"
	//"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	//"time"
)

//Takes in {username, password}
//
// Returns { {code, userDetails}, token }
//
//userDetails = {id, firstname, lastname, email, password}
func (db *Sessiondb) Login(w http.ResponseWriter, req *http.Request) {
	var user LoginCredentials
	var userDb models.UserCredentials
	err := json.NewDecoder(req.Body).Decode(&user) //fills up user from body
	check(err)


	if user.Email == "" || user.PassWord == "" {
		JsonError(&w, "please Fill in fields", http.StatusBadRequest)
		return
	}

	err = db.Session.QueryRow("SELECT id, firstname, lastname, email, password FROM users WHERE email=?", user.Email).Scan(&userDb.ID, &userDb.FirstName, &userDb.LastName, &userDb.Email, &userDb.PassWord)
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

	token := CreateToken(db, userDb)
	w.WriteHeader(http.StatusOK

	res := struct {
		MyStdResp
		Token string `json:"token"`
	}{MyStdResp{true, userDb },
		token,
	}
	err = json.NewEncoder(w).Encode(res)
	check(err)
}

func JsonError(w *http.ResponseWriter, ErrorMessage string, ErrorCode int) {
	(*w).WriteHeader(ErrorCode)
	res := MyStdResp{
		Status: false,
		Message: ErrorMessage,
	}
	json.NewEncoder(*w).Encode(res)
}

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}
