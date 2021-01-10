package extern


import (
	"context"
	"fmt"
	"github.com/IamNator/mysql-golang-web/models"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

type db models.DBData

//specifies how the response should be
var QP = QueryParameters{
	ResourceName: "people/me",
	PageSize: "1000",
	PersonFields: "names,phoneNumbers",
	RequestMaskIncludeField: SortOrder.FIRST_NAME_ASCENDING,
	Sources: Source.READ_SOURCE_TYPE_CONTACT,
}


//Returns URL for GET request for USERS contacts
func (qp *QueryParameters) url() (url string) {

	qp.SetApiKey()

	url, err := qp.SetURL()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	return url
}


//returns an array of contacts
func (db *db) getContact(){

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(5 * time.Second))
	defer cancelFunc()

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		QP.url(),
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

//extracts important information from app.env file { I should gitignore this file}
func LoadConfig(path string) (config QueryParameters, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)

	return
}


