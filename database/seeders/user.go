package seeders

import (
	"database/sql"
	"encoding/json"
	"github.com/IamNator/mysql-golang-web/models"
	"log"
	"os"
)

type DBData struct {
	DBType, User, Password, Host, DBName string
	Session                              *sql.DB
}

func (db DBData) FillDb() {


	file, err := os.OpenFile("data.json", os.O_CREATE, os.ModePerm)
	check(err)
	defer file.Close()

	var users []models.User

	json.NewDecoder(file).Decode(&users)


		for _, values := range users {

			stmt, err := db.Session.Prepare(`INSERT INTO phoneBook (FirstName,LastName,PhoneNumber)
	VALUES (?,?,?)`)

			_, err = stmt.Exec(values.FirstName, values.LastName, values.PhoneNumber)
			check(err)
		}

}


func check(err error) {
	if err != nil {
		log.Printf("%+v\n", err)
	}
}