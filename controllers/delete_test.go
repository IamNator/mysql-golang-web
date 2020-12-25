package controllers_test
//
//import (
//	"github.com/IamNator/mysql-golang-web/controllers"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//var dbs controllers.DBData
//
////Still need a ton of improvement
//func TestDBData_Delete(t *testing.T) {
//	req, err := http.NewRequest("POST", "/api/delete", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//	rr := httptest.NewRecorder()
//
//	handler := http.HandlerFunc(dbs.Delete)
//	handler.ServeHTTP(rr, req)
//
//	if rr.Code != http.StatusOK {
//		t.Errorf("handler returned unexpected status code : got %v wanted %v", rr.Code, http.StatusOK)
//	}
//
//	expected := "deleted"
//	if rr.Body.String() != expected {
//		t.Errorf("handler returned unexpected body : got %v wanted %v", rr.Body.String(), expected)
//	}
//
//}
