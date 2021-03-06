package controllers

import (
	"encoding/json"
	"github.com/IamNator/jsonWriter"
	"github.com/IamNator/mysql-golang-web/models"
	"net/http"
)

// swagger:response fetchResponse
type fetchResponseWrapper struct {
	//in body
	Body struct {
		Status  bool                      `json:"status"`
		Message []models.PhoneBookContact `json:"message"`
	}
}

// token not valid, login to get a new one
// swagger:response fetchUnauthorized
type fetchUnauthorizedWrapper struct {
	// in: body
	Body MyStdResp
}

// unable to respond to request
// swagger:response fetchInternalError
type fetchInternalErrorWrapper struct {
	// in: body
	Body MyStdResp
}

// swagger:parameters fetch
type headerWrapper struct {
	// token
	// in: header
	Authorization string
}

// swagger:route GET /api/contacts controllers fetch
// returns all phonebook contacts
// responses:
// 200: fetchResponse
// 401: fetchUnauthorized
// 500: fetchInternalError
func (db *Controllersdb) Fetch(w http.ResponseWriter, req *http.Request) {

	token := req.Header.Get("Authorization")

	//id, ok := db.SessionToken[reqBody.Token]
	id, ok := db.SessionToken[token]

	if !ok { //Check if user is logged in (id exists in the MAP)
		jsonWriter.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}

	_ = db.Session.Ping()
	rows, err := db.Session.Query(`SELECT id, FirstName, LastName, phoneNumber FROM phoneBook WHERE userID=` + id.ID)
	Check(err)

	var user models.PhoneBookContact
	var users []models.PhoneBookContact

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.PhoneNumber)
		Check(err)

		users = append(users, user)
	}

	resp := struct {
		Status  bool        `json:"status"`
		Message interface{} `json:"message"`
	}{
		Status:  true,
		Message: users,
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		jsonWriter.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

/* Request Body received
{
  "token":"fa3af482-4685-11eb-8c2d-a01d486a6c86"
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
