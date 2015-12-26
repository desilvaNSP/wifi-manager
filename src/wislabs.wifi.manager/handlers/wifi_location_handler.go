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
)

func GetLocations(w http.ResponseWriter, r *http.Request){
	if(!authenticator.IsAutherized("wifi_location", authenticator.ACTION_READ,r)){
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	locations := location.GetAllLocations()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(locations); err != nil {
		panic(err)
	}
}

/**
* POST
* @path /locations
* return
*/
func AddLocation(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var aplocation dao.ApLocation
	err := decoder.Decode(&aplocation)
	if(err != nil){
		log.Fatalln("Error while decoding location json")
	}
	location.AddLocation(&aplocation)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

/**
* DELETE
* @path /locations/{mac}/{ssid}/
* return
*/
func DeleteLocation(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	ssid := vars["ssid"]
	mac := vars["mac"]
	err :=location.DeleteApLocation(ssid, mac)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		log.Fatalln("Error while deleting location " + ssid +" from DB ", err)
		w.WriteHeader(http.StatusNotFound)
	}else{
		w.WriteHeader(http.StatusOK)
	}
}

/**
* DELETE
* @path /locations/{mac}
* return
*/
func DeleteAccessPoint(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	mac := vars["mac"]

	err := location.DeleteAccessPoint(mac)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		log.Fatalln("Error while deleting accesspoint : " + mac +" from DB ", err)
		w.WriteHeader(http.StatusNotFound)
	}else{
		w.WriteHeader(http.StatusOK)
	}
}