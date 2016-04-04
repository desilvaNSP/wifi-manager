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
	"fmt"
)

/**
* POST
* @path dashboard/apps/
*
*/
func CreateDashboardApp(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var dashboardApp dao.DashboardAppInfo
	err := decoder.Decode(&dashboardApp)
	if(err != nil){
		log.Fatalln("Error while decoding location json")
	}
	dashboard.CreateNewDashboardApp(dashboardApp)
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
	tenantId, err := strconv.Atoi(vars["tenantid"])
	if(err!= nil){
		log.Fatalln("Error while reading tenantid", err)
	}
	apps := dashboard.GetAllDashboardAppsOfUser(username, tenantId)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(apps); err != nil {
		panic(err)
	}
}

/**
* GET
* @path dashboard/apps/{appid}/users
*
*/
func GetUsersOfApp(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	appId, err := strconv.Atoi(vars["appid"])
	if(err!= nil){
		log.Fatalln("Error while reading tenantid", err)
	}
	appUsers := dashboard.GetDashboardUsersOfApp(appId)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(appUsers); err != nil {
		panic(err)
	}
}

/**
* GET
* @path dashboard/apps/{appid}/metrics
*
*/
func GetMetricsOfApp(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	appId, err := strconv.Atoi(vars["appid"])
	if(err!= nil){
		log.Fatalln("Error while reading tenantid", err)
	}
	appMetrics := dashboard.GetDashboardMetricsOfApp(appId)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(appMetrics); err != nil {
		panic(err)
	}
}

/**
* GET
* @path dashboard/apps/{appid}/groups
*
*/
func GetGroupsOfApp(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	appId, err := strconv.Atoi(vars["appid"])
	if(err!= nil){
		log.Fatalln("Error while reading tenantid", err)
	}
	appGroups := dashboard.GetDashboardGroupsOfApp(appId)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(appGroups); err != nil {
		panic(err)
	}
}
/**
* GET
* @path dashboard/apps/{appid}/groups
*
*/
func GetAclsOfApp(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	appId, err := strconv.Atoi(vars["appid"])
	if(err!= nil){
		log.Fatalln("Error while reading appid", err)
	}
	appacls := dashboard.GetDashboardAclsOfApp(appId)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(appacls); err != nil {
		panic(err)
	}
}

/**
* GET
* @path dashboard/{tenantid}/metrics
*
*/
func GetAllDashboardMetrics(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	tenantId, err := strconv.Atoi(vars["tenantid"])
	if(err!= nil){
		log.Fatalln("Error while reading tenantid", err)
	}
	metrics := dashboard.GetAllDashboardMetrics(tenantId)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(metrics); err != nil {
		panic(err)
	}
}


/**
* GET
* @path dashboard/{tenantid}/metrics
*
*/

func GetAclTypes(w http.ResponseWriter, r *http.Request){

	fmt.Printf("sasasas @ 2")
	aclTypes := dashboard.GetAllDashboardAclTypes()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(aclTypes); err != nil {
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
	tenantId, err := strconv.Atoi(vars["tenantid"])
	if(err!= nil){
		log.Fatalln("Error while reading tenantid", err)
	}
	err = dashboard.DeleteDashboardApp(appId, tenantId)

	if err != nil {
		log.Fatalln("Error while deleting dashboard app from DB ", err)
		w.WriteHeader(http.StatusNotFound)
	}else{
		w.WriteHeader(http.StatusOK)
	}
}