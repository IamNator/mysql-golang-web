package seeders

import (
	//"database/sql"
	"encoding/json"
	//"log"
	"os"
)

//Reads Json file and fill mysql database
func (db Seeddb) FillDb() {

	file, err := os.OpenFile("data.json", os.O_CREATE, os.ModePerm)
	check(err)
	defer file.Close()

	var users []struct {
		FirstName   string `json:"firstname" validate:"required"`
		LastName    string `json:"lastname" validate:"required"`
		PhoneNumber string `json:"phone_number" validate:"required"`
		ID          string `json:"id"`
	}

	json.NewDecoder(file).Decode(&users)
	//fmt.Println(users)

	for _, values := range users {

		stmt, err := db.Session.Prepare(`INSERT INTO phoneBook (userID, FirstName,LastName,phoneNumber)
	VALUES (?,?,?,?)`)

		_, err = stmt.Exec(values.ID, values.FirstName, values.LastName, values.PhoneNumber)
		check(err)
	}
}
