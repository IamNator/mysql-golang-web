package views

import (
	//"golang.org/x/net/html"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("html/login.html"))
	tpl.Execute(w, nil)
}
