package handlers

import (
	"encoding/json"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"wislabs.wifi.manager/controllers/dashboard"
	"strconv"
	"wislabs.wifi.manager/dao"
)

/**
* POST
* @path dashboard/apps/
*
*/

func CreateDashboardApp(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var dashboarApp dao.DashboardAppInfo
	err := decoder.Decode(&dashboarApp)
	if(err != nil){
		log.Fatalln("Error while decoding location json")
	}
	dashboard.CreateNewDashboardApp(dashboarApp)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

/**
* GET
* @path dashboard/{tenantid}/apps/{username}
*
*/
func GetAppsOfUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	username := vars["username"]
	tenantid, err := strconv.Atoi(vars["tenantid"])
	if(err!= nil){
		log.Fatalln("Error while reading tenantid", err)
	}
	apps := dashboard.GetAllDashboardAppsOfUser(username, tenantid)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(apps); err != nil {
		panic(err)
	}
}

/**
* GET
* @path dashboard/apps/{appid}
*
*/
func GetAllUsersOfApp(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	appId, err := strconv.Atoi(vars["appid"])
	if(err!= nil){
		log.Fatalln("Error while reading tenantid", err)
	}
	apps := dashboard.GetAllDashboardUsersOfApp(appId)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(apps); err != nil {
		panic(err)
	}
}

/**
* DELETE
* @path dashboard/{tenantid}/apps/{appid}/
*
*/
func DeleteDashboardApp(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	appId, err := strconv.Atoi(vars["appid"])
	tenantid, err := strconv.Atoi(vars["tenantid"])
	if(err!= nil){
		log.Fatalln("Error while reading tenantid", err)
	}
	err = dashboard.DeleteDashboardApp(appId, tenantid)

	if err != nil {
		log.Fatalln("Error while deleting dashboard app from DB ", err)
		w.WriteHeader(http.StatusNotFound)
	}else{
		w.WriteHeader(http.StatusOK)
	}
}