package migrations

import (
	"context"
	"log"
	"time"
)

func (db *Migrationdb) CreatePhoneBookDb() {

	query := `CREATE TABLE IF NOT EXISTS phoneBook(id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, userID INT NOT NULL, FirstName NCHAR(50),  
        LastName NCHAR(50), phoneNumber  VARCHAR(16), created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 10*time.Second)
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
