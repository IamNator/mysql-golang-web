package extern


import (
	"context"
	"fmt"
	"github.com/IamNator/mysql-golang-web/models"
	"net/http"
	"time"
)

type db models.DBData

var qp = QueryParameters{
	ApiKey: "AIzaSyAEuBpVzf3vDyaJ_tIwf_sLiIHDFOy8EGM",
	ResourceName: "people/me",
	PageSize: "1000",
	PersonFields: "names,phoneNumbers",
	RequestMaskIncludeField: SortOrder.FIRST_NAME_ASCENDING,
}

func (db *db) getContact(){

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(5 * time.Second))
	defer cancelFunc()

	url, err := qp.SetURL()
	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		url,
		nil,

		)

	client := http.Client{}
	req.Header.Set("Authorization", "Bearer [YOUR_ACCESS_TOKEN]")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(" Error : ", err )
		return
	}

	fmt.Println(resp.Body)
}


