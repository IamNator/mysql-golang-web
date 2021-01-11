package views

import (
	//"golang.org/x/net/html"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	tpl := template.Must(template.ParseFiles("web/home.html"))
	_ = tpl.Execute(w, nil)
}
