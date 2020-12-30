// Package classification Login API
//
// Documentation for login API
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
	"encoding/json"
	"fmt"
	"github.com/IamNator/mysql-golang-web/models"

	//"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	//"time"
)

// swagger:parameters idOfLoginEndpoint
type loginRequestWrapper struct {
	// in: body
	Body LoginCredentials
}

// Successful login request
// swagger:response loginResponse
type LoginResponseWrapper struct {
	//in: body
	Body LoginResponse
}

// Non registered user tries to login
// swagger:response loginNotFound
type LoginResponseNotFoundWrapper struct {
	//returns when a non registered user tries to login
	//in: body
	Body MyStdResp
}

// request does not conform to login api
// swagger:response loginBadRequest
type LoginBadRequestWrapper struct {
	//returns when the request does not conform to login api
	//in: body
	Body MyStdResp
}

// Incorrect password
// swagger:response loginIncorrectPassword
type LoginIncorrectPasswordWrapper struct {
	//returns when user enters an incorrect password
	//in: body
	Body MyStdResp
}

// LoginResponse defines the struct for API login
// swagger:model
type LoginResponse struct {
	// shows if process was successful or not
	Status bool `json:"status"`
	// Contains user information
	Message struct {
		// the user's first name
		FirstName string `json:"firstname" validate:"required"`
		// the user's last name
		LastName string `json:"lastname" validate:"required"`
		// the user's email address
		Email string `json:"email" validate:"required"`
		// the user's hashed password
		PassWord string `json:"password" validate:"required"`
		// the user's id
		ID string `json:"id"`
	} `json:"message"`
	// A unique number that defines authenticates the users transactions
	Token string `json:"token"`
}

// swagger:route POST /api/login session idOfLoginEndpoint
// Returns a session token
// responses:
// 200: loginResponse
// 400: loginBadRequest
// 404: loginIncorrectPassword && loginNotFound
// Login returns a token and user details from the user data
func (db *Sessiondb) Login(w http.ResponseWriter, req *http.Request) {

	var user LoginCredentials
	var userDb models.UserCredentials

	err := json.NewDecoder(req.Body).Decode(&user) //fills up user from body
	if err != nil {
		JsonError(&w, err.Error(), http.StatusBadRequest)
	}

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
	w.WriteHeader(http.StatusOK)

	res := LoginResponse{
		true,
		userDb,
		token,
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		JsonError(&w, err.Error(), http.StatusInternalServerError)
	}
}

func JsonError(w *http.ResponseWriter, ErrorMessage string, ErrorCode int) {
	(*w).WriteHeader(ErrorCode)
	res := MyStdResp{
		Status:  false,
		Message: ErrorMessage,
	}
	json.NewEncoder(*w).Encode(res)
}

/* Request body received
{
  "email":"natverior1@gmail.com",
  "password":"password"
}
*/

/* Response to sent
{
    "Status": true,
    "Message": {
        "id": "1",
        "firstname": "Nator",
        "lastname": "Verinumbe",
        "email": "natverior1@gmail.com",
        "password": "$2a$10$88kYQEH6sP2xPjGD3GeQ5e5hcq74yIHbl.Vo8SDdJgsmx28IdgPQu"
    },
    "token": "fa3af482-4685-11eb-8c2d-a01d486a6c86"
}
*/
