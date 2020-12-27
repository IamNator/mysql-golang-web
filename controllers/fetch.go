package controllers

import (
	"encoding/json"
	"github.com/IamNator/mysql-golang-web/models"
	"github.com/IamNator/mysql-golang-web/session"
	"net/http"
)

// takes req.Body = { "token": "ere-dfd-f3432", "id": "42cv"}
//
//returns w.Body = { "status": "true", "message": [ phone book contacts ] }
func (db *Controllersdb) Fetch(w http.ResponseWriter, req *http.Request) {
	var reqBody struct {
		Token string `json:"token"`
		ID string    `json:"id"`
	}
	json.NewDecoder(req.Body).Decode(&reqBody)

	if _, ok := db.SessionToken[reqBody.Token]; !ok { //Check if user is logged in (id exists in the MAP)
		session.JsonError(&w, "Unauthorized access", http.StatusUnauthorized)
		return
	}

	_ = db.Session.Ping()
	rows, err := db.Session.Query(`SELECT id, FirstName, LastName, phoneNumber FROM phoneBook WHERE userID=` + db.SessionToken[reqBody.Token].ID)
	Check(err)

	var user models.PhoneBookContact
	var users []models.PhoneBookContact

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.PhoneNumber)
		Check(err)

		users = append(users, user)
	}


	resp := struct{
		Status bool `json:"status"`
		Message interface{} `json:"message"`
	}{
		Status: true,
		Message: users,
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		session.JsonError(&w, err.Error(), http.StatusInternalServerError)
	}

}
/* Request Body received
  {
    "token":"fa3af482-4685-11eb-8c2d-a01d486a6c86",
    "id":"1"
  }

*/

/* Response Received  200 OK
{
    "Status": true,
    "Message": [
        {
            "firstname": "Nator",
            "lastname": "Verinumber",
            "phone_number": "09045057268",
            "id": "1"
        },
        {
            "firstname": "Peter",
            "lastname": "John",
            "phone_number": "09045689434",
            "id": "11"
        },
        {
            "firstname": "Ajibola",
            "lastname": "Tantoloun",
            "phone_number": "09055655946",
            "id": "51"
        },
        {
            "firstname": "John",
            "lastname": "Pop",
            "phone_number": "09045057268",
            "id": "111"
        },
        {
            "firstname": "humble",
            "lastname": "jack",
            "phone_number": "09043453433",
            "id": "151"
        }
    ]
}
*/