package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/IamNator/mysql-golang-web/backend/models"
	"log"
	"net/http"
	"os"
	//"sync"
)

func (db *DBData) Fetch_t(w http.ResponseWriter, req *http.Request) {

	file, _ := os.Open("data.json")
	defer file.Close()

	//var user models.User
	var users []models.User

	json.NewDecoder(file).Decode(&users)

	json.NewEncoder(w).Encode(&users) //Sends an array of user information
	log.Println("Data fetched")
	//	db.Close()
}

func (db *DBData) Delete_t(writer http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm() err: %v", err)
		return
	}

	var user models.User
	var users []models.User
	json.NewDecoder(req.Body).Decode(&user)

	file, _ := os.Open("data.json")
	json.NewDecoder(file).Decode(&users)

	for _, values := range users {
		if values.ID == user.ID {
			file.Close()
			os.Remove("data.json")
			filee, _ := os.Open("data.json")
			users = append(users, user)
			json.NewEncoder(filee).Encode(&users)
			filee.Close()
			s := "{\"deleted\":\"successfully\"}"
			json.NewEncoder(writer).Encode(s)
			break
		}
	}

	writer.Header().Set("Content-Type", "application/json")
	s := "{\"deleted\":\"successfully\"}"
	json.NewEncoder(writer).Encode(s)

}

func (db *DBData) Update_t(w http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	file, _ := os.OpenFile("data.json", os.O_CREATE, os.ModePerm)
	defer file.Close()

	var user models.User
	var users []models.User

	json.NewDecoder(req.Body).Decode(&user)
	json.NewDecoder(file).Decode(&users)

	for _, values := range users {
		if values.ID == user.ID {
			fmt.Fprintf(w, `Dupliacate id detected`)
			return
		}
	}

	users = append(users, user)

	if user.Fname != "" && user.Lname != "" && user.Phone_number != "" && string(user.ID) != "" {
		file.Close()
		os.Remove("data.json")
		file, _ := os.OpenFile("data.json", os.O_CREATE, os.ModePerm)

		json.NewEncoder(file).Encode(&users)
		fmt.Println("\nData Successfully Added")
		fmt.Fprintf(w, `Successful`)
	} else {
		fmt.Fprintf(w, `Please fill in all the fields `)
	}

}
