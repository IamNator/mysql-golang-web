package views

import (
	//"golang.org/x/net/html"
	"html/template"
	"net/http"
)

func Register(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("web/register.html"))
	tpl.Execute(w, nil)
}
