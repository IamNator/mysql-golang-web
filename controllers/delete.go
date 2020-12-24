package controllers

import (
	"encoding/json"
	//"fmt"
	"net/http"
)

//body takes ("token": "342-342s-fsd-343cv", "id": "23"}
//returns in body ("status": "true", "message": "deleted id, first and last name"
func (db *Controllersdb) Delete(writer http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
	//	http.Error(writer, fmt.Sprintf("ParseForm()  err : %v",err), http.StatusBadRequest )
	//	fmt.Fprintf(writer, "ParseForm() err: %v", err)
		return
	}
	var user struct {
		Token string `json:"token"`
		ID string `json:"id"` //id to be deleted
	}

	_ = json.NewDecoder(req.Body).Decode(&user)
	if _, ok := db.SessionToken[user.Token]; !ok {
		writer.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(writer).Encode("Cookie not found")
	}

	userID := db.SessionToken[user.Token]

	stmt, err := db.Session.Prepare(`DELETE FROM phoneBook WHERE id = ?, userID = ? ;`)
	_, err = stmt.Exec(user.ID, userID)
	Check(err)

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode("deleted")

}
