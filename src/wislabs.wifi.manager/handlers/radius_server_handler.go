package handlers

import (
	"net/http"
	"encoding/json"
	"wislabs.wifi.manager/dao"
	"wislabs.wifi.manager/controllers/radius"
	"github.com/gorilla/mux"
	"log"
	"strconv"

)


func TestRadiusAuthConnection(w http.ResponseWriter, r*http.Request){
	decoder := json.NewDecoder(r.Body)
	var radiusConfig dao.RadiusConfigsInfo
	decoder.Decode(&radiusConfig)

	status := radius.TestAuthenticationOnUser(radiusConfig);

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(status); err != nil {
		panic(err)
	}
}

func CreateRadiusServerHandler(w http.ResponseWriter, r*http.Request){
	decoder := json.NewDecoder(r.Body)
	var radiusConfig dao.RadiusConfigsInfo
	decoder.Decode(&radiusConfig)

	status := radius.TestAuthenticationOnUser(radiusConfig)
	if(status){
		radius.CreateRadiusServer(radiusConfig);
	}
	w.WriteHeader(http.StatusOK)
}


func GetRadiusServerDetailsHandler(w http.ResponseWriter, r*http.Request){
	vars := mux.Vars(r)
	tenantid, err := strconv.Atoi(vars["tenantid"])
	if(err!= nil){
		log.Fatalln("Error while reading tenantid", err)
	}
	username := vars["username"]

	allRadiusDetails := radius.GetAllRadiusDetails(tenantid,username)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(allRadiusDetails); err != nil {
		panic(err)
	}
}


func DeleteRadiusInstanceHandler(w http.ResponseWriter, r*http.Request){
	vars := mux.Vars(r)
	tenantid, err := strconv.Atoi(vars["tenantid"])
	if(err!= nil){
		log.Fatalln("Error while reading tenantid", err)
	}
	radiusinstid := vars["radiusinstid"]

	allRadiusDetails := radius.DeleteRadiusInstance(tenantid,radiusinstid)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(allRadiusDetails); err != nil {
		panic(err)
	}
}

func UpdateRadiusInstanceHandler(w http.ResponseWriter, r*http.Request)  {
	vars := mux.Vars(r)
	tenantid, err := strconv.Atoi(vars["tenantid"])
	if(err!= nil){
		log.Fatalln("Error while reading tenantid", err)
	}
	radiusinstid := vars["radiusinstid"]

	decoder := json.NewDecoder(r.Body)
	var radiusConfig dao.RadiusConfigs
	decoder.Decode(&radiusConfig)

	allRadiusDetails := radius.UpdateRadiusInstance(radiusConfig,tenantid,radiusinstid)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(allRadiusDetails); err != nil {
		panic(err)
	}
}