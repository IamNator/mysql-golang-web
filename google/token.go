package google
//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"log"
//	"net/http"
//	"net/url"
//	"time"
//)
//
////https://www.googleapis.com/auth/contacts.readonly
//
////#define GOOGLE_AUTH_URL "https://accounts.google.com/o/oauth2/device/code"
////#define GOOGLE_AUTH_POST "client_id="GOOGLE_AUTH_CLIENT_ID"&scope=email profile https://www.googleapis.com/auth/contacts.readonly"
////
////int main(void) {
////    char * res = handle_url(GOOGLE_AUTH_URL,GOOGLE_AUTH_POST); // use curl to make a POST
////
////    if (res==NULL) {
////        Report("Error");
////        return -1;
////    }
////    cJSON *obj = cJSON_Parse(res);
////    printf("Result=%s\n",cJSON_Print(obj));
////    return 0;
////}
//
//func setUrl() string {
//	u_rl, _ := url.ParseRequestURI("https://www.googleapis.com/auth/contacts.readonly")
//
//
//	return u_rl.String()
//}
//
//
//func GetToken(){
//	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(5 * time.Second))
//	defer cancelFunc()
//
//	req, err := http.NewRequestWithContext(
//		ctx,
//		"POST",
//		   "https://accounts.google.com/o/oauth2/device/code",
//		   nil,
//	)
//
//	if req == nil {
//		fmt.Println(" Error : nil request object" )
//		return
//	}
//
//	client := http.Client{}
//	req.Header.Set("Authorization", "Bearer [YOUR_ACCESS_TOKEN]")
//	req.Header.Set("Accept", "application/json")
//
//
//	resp, err := client.Do(req)
//	if err != nil {
//		fmt.Println(" Error : ", err )
//		return
//	}
//
//	var response struct{} //response for token
//	err = json.NewDecoder(resp.Body).Decode(&response)
//	defer log.Println(resp.Body.Close())
//	if err != nil {
//		fmt.Println(" Error : ", err )
//		return
//	}
//
//}
//

