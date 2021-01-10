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
	"net/url"
	"time"
)

type db models.DBData

type respBody struct {

	Connections []struct{
	  Person struct{
	  	Name []struct{
			FamilyName string 		`json:"familyName"`
			GivenName string 		`json:"givenName"`
			MiddleName string 		`json:"middleName"`
		}							`json:"names"`
		PhoneNumber []struct{
			Value string 			`json:"value"`
		}							`json:"phoneNumbers"`
	  }                             `json:"person"`
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

func setURL( apiKey, resourceName, pageToken, pageSize, personFields, requestMask_includeField, requestSyncToken, sortOrder, sources, syncToken string) (*url.URL, error) {
	url, err := url.Parse("https://people.googleapis.com/v1/%5BcontactGroups/all%5D/connections?key=[YOUR_API_KEY]")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//url.Scheme = "https"
	//url.Host = "people.googleapis.com"
	q := url.Query()
	q.Set("key", apiKey)
	q.Set("resourceName", resourceName)
	q.Set("pageToken", pageToken)
	q.Set("pageSize", pageSize)
	q.Set("personFields", personFields)
	q.Set("requestMask.includeField", requestMask_includeField)
	q.Set("requestSyncToken", requestSyncToken)
	q.Set("sortOrder", sortOrder)
	q.Set("sources", sources)
	q.Set("syncToken", syncToken)

	return url, nil
}


// Api key - AIzaSyAEuBpVzf3vDyaJ_tIwf_sLiIHDFOy8EGM
//
//OAuth2.0
//Client ID - 745773636757-afhodihnn83bva8rlhssapo98onb7qvt.apps.googleusercontent.com
//Client Secret - zhMJZ2SWYKq-PM6zintdwYmo


//{
//  "contactGroups": [
//    {
//      "resourceName": "contactGroups/65189743883aabaf",
//      "etag": "UPTmggxWH8g=",
//      "metadata": {
//        "updateTime": "2020-03-25T12:25:13.787Z"
//      },
//      "groupType": "USER_CONTACT_GROUP",
//      "name": "ICE",
//      "formattedName": "ICE",
//      "memberCount": 1
//    },
//    {
//      "resourceName": "contactGroups/844909808499f0a",
//      "etag": "c8EKNg24QFk=",
//      "metadata": {
//        "updateTime": "2020-03-03T22:06:17.806Z"
//      },
//      "groupType": "USER_CONTACT_GROUP",
//      "name": "Bae",
//      "formattedName": "Bae"
//    },
//    {
//      "resourceName": "contactGroups/chatBuddies",
//      "groupType": "SYSTEM_CONTACT_GROUP",
//      "name": "chatBuddies",
//      "formattedName": "Chat contacts"
//    },
//    {
//      "resourceName": "contactGroups/all",
//      "groupType": "SYSTEM_CONTACT_GROUP",
//      "name": "all",
//      "formattedName": "All Contacts",
//      "memberCount": 595
//    },
//    {
//      "resourceName": "contactGroups/myContacts",
//      "groupType": "SYSTEM_CONTACT_GROUP",
//      "name": "myContacts",
//      "formattedName": "My Contacts",
//      "memberCount": 594
//    },
//    {
//      "resourceName": "contactGroups/friends",
//      "groupType": "SYSTEM_CONTACT_GROUP",
//      "name": "friends",
//      "formattedName": "Friends"
//    },
//    {
//      "resourceName": "contactGroups/family",
//      "groupType": "SYSTEM_CONTACT_GROUP",
//      "name": "family",
//      "formattedName": "Family"
//    },
//    {
//      "resourceName": "contactGroups/coworkers",
//      "groupType": "SYSTEM_CONTACT_GROUP",
//      "name": "coworkers",
//      "formattedName": "Coworkers"
//    },
//    {
//      "resourceName": "contactGroups/blocked",
//      "groupType": "SYSTEM_CONTACT_GROUP",
//      "name": "blocked",
//      "formattedName": "Blocked"
//    },
//    {
//      "resourceName": "contactGroups/starred",
//      "groupType": "SYSTEM_CONTACT_GROUP",
//      "name": "starred",
//      "formattedName": "Starred"
//    }
//  ],
//  "totalItems": 10,
//  "nextSyncToken": "EJ6I9Pzwj-4C"
//}

//addresses,ageRanges,biographies,birthdays,calendarUrls,clientData,coverPhotos,events,externalIds,genders,imClients,interests,locales,locations,memberships,metadata,miscKeywords,nicknames,occupations,organizations,photos,relations,sipAddresses,skills,urls,userDefined

//Good Morning Sir.
//	Here is my full residential address
//
//Thank you for reaching out, I await your response.