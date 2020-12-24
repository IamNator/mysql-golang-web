package seeders

import (
	"encoding/json"
	"github.com/IamNator/mysql-golang-web/models"

	//"github.com/IamNator/mysql-golang-web/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
)

////Keeps track of logins and mysql login data
//type DBData struct {
//	DBType, User, Password, Host, DBName string
//	Session                              *sql.DB
//	SessionToken                         map[string]models.User
//}
type Seeddb models.DBData

//Old model, we will change this soon
type user struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

//Data collected and stored of user
type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	PassWord  string `json:"password"`
}

//
//Now implemented
func (db Seeddb) FillUserDb() {

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

//Checks for errors
func check(err error) {
	if err != nil {
		log.Printf("%+v\n", err)
	}
}
