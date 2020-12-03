package seeders

import (
	"database/sql"
	"encoding/json"
	"github.com/IamNator/mysql-golang-web/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
)

type DBData struct {
	DBType, User, Password, Host, DBName string
	Session                              *sql.DB
	SessionIDs							 map[string]string
	SessionUsers						 map[string]string
}

type user struct{
	userName string `json:"username"`
	passWord string `json:"password"`
}



func (db DBData) FillUSerDb() {


	file, err := os.OpenFile("user.json", os.O_CREATE, os.ModePerm)
	check(err)
	defer file.Close()


	var users []user

	json.NewDecoder(file).Decode(&users)
	for _, values := range users {
		v, _ := bcrypt.GenerateFromPassword([]byte(values.passWord), bcrypt.DefaultCost)
		values.passWord = string(v)
	}


	for _, values := range users {

		stmt, err := db.Session.Prepare(`INSERT INTO users ( username, password)
	VALUES (?,?)`)

		_, err = stmt.Exec( values.userName, values.passWord)
		check(err)
	}

}

func (db DBData) FillDb() {


	file, err := os.OpenFile("data.json", os.O_CREATE, os.ModePerm)
	check(err)
	defer file.Close()

	var users []models.User

	json.NewDecoder(file).Decode(&users)


		for _, values := range users {

			stmt, err := db.Session.Prepare(`INSERT INTO phoneBook (userID, FirstName,LastName,PhoneNumber)
	VALUES (?,?,?)`)

			_, err = stmt.Exec(values.ID, values.FirstName, values.LastName, values.PhoneNumber)
			check(err)
		}

}


func check(err error) {
	if err != nil {
		log.Printf("%+v\n", err)
	}
}