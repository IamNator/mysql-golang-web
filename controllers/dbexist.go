package controllers

func (db *Controllersdb) DbExists() bool {
	var id int
	idn := 1
	err := db.Session.QueryRow("Select id From phoneBook WHERE id=?", idn).Scan(&id)
	if err != nil {
		//	fmt.Println(err)
		return false
	} else {
		//	fmt.Println(err)
		return true
	}
}

