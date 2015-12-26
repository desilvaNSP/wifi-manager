package main

import (
	"database/sql"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"wislabs.wifi.manager/dao"
	"wislabs.wifi.manager/routes"
	"wislabs.wifi.manager/utils"
)

var m *mux.Router
var req *http.Request
var err error
var respRec *httptest.ResponseRecorder
var username string

func setup() {

	loadConfigs("/home/anuruddha/git/wifi-manager/server")
	username = "erty"
	//mux router with added question routes
	m = routes.NewRouter()
	//The response recorder used to record HTTP responses
	respRec = httptest.NewRecorder()
}

func TestCreateUser(t *testing.T) {
	setup()
	radiusUser := dao.PortalUser{}
	radiusUser.Username = username
	radiusUser.Location = utils.NullInt64{sql.NullInt64{2, true}}
	//Testing get of non existent question type

	b, err := json.Marshal(radiusUser)
	req, err = http.NewRequest("POST", "/wifi/users", strings.NewReader(string(b)))
	if err != nil {
		t.Fatal("Creating 'POST /questions/1/SC' request failed!")
	}

	m.ServeHTTP(respRec, req)
	if respRec.Code != http.StatusOK {
		//TestDeleteUser(t)
		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusBadRequest)
	}

}

func TestGetUsers(t *testing.T) {
	setup()
	//Testing get of non existent question type
	req, err = http.NewRequest("GET", "/wifi/users", nil)
	if err != nil {
		t.Fatal("Creating 'GET /questions/1/SC' request failed!")
	}

	m.ServeHTTP(respRec, req)
	log.Info(respRec)
	if respRec.Code != http.StatusOK {
		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusBadRequest)
	}
}

func TestDeleteUser(t *testing.T) {
	setup()
	//Testing get of non existent question type
	req, err = http.NewRequest("DELETE", "/wifi/users/"+username, nil)
	if err != nil {
		t.Fatal("Creating 'GET /questions/1/SC' request failed!")
	}

	m.ServeHTTP(respRec, req)

	if respRec.Code != http.StatusOK {
		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusBadRequest)
	}
}
