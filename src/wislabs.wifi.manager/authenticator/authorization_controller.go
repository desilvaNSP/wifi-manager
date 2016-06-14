package authenticator
import (
	"net/http"
	"encoding/json"
)

/*Permission Constant*/
const CSV_DOWNLOAD  = "csv_download"
const WIFI_LOCATION = "wifi_location"
const WIFI_USERS  =  "wifi_users"
const DASHBOARD_USERS =  "dashboard_users"
const ADMIN =  "admin"

const ACTION_EXECUTE string = "execute"
const ACTION_WRITE string = "write"
const ACTION_READ string = "read"

type Permission struct {
	permission string
}
/**
* get scope from jwt and check for permission
* "scopes": {
*    "wifi_location": [
*      "read",
*     "write",
*      "execute"
*    ]
*  }
*/
func IsAuthorized(resourceId string, permission string, r *http.Request) bool {
	m1 := make(map[string][]string)
	json.Unmarshal([]byte(r.Header.Get("scopes")), &m1)
	m2 := m1[resourceId]
	if m2 != nil {
		for _, element := range m2 {
			if element == permission {
				return true
			}
		}
	}
	return false
}

func IsUserAuthorized(username string, resourceId string, permission string, r *http.Request) bool {
	m1 := make(map[string][]string)
	json.Unmarshal([]byte(r.Header.Get("scopes")), &m1)

	m2 := m1[resourceId]
	if m2 != nil && username == r.Header.Get("username"){
		for _, element := range m2 {
			if element == permission {
				return true
			}
		}
	}
	return false
}