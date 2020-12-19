package migrations

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type DBData struct {
	DBType, User, Password, Host, DBName string
	Session                              *sql.DB
	SessionIDs                           map[string]string
	SessionUsers                         map[string]string
}


func (db *DBData) CreateUserDb() {

	query := `CREATE TABLE IF NOT EXISTS users(id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, firstname NCHAR(50), lastname NCHAR(50), email NCHAR(50), password NCHAR(120), created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 15*time.Second)
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
