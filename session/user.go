package session

import "database/sql"


//For registering new users
type Credentials struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	PassWord  string `json:"password"`
}

type LoginCredentials struct {
	Email string `json:"email"`
	PassWord string `json:"password"`
}

type DBData struct {
	DBType, User, Password, Host, DBName string
	Session                              *sql.DB
	SessionIDs                           map[string]string
	SessionUsers                         map[string]string
}
