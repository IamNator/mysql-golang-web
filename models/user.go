package models

type User struct {
	FirstName        string `json:"fname"`
	LastName        string `json:"lname"`
	PhoneNumber string `json:"phone_number"`
	ID           string `json:"id"`
}
