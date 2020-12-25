package session

import (
	"encoding/json"
	"net/http"
)

//Takes in req.body = {token}
//
// Returns w.Body = { status:true, message:"logged out successfully" }
//
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
		delete(db.SessionToken, user.Token) //token is deleted here

		w.WriteHeader(http.StatusOK)

		res := MyStdResp {
			true,
			"Logged out successfully",
		}


		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			JsonError(&w, err.Error(), http.StatusInternalServerError)
		}
	}

}


/* Request body received
{
  "token":"fa3af482-4685-11eb-8c2d-a01d486a6c86"
}
*/


/* Response to sent
{
    "Status": true,
    "Message": "Logged out successfully"
}
*/