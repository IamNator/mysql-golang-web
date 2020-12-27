// Package classification Login API
//
//Documentation for register API
//
// schemes: http
// BasePath: /
// Version: 1.0.0
// Contact: natverior1@gmail.com
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package session

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/IamNator/mysql-golang-web/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// swagger:response registerResponse
type registerResponseWrapper struct {
//	in:body
	Body MyStdResp
}

// swagger:route POST /api/register session register
// adds new user to user database
// responses:
// 200: registerResponse
func (db *Sessiondb) Register(w http.ResponseWriter, req *http.Request) {
	var user models.UserCredentials
	_ = json.NewDecoder(req.Body).Decode(&user)

	if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.PassWord == "" {
		JsonError(&w, "please Fill in fields", http.StatusBadRequest)
		return
	}
	err := db.Session.QueryRow("Select email From users WHERE email=?", user.Email).Scan(&user.Email)

	switch {

	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PassWord), bcrypt.DefaultCost)
		if err != nil {
			JsonError(&w, fmt.Sprintf("server Error, unable to create your account (hash problem) : %v ", err), http.StatusInternalServerError)
			log.Fatal(err)
			return
		}

		_, err = db.Session.Exec("INSERT INTO users(firstname, lastname, email, password) VALUES(?,?,?,?)", user.FirstName, user.LastName, user.Email, hashedPassword)
		if err != nil {
			JsonError(&w, fmt.Sprintf("server Error, unable to create your account (Database problem) : %v ", err), http.StatusInternalServerError)
			log.Fatal(err)
			return
		}
		//http.SetCookie(w, LoginCookie(user.userName, db))
		res := MyStdResp{
			Status: true,
			Message: "User Created",
		}
		_ = json.NewEncoder(w).Encode(res)
		return


	case err != nil:
		JsonError(&w, "server error, Unable to access Database", http.StatusInternalServerError)
		fmt.Print(err)
		return


	default:
		JsonError(&w, "User Already Exists, Please login", http.StatusFound)
	}
}
