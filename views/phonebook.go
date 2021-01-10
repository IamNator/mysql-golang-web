package views

import (
	//"golang.org/x/net/html"
	"html/template"
	"net/http"
)

func PhoneBook(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("html/index.html"))
	_ = tpl.Execute(w, nil)
}

