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

func IsAutherized(resourceId string, permission string, r *http.Request) bool{
	m1 := make(map[string][]string)
	print(r.Header.Get("scopes"))
	json.Unmarshal([]byte(r.Header.Get("scopes")),&m1)
    m2 := m1[resourceId]
	for i :=0; i<3 ; i++{
		if(m2[i]==permission){
			return true
		}
	}
	return false
}