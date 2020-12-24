package controllers

import "fmt"

func (db *Controllersdb) CloseDB() string {
	err := db.Session.Close()
	if err != nil {
		return fmt.Sprintf("%v", err)
	} else {
		return fmt.Sprintln("Data base closed")
	}
}
