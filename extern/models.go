package extern

import (
	"net/url"
	"fmt"
)

type Metadata struct{
	Primary bool	    `json:"primary"` // "true" if verified and "false" if otherwise
	Source struct{
		Type string		`json:"type"`    //"CONTACT"
		ID string 		`json:"id"`      //"48b4b2c68e9125a3"
	}					`json:"source"`
}

type RespBody struct {

	Connections []struct{
		ResourceName string         `json:"resourceName"` // e.g "people/c5239008832403875235"
		Etag 		 string 		`json:"etag"`   // e.g "%EgcBAj0LPzcuGgQBAgUHIgxUNjQrRDhoUmpHMD0="
		Names []struct{
			Metadata                `json:"metadata"`
			DisplayName string      `json:"displayName"`
			FamilyName string 		`json:"familyName"`
			GivenName string 		`json:"givenName"`
			MiddleName string 		`json:"middleName"`
		}							`json:"names"`
		PhoneNumbers []struct{
			Metadata                `json:"metadata"`
			CanonicalForm string 	`json:"canonicalForm"`  // e.g "+2347087404908"
			Type string 			`json:"type"`			// e.g "mobile"
		}							`json:"phoneNumbers"`

	}					 			`json:"connections"`
	NextPageToken string		    `json:"nextPageToken"`
	NextSyncToken string			`json:"nextSyncToken"`
	TotalPeople int64				`json:"totalPeople"`
	TotalItems int64				`json:"totalItems"`

}



type sortOrder string
const (
	FIRST_NAME_ASCENDING = sortOrder("FIRST_NAME_ASCENDING")
	LAST_NAME_ASCENDING = sortOrder("LAST_NAME_ASCENDING")
	LAST_MODIFIED_ASCENDING = sortOrder("LAST_MODIFIED_ASCENDING")
	LAST_MODIFIED_DESCENDING = sortOrder("LAST_MODIFIED_DESCENDING")
)

var SortOrder = [4]sortOrder{ FIRST_NAME_ASCENDING, LAST_NAME_ASCENDING, LAST_MODIFIED_ASCENDING, LAST_MODIFIED_DESCENDING,
}




type QueryParameters struct {
	ApiKey string           			`json:"apiKey" validate:"required"`
	ResourceName string     			`json:"resourceName"    validate:"required"`
	PageToken string       			    `json:"pageToken"`
	PageSize string 					`json:"pageSize"        validate:"required"`
	PersonFields string  			    `json:"personFields"    validate:"required"`
	RequestMaskIncludeField string		`json:"requestMask.includeField"`
	RequestSyncToken string 			`json:"requestSyncToken"`
	SortOrder sortOrder 				`json:"sortOrder"        validate:"required"`
	Sources string 						`json:"sources"`
	SyncToken string					`json:"syncToken"`
}



func (qp QueryParameters) setURL() (*url.URL, error) {
	url, err := url.Parse("https://people.googleapis.com/v1/%5BcontactGroups/all%5D/connections?key=[YOUR_API_KEY]")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	//url.Scheme = "https"
	//url.Host = "people.googleapis.com"
	q := url.Query()
	q.Set("key", qp.ApiKey)
	q.Set("resourceName", qp.ResourceName)
	q.Set("pageToken", qp.PageToken)
	q.Set("pageSize", qp.PageSize)
	q.Set("personFields", qp.PersonFields)
	q.Set("requestMask.includeField", qp.RequestMaskIncludeField)
	q.Set("requestSyncToken", qp.RequestSyncToken)
	q.Set("sortOrder", string(qp.SortOrder) )
	q.Set("sources", qp.Sources)
	q.Set("syncToken", qp.SyncToken)

	return url, nil
}
