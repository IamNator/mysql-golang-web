package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/IamNator/mysql-golang-web/models"
	"github.com/IamNator/mysql-golang-web/session"
	"net/http"
	validate "github.com/go-playground/validator/v10"
)

func (db *DBData) Update(w http.ResponseWriter, req *http.Request) {

	var user models.User
	json.NewDecoder(req.Body).Decode(&user)
	validator := validate.New()
	err := validator.Struct(user)
	if err != nil {

	//}
	//if user.FirstName != "" && user.LastName != "" && user.PhoneNumber != "" && string(user.ID) != "" {
		//code needs optimizations
		var userid string
		ck, _ := req.Cookie("sessionID")
		userid = db.SessionToken[ck.Value]
		stmt, err := db.Session.Prepare(`INSERT INTO phoneBook (userID, FirstName,LastName,phoneNumber)
	VALUES (?,?,?,?)`)

		_, err = stmt.Exec(userid, user.FirstName, user.LastName, user.PhoneNumber)
		Check(err)

		if err != nil {
			session.JsonError(&w, "Unable to create user Database Error", http.StatusInternalServerError)
		} else {
			fmt.Println("Data Successfully Added")
		}
		fmt.Fprintf(w, `Successful\n`)
	} else {
		fmt.Fprintf(w, `Please fill in all the fields\n`)
	}

}
