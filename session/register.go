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

// user successfully created
// swagger:response registerResponse
type registerResponseWrapper struct {
	//	in:body
	Body MyStdResp
}

// when a user already exists
// swagger:response registerUserExist
type registerUserExistWrapper struct {
	//	in:body
	Body MyStdResp
}

// returns when there is an internal server error (likely unable access database)
// swagger:response registerInternalError
type registerInternalErrorWrapper struct {
	//	in:body
	Body MyStdResp
}

// swagger:parameters register
type registerRequestWrapper struct {
	// in: body
	Body models.UserCredentials
}

// adds new user to user database
// swagger:route POST /api/register session register
// responses:
// 201: registerResponse
// 302: registerUserExist
// 500: registerInternalError
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
		w.WriteHeader(http.StatusCreated)
		res := MyStdResp{
			Status:  true,
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
