package user

import "database/sql"


//Soon to adopts this for database
type RegisterUser struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	PassWord  string `json:"password"`
}

type Credentials struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type DBData struct {
	DBType, User, Password, Host, DBName string
	Session                              *sql.DB
	SessionIDs                           map[string]string
	SessionUsers                         map[string]string
}
