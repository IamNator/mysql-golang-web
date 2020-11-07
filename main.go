package main

import (
	"fmt"
	"github.com/IamNator/Projects/mysql-stuff/load_data"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./"))

	fmt.Println("server running...@localhost:9080")
	http.HandleFunc("/index", load_data.Index)
	http.Handle("/", fileServer)
	http.HandleFunc("/insert", load_data.Insert)
	http.HandleFunc("/form", load_data.FormHandler)
	http.HandleFunc("/del", load_data.Delete)
	err := http.ListenAndServe(":9080", nil)
	check(err)

}

// <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
// 		<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js" integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM" crossorigin="anonymous"></script>

func check(err error) {
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
