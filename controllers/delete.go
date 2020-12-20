package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type DBData struct {
	DBType, User, Password, Host, DBName string
	Session                              *sql.DB
	SessionIDs                           map[string]string
	SessionUsers                         map[string]string
}

func (db DBData) OpenDB() (*sql.DB, error) {
	openDB, err := sql.Open(db.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s", db.User, db.Password, db.Host, db.DBName))
	//db.Session.SetMaxOpenConns(20)
	//db.Session.SetMaxIdleConns(20)
	db.Session.SetConnMaxLifetime(time.Minute * 5)
	return openDB, err
}

func (db DBData) CloseDB() string {
	err := db.Session.Close()
	if err != nil {
		return fmt.Sprintf("%v", err)
	} else {
		return fmt.Sprintln("Data base closed")
	}
}

func (db *DBData) DbExists() bool {
	var id int
	idn := 1
	err := db.Session.QueryRow("Select id From phoneBook WHERE id=?", idn).Scan(&id)
	if err != nil {
		//	fmt.Println(err)
		return false
	} else {
		//	fmt.Println(err)
		return true
	}
}


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

func Check(err error) {
	if err != nil {
		log.Printf("%+v\n", err)
	}
}
