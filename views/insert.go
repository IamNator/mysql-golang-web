package views

import (
	//"golang.org/x/net/html"
	"html/template"
	"net/http"
)

func Insert(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("html/insert.html"))
	tpl.Execute(w, nil)
}
