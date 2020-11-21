package models

type User struct {
	Fname        string `json:"fname"`
	Lname        string `json:"lname"`
	Phone_number string `json:"phone_number"`
	ID           int    `json:"id"`
}
