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
	if user.UserName == "" || user.Password == "" {
		JsonRegisterError(&w, "please Fill in fields", http.StatusBadRequest)
		return
	}
	err := db.Session.QueryRow("Select username From users WHERE username=?", user.UserName).Scan(&user.UserName)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			JsonRegisterError(&w,"server Error, unable to create your account (hash problem)", http.StatusInternalServerError)
			log.Fatal(err)
			return
		}

		_, err = db.Session.Exec("INSERT INTO users(username, password) VALUES(?,?)", user.UserName, hashedPassword)
		if err != nil {
			JsonRegisterError(&w,"server Error, unable to create your account (Database problem)", http.StatusInternalServerError)
			log.Fatal(err)
			return
		}
		//http.SetCookie(w, LoginCookie(user.userName, db))
		w.Write([]byte("User Created!"))
		return
	case err != nil:
		JsonRegisterError(&w,"server error, Unable to access Database",http.StatusInternalServerError)
		fmt.Print(err)
		return
	default:
		JsonRegisterError(&w,"User Already Exists, Please login",http.StatusFound)
	}

}


func (db *DBData) Register_new(w http.ResponseWriter, req *http.Request) {
	var user User
	json.NewDecoder(req.Body).Decode(&user)

	if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.PassWord == "" {
		JsonRegisterError(&w, "please Fill in fields", http.StatusBadRequest)
		return
	}
	err := db.Session.QueryRow("Select email From users WHERE email=?", user.Email).Scan(&user.Email)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PassWord), bcrypt.DefaultCost)
		if err != nil {
			JsonRegisterError(&w,"server Error, unable to create your account (hash problem)", http.StatusInternalServerError)
			log.Fatal(err)
			return
		}

		_, err = db.Session.Exec("INSERT INTO users(firstname, lastname, email, password) VALUES(?,?,?,?)", user.FirstName, user.LastName, user.Email, hashedPassword)
		if err != nil {
			JsonRegisterError(&w,"server Error, unable to create your account (Database problem)", http.StatusInternalServerError)
			log.Fatal(err)
			return
		}
		//http.SetCookie(w, LoginCookie(user.userName, db))
		json.NewEncoder(w).Encode("USer Created")
		return
	case err != nil:
		JsonRegisterError(&w,"server error, Unable to access Database",http.StatusInternalServerError)
		fmt.Print(err)
		return
	default:
		JsonRegisterError(&w,"User Already Exists, Please login",http.StatusFound)
	}

}

func JsonRegisterError(w * http.ResponseWriter, ErrorMessage string, ErrorCode int) {
	(*w).WriteHeader(ErrorCode)
	json.NewEncoder(*w).Encode(struct{
		Login string `json:"login"`
	}{
		ErrorMessage,
	})
}

