package google_test

import (
	"github.com/IamNator/mysql-golang-web/google"
	"testing"
)

func TestGetContacts(t *testing.T){


	var resp google.RespBody
	google.GetContacts(&resp)
	//fmt.Println(resp)
	i:=1
	if i == 1 {
		t.Log(resp)
		t.Error(resp)
	}

}