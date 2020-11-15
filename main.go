package main

import (
	"fmt"
	"github.com/IamNator/mysql-golang-web/controllers"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./"))

	fmt.Println("server running...@localhost:9080")
	http.HandleFunc("/index", controllers.Index)
	http.Handle("/css/bootstrap.min.css", fileServer)
	http.Handle("/js/bootstrap.min.js", fileServer)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/form", controllers.FormHandler)
	http.HandleFunc("/del", controllers.Delete)
	err := http.ListenAndServe(":9080", nil)
	check(err)

}

func check(err error) {
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
