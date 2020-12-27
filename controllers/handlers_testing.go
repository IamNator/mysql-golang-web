package controllers

//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/IamNator/mysql-golang-web/models"
//	"log"
//	"net/http"
//	"os"
//	"strconv"
//	//"sync"
//)
//
//func (db *DBData) Fetch_t(w http.ResponseWriter, _ *http.Request) {
//
//	file, err := os.OpenFile("data.json", os.O_CREATE, os.ModePerm)
//	Check(err)
//
//	//var user models.User
//	var users []models.User
//
//	json.NewDecoder(file).Decode(&users)
//	file.Close() //Closes file after it's read
//
//	json.NewEncoder(w).Encode(&users) //Sends an array of user information
//	log.Println("Data fetched")
//}
//
//func (db *DBData) Delete_t(writer http.ResponseWriter, req *http.Request) {
//	if err := req.ParseForm(); err != nil {
//		fmt.Fprintf(writer, "ParseForm() err: %v", err)
//		return
//	}
//
//	var user models.User
//	var users []models.User
//	json.NewDecoder(req.Body).Decode(&user)
//
//	file, err := os.Open("data.json")
//	Check(err)
//	json.NewDecoder(file).Decode(&users)
//	file.Close()
//
//	for i, values := range users {
//		if values.ID == user.ID {
//			os.Remove("data.json")
//			file, err := os.OpenFile("data.json", os.O_CREATE, os.ModePerm)
//			Check(err)
//			fmt.Println("About to delete")
//			users = append(users[:i], users[i+1:]...)
//			jd := json.NewEncoder(file)
//			jd.SetIndent("", "   ")
//			jd.Encode(&users)
//			file.Close()
//			break
//		}
//	}
//
//	writer.Header().Set("Content-Type", "application/json")
//	s := "deleted"
//	json.NewEncoder(writer).Encode(s)
//
//}
//
//func (db *DBData) Update_t(w http.ResponseWriter, req *http.Request) {
//
//	if err := req.ParseForm(); err != nil {
//		fmt.Fprintf(w, "ParseForm() err: %v", err)
//		log.Println(err)
//		return
//	}
//
//	file, err := os.OpenFile("data.json", os.O_CREATE, os.ModePerm)
//	Check(err)
//	defer file.Close()
//
//	var user models.User
//	var users []models.User
//
//	json.NewDecoder(req.Body).Decode(&user)
//	json.NewDecoder(file).Decode(&users)
//
//	{
//		i := 0
//		for _, values := range users {
//			if values.ID != strconv.Itoa(i) {
//				user.ID = strconv.Itoa(i)
//				break
//			}
//			i++
//		}
//	}
//
//	users = append(users, user)
//	fmt.Println("About to enter data")
//	if user.FirstName != "" && user.LastName != "" && user.PhoneNumber != "" && user.ID != "" {
//		file.Close()
//		os.Remove("data.json")
//		file, err := os.OpenFile("data.json", os.O_CREATE, os.ModePerm)
//		Check(err)
//		jd := json.NewEncoder(file)
//		jd.SetIndent("", "   ")
//		jd.Encode(&users)
//		fmt.Println("\nData Successfully Added")
//		fmt.Fprintf(w, `Successful`)
//	} else {
//		fmt.Fprintf(w, `Please fill in all the fields `)
//	}
//
//}
