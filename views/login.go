package views

import (
	//"golang.org/x/net/html"
	"html/template"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, _ *http.Request) {
	tpl := template.Must(template.ParseFiles("web/login.html"))
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
