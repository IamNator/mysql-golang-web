package session

import (
	"encoding/json"

	//"github.com/satori/go.uuid"
	"net/http"
	//"time"
)

//Takes in req.body = {token}
//
// Returns w.Body = { status:true, message:"logged out successfully" }
//
//userDetails = {id, firstname, lastname, email, password}
func (db *Sessiondb) Logout(w http.ResponseWriter, req *http.Request) {
	var user struct{
		Token string `json:"token"`
	}


	err := json.NewDecoder(req.Body).Decode(&user) //fills up user from body
	if err != nil {
		JsonError(&w, err.Error(), http.StatusBadRequest)
	}

	if user.Token == ""  {
		JsonError(&w, "Token not present", http.StatusBadRequest)
		return
	}

	{
		delete(db.SessionToken, user.Token)

		w.WriteHeader(http.StatusOK)

		res := struct {
			MyStdResp
			Token string `json:"token"`
		}{MyStdResp{true, "Logged out successfully" },
			user.Token,
		}

		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			JsonError(&w, err.Error(), http.StatusInternalServerError)
		}
	}

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