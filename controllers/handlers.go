package controllers

import (
	"database/sql"
	"fmt"
	//"golang.org/x/net/html"
	"encoding/json"
	"github.com/IamNator/mysql-golang-web/models"
	"log"
	"net/http"
)

type DBData struct {
	DBType, User, Password, Host, DBName string
	Session                              *sql.DB
	SessionIDs                           map[string]string
	SessionUsers                         map[string]string
}

func (db DBData) OpenDB() (*sql.DB, error) {
	opendb, err := sql.Open(db.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s", db.User, db.Password, db.Host, db.DBName))
	//db.Session.SetMaxOpenConns(20)
	//db.Session.SetMaxIdleConns(20)
	//db.Session.SetConnMaxLifetime(time.Minute * 5)
	return opendb, err
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

func (db *DBData) Fetch(w http.ResponseWriter, req *http.Request) {
	var SessionUserID string
	cookie, err := req.Cookie("sessionID")
	if err != nil {
		http.Redirect(w, req, "/", 301)
		fmt.Println("Cookie not found")
		return

	}
	userName := db.SessionIDs[cookie.Value]      //returns the username
	if id, ok := db.SessionUsers[userName]; ok { //Check if user is logged in (id exists in the MAP)
		SessionUserID = id
	} else {
		http.Redirect(w, req, "/", 301)
		fmt.Println("Cookie.Value does not match userNAme")
		return
	}

	db.Session.Ping()
	rows, err := db.Session.Query(`SELECT id, FirstName, LastName, PhoneNumber FROM phoneBook WHERE userID=` + SessionUserID)
	check(err)

	var user models.User
	var users []models.User

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.PhoneNumber)
		check(err)

		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users) //Sends an array of user information

}

func (db *DBData) Delete(writer http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm() err: %v", err)
		return
	}
	var user struct {
		ID string `json:"id"`
	}
	json.NewDecoder(req.Body).Decode(&user)
	ck, _ := req.Cookie("sessionID")
	username := db.SessionIDs[ck.Value]
	userID := db.SessionUsers[username]

	stmt, err := db.Session.Prepare(`DELETE FROM phoneBook WHERE id = ?, userID = ? ;`)
	_, err = stmt.Exec(user.ID, userID)
	check(err)

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode("deleted")

}

func (db *DBData) Update(w http.ResponseWriter, req *http.Request) {

	var user models.User
	json.NewDecoder(req.Body).Decode(&user)

	if user.FirstName != "" && user.LastName != "" && user.PhoneNumber != "" && string(user.ID) != "" {
		//code needs optimization
		var userid string
		ck, _ := req.Cookie("sessionID")
		userid = db.SessionUsers[ck.Value]
		stmt, err := db.Session.Prepare(`INSERT INTO phoneBook (userID, FirstName,LastName,phoneNumber)
	VALUES (?,?,?,?)`)

		_, err = stmt.Exec(userid, user.FirstName, user.LastName, user.PhoneNumber)
		check(err)

		if err != nil {
			fmt.Fprintln(w, err)
		} else {
			fmt.Println("Data Successfully Added")
		}

		fmt.Fprintf(w, `Successful\n`)
	} else {
		fmt.Fprintf(w, `Please fill in all the fields\n`)
	}

}

func check(err error) {
	if err != nil {
		log.Printf("%+v\n", err)
	}
}
