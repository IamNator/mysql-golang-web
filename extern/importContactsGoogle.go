package extern

//GET https://people.googleapis.com/v1/%5BRESOURCENAME%5D/connections?key=[YOUR_API_KEY] HTTP/1.1
//
//Authorization: Bearer [YOUR_ACCESS_TOKEN]
//Accept: application/json
import (
	"context"
	"fmt"
	"github.com/IamNator/mysql-golang-web/models"
	"net/http"
	"time"
)

type db models.DBData

type respBody struct {

	Connections []struct{
	  Person []struct{
	  	Name []struct{
			FamilyName string `json:"familyName"`
			GivenName string `json:"givenName"`
			MiddleName string `json:"middleName"`
		}
		PhoneNumber []struct{
			Value string `json:"value"`
		}
	  }
	}					 			`json:"connections"`
	NextPageToken string		    `json:"nextPageToken"`
	NextSyncToken string			`json:"nextSyncToken"`
	TotalPeople int64				`json:"totalPeople"`
	TotalItems int64				`json:"totalItems"`

}

func (db *db) getContact(){

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(5 * time.Second))
	defer cancelFunc()

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		"https://people.googleapis.com/v1/%5BRESOURCENAME%5D/connections?key=[YOUR_API_KEY] HTTP/1.1",
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

