package google
//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"log"
//	"net/http"
//	"time"
//)
//
////url parameters
//var queryParams = QueryParameters{
//	ResourceName: "people/me",
//	PageSize: "2000",
//	PersonFields: "names,phoneNumbers",
//	RequestMaskIncludeField: SortOrder.FIRST_NAME_ASCENDING,
//	Sources: Source.READ_SOURCE_TYPE_CONTACT,
//}
//
//
////Returns URL string for GET request for USERS contacts
//func (qp *QueryParameters) url() (url string) {
//
//	qp.SetApiKey()
//
//	url, err := qp.SetURL()
//	if err != nil {
//		log.Fatal("cannot load config:", err)
//	}
//	return url
//}
//
//
////returns an array of contacts
//func GetContacts(response *RespBody) {
//
//
//	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(5 * time.Second))
//	defer cancelFunc()
//
//
//	url := queryParams.url()
//	fmt.Println("URL :", url)
//	req, err := http.NewRequestWithContext(
//				ctx,
//		"GET",
//		 		url,
//		   nil,
//		)
//
//	if req == nil {
//		fmt.Println(" Error : nil request object" )
//		return
//	}
//
//
//	client := http.Client{}
//	req.Header.Set("Authorization", "")
//	req.Header.Set("Accept", "application/json")
//
//
//	resp, err := client.Do(req)
//	if err != nil {
//		fmt.Println(" Error : ", err )
//		return
//	}
//
//	err = json.NewDecoder(resp.Body).Decode(&response)
//	if err != nil {
//		fmt.Println(" Error : ", err )
//		return
//	}
//
//}
//
//
