package seeders

import (
	//"database/sql"
	"encoding/json"
	"github.com/IamNator/mysql-golang-web/models"

	//"log"
	"os"
)


func (db DBData) FillDb() {

	file, err := os.OpenFile("data.json", os.O_CREATE, os.ModePerm)
	check(err)
	defer file.Close()

	var users []models.User

	json.NewDecoder(file).Decode(&users)
	//fmt.Println(users)

	for _, values := range users {

		stmt, err := db.Session.Prepare(`INSERT INTO phoneBook (userID, FirstName,LastName,phoneNumber)
	VALUES (?,?,?,?)`)

		_, err = stmt.Exec(values.ID, values.FirstName, values.LastName, values.PhoneNumber)
		check(err)
	}

}
