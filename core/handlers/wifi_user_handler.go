package handlers

import (
	wifi_controller "dashboard-core/controllers/wifi"
	"dashboard-core/dao"
	"encoding/json"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"strconv"
)


/**
* POST
* @path /users
*/
func AddUserHandler(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var user dao.RadiusUser
	decoder.Decode(&user)
	wifi_controller.AddWiFiUser(&user)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

/**
* GET
* @path /users
* return [{"id":0,"username":"anu","password":"","acctstarttime":"2015-09-20 18:49:32",
*         "acctlastupdatedtime":"2015-09-20 18:49:32","acctactivationtime":"","acctstoptime":"2015-09-20 19:49:32"}]
*/
func GetUsersHandler(w http.ResponseWriter, r *http.Request){

	users := wifi_controller.GetAllWiFiUsers()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}

/**
* DELETE
* @path /users/<user-id>
*/
func DeleteUserHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	username := vars["username"]
	err := wifi_controller.DeleteUserAccountingSession(username)
	err = wifi_controller.DeleteUserFromRadAcct(username)
	err = wifi_controller.DeleteUserFromRadCheck(username)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		log.Fatalln("Error while deleting user from accounting table" + username +" from DB ", err)
		w.WriteHeader(http.StatusNotFound)
	}else{
		w.WriteHeader(http.StatusOK)
	}

	if err := json.NewEncoder(w).Encode("{}"); err != nil {
		panic(err)
	}
}

/**
* GET
* @path /users/<user-id>
*/
func GetUserHandler(w http.ResponseWriter, r *http.Request){

}

/**
* POST
* @path /users
*/
func UpdateUserHandler(){

}

/**
* POST
* @path /users
* return [{"id":0,"username":"anu","password":"","acctstarttime":"2015-09-20 18:49:32",
*         "acctlastupdatedtime":"2015-09-20 18:49:32","acctactivationtime":"","acctstoptime":"2015-09-20 19:49:32"}]
*/
func GetUsersCountFromToHandler(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var constrains dao.Constrains
	err := decoder.Decode(&constrains)

    count := wifi_controller.GetUsersCountFromTo(constrains.From, constrains.To, constrains.LocationId)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(count); err != nil {
		panic(err)
	}
}

func GetReturningUsersCountFromToHandler(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var constrains dao.Constrains
	err := decoder.Decode(&constrains)

	count := wifi_controller.GetReturningUsers(constrains.From, constrains.To, constrains.LocationId)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(count); err != nil {
		panic(err)
	}
}

func GetUserCountOfDownloadsOverHandler(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var constrains dao.Constrains
	decoder.Decode(&constrains)
	vars := mux.Vars(r)
	threshold := vars["threshold"]
	value,_  := strconv.Atoi(threshold)
	count := wifi_controller.GetUserCountOfDownloadsOver(constrains,value)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(count); err != nil {
		panic(err)
	}
}


func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}