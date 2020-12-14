package views

import (
	//"golang.org/x/net/html"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("html/home.html"))
	tpl.Execute(w, nil)
}
