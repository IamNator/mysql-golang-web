package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (db *DBData) Delete(writer http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm() err: %v", err)
		return
	}
	var user struct {
		ID string `json:"id"`
	}
	_ = json.NewDecoder(req.Body).Decode(&user)
	ck, er := req.Cookie("sessionID")
	if er == http.ErrNoCookie {
		writer.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(writer).Encode("Cookie not found")
	}

	username := db.SessionIDs[ck.Value]
	userID := db.SessionUsers[username]

	stmt, err := db.Session.Prepare(`DELETE FROM phoneBook WHERE id = ?, userID = ? ;`)
	_, err = stmt.Exec(user.ID, userID)
	Check(err)

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode("deleted")

}
