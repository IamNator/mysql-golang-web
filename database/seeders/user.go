package seeders

import (
	"database/sql"
	"encoding/json"
	//"github.com/IamNator/mysql-golang-web/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
)

type DBData struct {
	DBType, User, Password, Host, DBName string
	Session                              *sql.DB
	SessionIDs                           map[string]string
	SessionUsers                         map[string]string
}

type user struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}
type User struct {
	FirstName string `json:"firstname"`
	LastName string	 `json:"lastname"`
	Email string 	 `json:"email"`
	PassWord string  `json:"password"`
}


func (db DBData) FillUSerDb() {

	file, err := os.OpenFile("user.json", os.O_CREATE, os.ModePerm)
	check(err)
	defer file.Close()

	var users []user

	json.NewDecoder(file).Decode(&users)

	for index, values := range users {
		v, _ := bcrypt.GenerateFromPassword([]byte(values.PassWord), bcrypt.DefaultCost)
		users[index].PassWord = string(v)
	}

	for _, values := range users {

		stmt, err := db.Session.Prepare(`INSERT INTO users ( username, password)
	VALUES (?,?)`)

		_, err = stmt.Exec(values.UserName, values.PassWord)
		check(err)
	}

}

func (db DBData) FillUserDb_new() {

	file, err := os.OpenFile("user_new.json", os.O_CREATE, os.ModePerm)
	check(err)
	defer file.Close()

	var users []User

	json.NewDecoder(file).Decode(&users)

	for index, values := range users {
		v, _ := bcrypt.GenerateFromPassword([]byte(values.PassWord), bcrypt.DefaultCost)
		users[index].PassWord = string(v)
	}

	for _, values := range users {

		stmt, err := db.Session.Prepare(`INSERT INTO users ( firstname, lastname, email, password )
	VALUES (?,?,?,?)`)

		_, err = stmt.Exec(values.FirstName, values.LastName, values.Email, values.PassWord)
		check(err)
	}

}


func check(err error) {
	if err != nil {
		log.Printf("%+v\n", err)
	}
}
