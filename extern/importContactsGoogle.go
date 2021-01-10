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
		ResourceName string         `json:"resourceName"`
		Etag 		 string 		`json:"etag"`
	  	Names []struct{
			FamilyName string 		`json:"familyName"`
			GivenName string 		`json:"givenName"`
			MiddleName string 		`json:"middleName"`
		}							`json:"names"`
		PhoneNumbers []struct{
			Value string 			`json:"value"`
		}							`json:"phoneNumbers"`

	}					 			`json:"connections"`
	NextPageToken string		    `json:"nextPageToken"`
	NextSyncToken string			`json:"nextSyncToken"`
	TotalPeople int64				`json:"totalPeople"`
	TotalItems int64				`json:"totalItems"`

}

// [
//   {
//      "resourceName": "people/c5239008832403875235",
//      "etag": "%EgcBAj0LPzcuGgQBAgUHIgxUNjQrRDhoUmpHMD0=",
//      "names": [
//        {
//          "metadata": {
//            "primary": true,
//            "source": {
//              "type": "CONTACT",
//              "id": "48b4b2c68e9125a3"
//            }
//          },
//          "displayName": "Abibat Janey Laws Friend SelfCrush",
//          "familyName": "SelfCrush",
//          "givenName": "Abibat Janey Laws",
//          "middleName": "Friend",
//          "displayNameLastFirst": "SelfCrush, Abibat Janey Laws Friend",
//          "unstructuredName": "Abibat Janey Laws Friend SelfCrush"
//        }
//      ],
//      "phoneNumbers": [
//        {
//          "metadata": {
//            "primary": true,
//            "source": {
//              "type": "CONTACT",
//              "id": "48b4b2c68e9125a3"
//            }
//          },
//          "value": "+234 708 740 4908",
//          "canonicalForm": "+2347087404908",
//          "type": "mobile",
//          "formattedType": "Mobile"
//        }
//      ]
//   }
// ]

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
//  "connections": [
//    {
//      "resourceName": "people/c5588552881752772955",
//      "etag": "%EgcBAj0LPzcuGgQBAgUHIgxzY05YOERaSEllRT0=",
//      "phoneNumbers": [
//        {
//          "metadata": {
//            "primary": true,
//            "source": {
//              "type": "CONTACT",
//              "id": "4d8e873c8d5c115b"
//            }
//          },
//          "value": "+2348100642925",
//          "canonicalForm": "+2348100642925",
//          "type": "mobile",
//          "formattedType": "Mobile"
//        }
//      ]
//    },
//    {
//      "resourceName": "people/c4795508802158220279",
//      "etag": "%EgcBAj0LPzcuGgQBAgUHIgxNbDZuSEQvQSs2RT0=",
//      "phoneNumbers": [
//        {
//          "metadata": {
//            "primary": true,
//            "source": {
//              "type": "CONTACT",
//              "id": "428d11d90de653f7"
//            }
//          },
//          "value": "+2348132090452",
//          "canonicalForm": "+2348132090452",
//          "type": "mobile",
//          "formattedType": "Mobile"
//        }
//      ]
//    },
//    {
//      "resourceName": "people/c7842939431841336216",
//      "etag": "%EgcBAj0LPzcuGgQBAgUHIgxNbmVWUG1Mamt4ST0=",
//      "names": [
//        {
//          "metadata": {
//            "primary": true,
//            "source": {
//              "type": "CONTACT",
//              "id": "6cd7b7dc0dce8b98"
//            }
//          },
//          "displayName": "Mr. Abba",
//          "familyName": "Abba",
//          "honorificPrefix": "Mr.",
//          "displayNameLastFirst": "Abba, Mr.",
//          "unstructuredName": "Mr. Abba"
//        }
//      ],
//      "phoneNumbers": [
//        {
//          "metadata": {
//            "primary": true,
//            "source": {
//              "type": "CONTACT",
//              "id": "6cd7b7dc0dce8b98"
//            }
//          },
//          "value": "08150759317",
//          "canonicalForm": "+2348150759317",
//          "type": "mobile",
//          "formattedType": "Mobile"
//        }
//      ]
//    },
//    {
//      "resourceName": "people/c5239008832403875235",
//      "etag": "%EgcBAj0LPzcuGgQBAgUHIgxUNjQrRDhoUmpHMD0=",
//      "names": [
//        {
//          "metadata": {
//            "primary": true,
//            "source": {
//              "type": "CONTACT",
//              "id": "48b4b2c68e9125a3"
//            }
//          },
//          "displayName": "Abibat Janey Laws Friend SelfCrush",
//          "familyName": "SelfCrush",
//          "givenName": "Abibat Janey Laws",
//          "middleName": "Friend",
//          "displayNameLastFirst": "SelfCrush, Abibat Janey Laws Friend",
//          "unstructuredName": "Abibat Janey Laws Friend SelfCrush"
//        }
//      ],
//      "phoneNumbers": [
//        {
//          "metadata": {
//            "primary": true,
//            "source": {
//              "type": "CONTACT",
//              "id": "48b4b2c68e9125a3"
//            }
//          },
//          "value": "+234 708 740 4908",
//          "canonicalForm": "+2347087404908",
//          "type": "mobile",
//          "formattedType": "Mobile"
//        }
//      ]
//    },
//    {
//      "resourceName": "people/c6605848926967772260",
//      "etag": "%EgcBAj0LPzcuGgQBAgUHIgxuUHBWMkxoOWhURT0=",
//      "names": [
//        {
//          "metadata": {
//            "primary": true,
//            "source": {
//              "type": "CONTACT",
//              "id": "5bacb0998bcf3464"
//            }
//          },
//          "displayName": "Abiola Friend Indy Stalite First Squat",
//          "familyName": "Squat",
//          "givenName": "Abiola Friend Indy Stalite",
//          "middleName": "First",
//          "displayNameLastFirst": "Squat, Abiola Friend Indy Stalite First",
//          "unstructuredName": "Abiola Friend Indy Stalite First Squat"
//        }
//      ],
//      "phoneNumbers": [
//        {
//          "metadata": {
//            "primary": true,
//            "source": {
//              "type": "CONTACT",
//              "id": "5bacb0998bcf3464"
//            }
//          },
//          "value": "0705 966 5200",
//          "canonicalForm": "+2347059665200",
//          "type": "mobile",
//          "formattedType": "Mobile"
//        }
//      ]
//    },
//    {
//      "resourceName": "people/c4572580199681882533",
//      "etag": "%EgcBAj0LPzcuGgQBAgUHIgxJMTYxSFBMT21pOD0=",
//      "names": [
//        {
//          "metadata": {
//            "primary": true,
//            "source": {
//              "type": "CONTACT",
//              "id": "3f7511770d05f9a5"
//            }
//          },
//          "displayName": "Dr. Ade EEE Lecturer",
//          "familyName": "Lecturer",
//          "givenName": "Ade",
//          "middleName": "EEE",
//          "honorificPrefix": "Dr.",
//          "displayNameLastFirst": "Lecturer, Dr. Ade EEE",
//          "unstructuredName": "Dr. Ade EEE Lecturer "
//        }
//      ],
//      "phoneNumbers": [
//        {
//          "metadata": {
//            "primary": true,
//            "source": {
//              "type": "CONTACT",
//              "id": "3f7511770d05f9a5"
//            }
//          },
//          "value": "+234 805 801 5437",
//          "canonicalForm": "+2348058015437",
//          "type": "mobile",
//          "formattedType": "Mobile"
//        }
//      ]
//    },
//    {
//      "resourceName": "people/c5536759787340758547",
//      "etag": "%EgcBAj0LPzcuGgQBAgUHIgx0RTMwRTAzWGpjUT0=",
//      "names": [
//        {
//          "metadata": {
//            "primary": true,
//            "source": {
//              "type": "CONTACT",
//              "id": "4cd685b10b1cfe13"
//            }
//          },
//          "displayName": "Ade Micheals Friend Secondary school Ecwa",
//          "familyName": "Ecwa",
//          "givenName": "Ade Micheals Friend Secondary",
//          "middleName": "school",
//          "displayNameLastFirst": "Ecwa, Ade Micheals Friend Secondary school",
//          "unstructuredName": "Ade Micheals Friend Secondary school Ecwa"
//        }
//      ],
//      "phoneNumbers": [
//        {
//          "metadata": {
//            "primary": true,
//            "source": {
//              "type": "CONTACT",
//              "id": "4cd685b10b1cfe13"
//            }
//          },
//          "value": "+234 816 764 9977",
//          "canonicalForm": "+2348167649977",
//          "type": "mobile",
//          "formattedType": "Mobile"
//        }
//      ]
//    },
//    {
//      "resourceName": "people/c5066568094512725021",
//      "etag": "%EgcBAj0LPzcuGgQBAgUHIgwrM2hLRjFnaVJhTT0=",
//      "names": [
//        {
//          "metadata": {
//            "primary": true,
//            "source": {
//              "type": "CONTACT",
//              "id": "465010d68fc23c1d"
//            }
//          },
//          "displayName": "Ade Tailor Friend AGBOWO",
//          "familyName": "AGBOWO",
//          "givenName": "Ade Tailor",
//          "middleName": "Friend",
//          "displayNameLastFirst": "AGBOWO, Ade Tailor Friend",
//          "unstructuredName": "Ade Tailor Friend AGBOWO"
//        }
//      ],
//      "phoneNumbers": [
//        {
//          "metadata": {
//            "primary": true,
//            "source": {
//              "type": "CONTACT",
//              "id": "465010d68fc23c1d"
//            }
//          },
//          "value": "0704 574 1427",
//          "canonicalForm": "+2347045741427",
//          "type": "mobile",
//          "formattedType": "Mobile"
//        }
//      ]
//    }
//}

//addresses,ageRanges,biographies,birthdays,calendarUrls,clientData,coverPhotos,events,externalIds,genders,imClients,interests,locales,locations,memberships,metadata,miscKeywords,nicknames,occupations,organizations,photos,relations,sipAddresses,skills,urls,userDefined

//Good Morning Sir.
//	Here is my full residential address
//
//Thank you for reaching out, I await your response.