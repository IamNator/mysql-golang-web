package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/IamNator/mysql-golang-web/models"
	"log"
	"net/http"
	"os"
	//"sync"
)

func (db *DBData) Fetch_t(w http.ResponseWriter, req *http.Request) {

	file, err := os.Open("data.json")
	check(err)
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

	file, err := os.Open("data.json")
	check(err)
	json.NewDecoder(file).Decode(&users)

	for i, values := range users {
		if values.ID == user.ID {
			file.Close()
			os.Remove("data.json")
			filee, _ := os.Open("data.json")
			users = append(users[i:], users[i+1:]...)
			json.NewEncoder(filee).Encode(&users)
			filee.Close()
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
		log.Println(err)
		return
	}

	file, err := os.OpenFile("data.json", os.O_CREATE, os.ModePerm)
	check(err)
	defer file.Close()

	var user models.User
	var users []models.User

	json.NewDecoder(req.Body).Decode(&user)
	json.NewDecoder(file).Decode(&users)

	{
		i := 0
		for _, values := range users {
			i++
			if values.ID != string(i) {
				user.ID = string(i)
				break
			}
		}
	}


	users = append(users, user)
	fmt.Println("About to enter data")
	if user.Fname != "" && user.Lname != "" && user.Phone_number != "" && user.ID != "" {
		file.Close()
		os.Remove("data.json")
		file, err := os.OpenFile("data.json", os.O_CREATE, os.ModePerm)
		check(err)
		
		json.NewEncoder(file).Encode(&users)
		fmt.Println("\nData Successfully Added")
		fmt.Fprintf(w, `Successful`)
	} else {
		fmt.Fprintf(w, `Please fill in all the fields `)
	}

}
