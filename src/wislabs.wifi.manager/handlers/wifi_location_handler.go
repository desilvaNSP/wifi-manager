package handlers

import (
	"wislabs.wifi.manager/dao"
	"wislabs.wifi.manager/authenticator"
	"encoding/json"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"wislabs.wifi.manager/controllers/location"
	"strconv"
)

func GetLocations(w http.ResponseWriter, r *http.Request){
	if(!authenticator.IsAuthorized("wifi_location", authenticator.ACTION_READ,r)){
		w.WriteHeader(http.StatusForbidden)
		return
	}
	vars := mux.Vars(r)
	tenantid, err := strconv.Atoi(vars["tenantid"])
	if(err!= nil){
		log.Fatalln("Error while reading tenantid", err)
	}
	locations := location.GetAllLocations(tenantid)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(locations); err != nil {
		panic(err)
	}
}

func GetLocationGroups(w http.ResponseWriter, r *http.Request){
	if (!authenticator.IsAuthorized("wifi_location", authenticator.ACTION_READ, r)) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	vars := mux.Vars(r)
	tenantid, err := strconv.Atoi(vars["tenantid"])
	if(err!= nil){
		log.Fatalln("Error while reading tenantid", err)
	}
	locationGroups := location.GetAllLocationGroups(tenantid)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(locationGroups); err != nil {
		panic(err)
	}
}

/**
* POST
* @path /wifi/locations
* return
*/
func AddWiFiLocationHandler(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var apLocation dao.ApLocation
	err := decoder.Decode(&apLocation)
	if(err != nil){
		log.Fatalln("Error while decoding location json")
	}
	location.AddWiFiLocation(&apLocation)
	w.WriteHeader(http.StatusOK)
}

/**
* POST
* @path /wifi/locationsupdate
* return
*/
func UpdateWiFiLocationHandler(w http.ResponseWriter,r *http.Request)  {
	decoder := json.NewDecoder(r.Body)
	var apLocation dao.ApLocation
	err := decoder.Decode(&apLocation)
	if(err != nil){
		log.Fatalln("Error while decoding location json")
	}
	location.UpdateWifiLocation(&apLocation)
	w.WriteHeader(http.StatusOK)
}

/**
* POST
* @path /{tenantid}/locations
* return
*/
func AddWiFiGroupHandler(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var apGroup dao.ApGroup
	err := decoder.Decode(&apGroup)
	if(err != nil){
		log.Fatalln("Error while decoding location json")
	}
	location.AddWiFiGroup(&apGroup)
	w.WriteHeader(http.StatusOK)
}


/**
* DELETE
* @path /{tenantid}/locations/{mac}/{ssid}/{groupname}
* return
*/
func DeleteLocation(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	tenantid, err := strconv.Atoi(vars["tenantid"])
	if(err!= nil){
		log.Fatalln("Error while reading tenantid", err)
	}
	ssid := vars["ssid"]
	mac := vars["mac"]
	groupname := vars["groupname"]
	err =location.DeleteApLocation(tenantid, ssid, mac, groupname)

	if err != nil {
		log.Fatalln("Error while deleting location " + ssid +" from DB ", err)
		w.WriteHeader(http.StatusNotFound)
	}else{
		w.WriteHeader(http.StatusOK)
	}
}

/**
* DELETE
* @path /{tenantid}/locations/{groupname}
* return
*/
func DeleteLocationGroup(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	tenantid, err := strconv.Atoi(vars["tenantid"])
	if(err!= nil){
		log.Fatalln("Error while reading tenantid", err)
	}
	groupname := vars["groupname"]
	err =location.DeleteApGroup(groupname, tenantid)

	if err != nil {
		log.Fatalln("Error while deleting location " + groupname +" from DB ", err)
		w.WriteHeader(http.StatusNotFound)
	}else{
		w.WriteHeader(http.StatusOK)
	}
}

/**
* DELETE
* @path /{tenantid}/locations/{mac}
* return
*/
func DeleteAccessPoint(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	tenantid, err := strconv.Atoi(vars["tenantid"])
	if(err!= nil){
		log.Fatalln("Error while reading tenantid", err)
	}
	mac := vars["mac"]

	err = location.DeleteAccessPoint(mac, tenantid)
	if err != nil {
		log.Fatalln("Error while deleting accesspoint : " + mac +" from DB ", err)
		w.WriteHeader(http.StatusNotFound)
	}else{
		w.WriteHeader(http.StatusOK)
	}
}

/**
* POST
* @path /wifi/summary/locationpoint,
*
*/

/*
func GetAPLongLatandSummaryDataFromToHandler(w http.ResponseWriter, r*http.Request){
	decoder := json.NewDecoder(r.Body)
	var constrains dao.Constrains
	decoder.Decode(&constrains)

	locationData := location.GetAccessPointLongLatByMac()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(accespoint); err != nil {
		panic(err)
	}
}*/
