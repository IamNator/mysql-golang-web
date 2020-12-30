package session

import (
	"encoding/json"
	"net/http"
)


// swagger: parameters idOfLogoutEndpoint
type logoutRequestWrapper struct {
	// in: query
	Body struct {
		Token string `json:"token"`
	}
}


// swagger:response logoutResponse
type logoutResponseWrapper struct {
	// in: body
	Body MyStdResp
}

// unable to process request
// swagger:response logoutBadRequest
type logoutBadRequestWrapper struct {
	// in: body
	Body MyStdResp
}

// Unable to respond
// swagger:response logoutInternalError
type logoutInternalErrorWrapper struct {
	// in: body
	Body MyStdResp
}


// swagger:route POST /api/logout session idOfLogoutEndpoint
// logs the user out
// responses:
// 200: logoutResponse
// 400: logoutBadRequest
// 500: logoutInternalError
func (db *Sessiondb) Logout(w http.ResponseWriter, req *http.Request) {
	var user struct {
		Token string `json:"token"`
	}

	err := json.NewDecoder(req.Body).Decode(&user) //fills up user from body
	if err != nil {
		JsonError(&w, err.Error(), http.StatusBadRequest)
	}

	if user.Token == "" {
		JsonError(&w, "Token not present", http.StatusBadRequest)
		return
	}

	{
		Mutex.Lock()
		delete(db.SessionToken, user.Token) //token is deleted here
		Mutex.Unlock()

		w.WriteHeader(http.StatusOK)

		res := MyStdResp{
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
