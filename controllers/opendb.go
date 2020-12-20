package controllers

import (
	"database/sql"
	"fmt"
	"time"
)

func (db DBData) OpenDB() (*sql.DB, error) {
	openDB, err := sql.Open(db.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s", db.User, db.Password, db.Host, db.DBName))
	//db.Session.SetMaxOpenConns(20)
	//db.Session.SetMaxIdleConns(20)
	db.Session.SetConnMaxLifetime(time.Minute * 5)
	return openDB, err
}
