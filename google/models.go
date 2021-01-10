package google
//
//import (
//	"github.com/spf13/viper"
//	"log"
//	"net/url"
//)
//
//type Metadata struct{
//	Primary bool	    `json:"primary"` // "true" if verified and "false" if otherwise
//	Source struct{
//		Type string		`json:"type"`    //"CONTACT"
//		ID string 		`json:"id"`      //"48b4b2c68e9125a3"
//	}					`json:"source"`
//}
//
//type RespBody struct {
//
//	Connections []struct{
//		ResourceName string         `json:"resourceName"` // e.g "people/c5239008832403875235"
//		Etag 		 string 		`json:"etag"`   // e.g "%EgcBAj0LPzcuGgQBAgUHIgxUNjQrRDhoUmpHMD0="
//		Names []struct{
//			Metadata                `json:"metadata"`
//			DisplayName string      `json:"displayName"`
//			FamilyName string 		`json:"familyName"`
//			GivenName string 		`json:"givenName"`
//			MiddleName string 		`json:"middleName"`
//		}							`json:"names"`
//		PhoneNumbers []struct{
//			Metadata                `json:"metadata"`
//			CanonicalForm string 	`json:"canonicalForm"`  // e.g "+2347087404908"
//			Type string 			`json:"type"`			// e.g "mobile"
//		}							`json:"phoneNumbers"`
//
//	}					 			`json:"connections"`
//	NextPageToken string		    `json:"nextPageToken"`
//	NextSyncToken string			`json:"nextSyncToken"`
//	TotalPeople int64				`json:"totalPeople"`
//	TotalItems int64				`json:"totalItems"`
//
//}
//
//
//type sortOrfder struct {
//	FIRST_NAME_ASCENDING,
//	LAST_NAME_ASCENDING,
//	LAST_MODIFIED_ASCENDING,
//	LAST_MODIFIED_DESCENDING string
//}
//
//var SortOrder = sortOrfder{ "FIRST_NAME_ASCENDING", "LAST_NAME_ASCENDING", "LAST_MODIFIED_ASCENDING", "LAST_MODIFIED_DESCENDING",
//}
//
//
//type source struct {
//	READ_SOURCE_TYPE_UNSPECIFIED,
//	READ_SOURCE_TYPE_PROFILE,
//	READ_SOURCE_TYPE_CONTACT,
//	READ_SOURCE_TYPE_DOMAIN_CONTACT string
//}
//
//var Source = source{ "READ_SOURCE_TYPE_UNSPECIFIED", "READ_SOURCE_TYPE_PROFILE", "READ_SOURCE_TYPE_CONTACT", "READ_SOURCE_TYPE_DOMAIN_CONTACT",
//}
//
//
//
//
//type QueryParameters struct {
//	ApiKey string           			`json:"apiKey" validate:"required" viper:"API_KEY"`
//	ResourceName string     			`json:"resourceName"    validate:"required"`
//	PageToken string       			    `json:"pageToken"`
//	PageSize string 					`json:"pageSize"        validate:"required"`
//	PersonFields string  			    `json:"personFields"    validate:"required"`
//	RequestMaskIncludeField string		`json:"requestMask.includeField"`
//	RequestSyncToken string 			`json:"requestSyncToken"`
//	SortOrder string  	 				`json:"sortOrder"        validate:"required"`
//	Sources string 						`json:"sources"`
//	SyncToken string					`json:"syncToken"`
//	PrettyPrint string                  `json:"prettyPrint"`
//}
//
//
//func (qp *QueryParameters) SetApiKey(){
//	config, err := loadConfig("C:/gocode/src/github.com/IamNator/mysql-golang-web/" )
//	if err != nil {
//		log.Fatal("cannot load config:", err)
//	}
//	qp.ApiKey = config.ApiKey
//}
//
////extracts important information from app.env file { I should gitignore this file}
//func loadConfig(path string) (config QueryParameters, err error) {
//
//	viper.AddConfigPath(path)
//	viper.SetConfigName("app")
//	viper.SetConfigType("env")
//
//	viper.AutomaticEnv()
//
//	err = viper.ReadInConfig()
//	if err != nil {
//		return
//	}
//	err = viper.Unmarshal(&config)
//
//	return
//}
//
//
//
//
//func (qp *QueryParameters) SetURL() (string, error) {
//	 var urL = url.URL{
//			 Scheme: "https",
//			 Host: "people.googleapis.com",
//			 Path: "v1/people/me/connections",
//	}
//
//	//https://accounts.google.com/o/oauth2/auth?response_type=permission%20id_token&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fcontacts.readonly&openid.realm=&authuser=0&redirect_uri=storagerelay%3A%2F%2Fhttps%2Fexplorer.apis.google.com%3Fid%3Dauth791925&client_id=292824132082.apps.googleusercontent.com&ss_domain=https%3A%2F%2Fexplorer.apis.google.com&gsiwebsdk=shim
//
//	q := urL.Query()
//	q.Add("key", qp.ApiKey)
//	q.Set("resourceName", qp.ResourceName)
//	q.Add("pageToken", qp.PageToken)
//	q.Add("pageSize", qp.PageSize)
//	q.Add("personFields", qp.PersonFields)
//	q.Add("requestMask.includeField", qp.RequestMaskIncludeField)
//	q.Add("requestSyncToken", qp.RequestSyncToken)
//	q.Add("sortOrder", string(qp.SortOrder) )
//	q.Add("sources", qp.Sources)
//	q.Add("syncToken", qp.SyncToken)
//
//	urL.RawQuery = q.Encode()
//	u_rl :=  urL.String()
//	return u_rl, nil
//}
//
//
////https://developers.google.com/people/api/rest/v1/people.connections/list?apix_params=%7B%22resourceName%22%3A%22people%2Fme%22%2C%22pageSize%22%3A1000%2C%22personFields%22%3A%22names%2CphoneNumbers%22%2C%22sortOrder%22%3A%22FIRST_NAME_ASCENDING%22%2C%22sources%22%3A%5B%22READ_SOURCE_TYPE_CONTACT%22%5D%2C%22prettyPrint%22%3Atrue%7D#authorization-scopes
