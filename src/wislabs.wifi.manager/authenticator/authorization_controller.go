package authenticator
import (
	"net/http"
	"encoding/json"
)

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