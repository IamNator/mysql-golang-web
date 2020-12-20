package controllers_test

import (
	"github.com/IamNator/mysql-golang-web/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

//Test func (db *DBData) Fetch(w http.ResponseWriter, req *http.Request)
var db controllers.DBData

func TestDBData_Fetch(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/fetch", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(db.Fetch)
	handler.ServeHTTP(rr, req)
}

//Test func (db *DBData) Delete(writer http.ResponseWriter, req *http.Request)
//Test func (db *DBData) Update(w http.ResponseWriter, req *http.Request)
//
