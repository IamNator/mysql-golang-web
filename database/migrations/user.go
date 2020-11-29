package migrations

import (
	"context"
	"database/sql"
	"log"
	"time"
)

//
//import "database/sql"


type DBData struct {
	DBType, User, Password, Host, DBName string
	Session                              *sql.DB
}

func (db *DBData) CreateUserDb() {

	query := `CREATE TABLE IF NOT EXISTS users(id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, username VARCHAR(50),  
        password VARCHAR(120), created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	res, err := db.Session.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return
	}
	log.Printf("Rows affected when creating table: %d", rows)

}


func (db *DBData) CreatePhoneBookDb() {

	query := `CREATE TABLE IF NOT EXISTS phoneBook(id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, fname VARCHAR(50),  
        lname VARCHAR(50), phone_number  VARCHAR(16), created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)`

	//SELECT fname, lname, phone_number, id FROM phoneBook`)
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	res, err := db.Session.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return
	}
	log.Printf("Rows affected when creating table: %d", rows)

}
//
//
//type User struct {
//	Fname        string `json:"fname"`
//	Lname        string `json:"lname"`
//	Phone_number string `json:"phone_number"`
//	ID           string `json:"id"`
//}


//CREATE TABLE users(
//id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
//username VARCHAR(50),
//password VARCHAR(120)
//);