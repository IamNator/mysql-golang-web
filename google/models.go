package google

import (
	"github.com/spf13/viper"
	"log"
	"net/url"
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


type sortOrfder struct {
	FIRST_NAME_ASCENDING,
	LAST_NAME_ASCENDING,
	LAST_MODIFIED_ASCENDING,
	LAST_MODIFIED_DESCENDING string
}

var SortOrder = sortOrfder{ "FIRST_NAME_ASCENDING", "LAST_NAME_ASCENDING", "LAST_MODIFIED_ASCENDING", "LAST_MODIFIED_DESCENDING",
}


type source struct {
	READ_SOURCE_TYPE_UNSPECIFIED,
	READ_SOURCE_TYPE_PROFILE,
	READ_SOURCE_TYPE_CONTACT,
	READ_SOURCE_TYPE_DOMAIN_CONTACT string
}

var Source = source{ "READ_SOURCE_TYPE_UNSPECIFIED", "READ_SOURCE_TYPE_PROFILE", "READ_SOURCE_TYPE_CONTACT", "READ_SOURCE_TYPE_DOMAIN_CONTACT",
}




type QueryParameters struct {
	ApiKey string           			`json:"apiKey" validate:"required" viper:"API_KEY"`
	ResourceName string     			`json:"resourceName"    validate:"required"`
	PageToken string       			    `json:"pageToken"`
	PageSize string 					`json:"pageSize"        validate:"required"`
	PersonFields string  			    `json:"personFields"    validate:"required"`
	RequestMaskIncludeField string		`json:"requestMask.includeField"`
	RequestSyncToken string 			`json:"requestSyncToken"`
	SortOrder string  	 				`json:"sortOrder"        validate:"required"`
	Sources string 						`json:"sources"`
	SyncToken string					`json:"syncToken"`
	PrettyPrint string                  `json:"prettyPrint"`
}


func (qp *QueryParameters) SetApiKey(){
	config, err := loadConfig("C:/gocode/src/github.com/IamNator/mysql-golang-web/" )
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	qp.ApiKey = config.ApiKey
}

//extracts important information from app.env file { I should gitignore this file}
func loadConfig(path string) (config QueryParameters, err error) {

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




func (qp *QueryParameters) SetURL() (string, error) {
	 var urL = url.URL{
			 Scheme: "https",
			 Host: "people.googleapis.com",
			 Path: "v1/people/me/connections",
	}

	//https://people.googleapis.com/v1/{resourceName=people/*}/connections

	q := urL.Query()
	q.Add("key", qp.ApiKey)
	q.Set("resourceName", qp.ResourceName)
	q.Add("pageToken", qp.PageToken)
	q.Add("pageSize", qp.PageSize)
	q.Add("personFields", qp.PersonFields)
	q.Add("requestMask.includeField", qp.RequestMaskIncludeField)
	q.Add("requestSyncToken", qp.RequestSyncToken)
	q.Add("sortOrder", string(qp.SortOrder) )
	q.Add("sources", qp.Sources)
	q.Add("syncToken", qp.SyncToken)

	urL.RawQuery = q.Encode()
	u_rl :=  urL.String()
	return u_rl, nil
}

