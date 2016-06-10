package handlers

import (
	"net/http"
	"encoding/json"
	"wislabs.wifi.manager/dao"
	"wislabs.wifi.manager/controllers/radius"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"strconv"
)

/**
* POST
* @path /radius/user/connection
*
*/
func TestRadiusAuthConnection(w http.ResponseWriter, r*http.Request){
	decoder := json.NewDecoder(r.Body)
	var nasClientInfo dao.NasClientTestInfo
	decoder.Decode(&nasClientInfo)

	status := radius.TestAuthenticationOnUser(nasClientInfo);

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(status); err != nil {
		panic(err)
	}
}

/**
* POST
* @path /radius/server
*
*/
func CreateRadiusServerHandler(w http.ResponseWriter, r*http.Request){
	decoder := json.NewDecoder(r.Body)
	var radiusConfig dao.RadiusServer
	err := decoder.Decode(&radiusConfig)
	if err != nil {
		panic("Error while decoding json")
	}
	err_db := radius.CreateRadiusServer(radiusConfig);
	if err_db!= nil {
		checkErr(err_db,"Error happening while create radius server")
	}
	w.WriteHeader(http.StatusOK)
}

/**
* GET
* @path /radius/{tenantid}/radiusdetails
*
*/
func GetRadiusServerDetailsHandler(w http.ResponseWriter, r*http.Request){
	vars := mux.Vars(r)
	tenantId, err := strconv.Atoi(vars["tenantid"])
	if(err!= nil){
		log.Error("Error while reading tenantid", err)
	}
	allRadiusDetails, err := radius.GetAllRadiusDetails(tenantId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(allRadiusDetails); err != nil {
		checkErr(err, "Error occured while getting Radius Details")
	}
}

/**
* GET
* @path /radius/server/clients/{instanceid}
*
*/
func GetRadiusClientsInServerHanlder(w http.ResponseWriter, r*http.Request){
	vars := mux.Vars(r)
	instanceId, err := strconv.Atoi(vars["instanceid"])
	if(err!= nil){
		log.Error("Error while reading instanceid", err)
	}

	tenantId, err := strconv.Atoi(r.Header.Get("tenantid"))
	if (err != nil) {
		log.Error("Error while reading tenantid", err)
	}

	radiusConfig, err := radius.GetInstanceConfigsById(instanceId, tenantId)
	if err != nil {
		checkErr(err,"Error while get server instance configs by ID")
	}
	nasClients, err := radius.GetRadiusClientsInServer(radiusConfig)
	if err != nil {
		checkErr(err, "Error occured while getting Nas Clients in Server ")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(nasClients); err != nil {
			checkErr(err, "Error occured while getting Nas Clients in Server")
		}
	}else{
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(nasClients); err != nil {
			checkErr(err, "Error occured while getting Nas Clients in Server")
		}
	}

}

/**
* DELETE
* @path /radius/server/{serverinstid}
*
*/
func DeleteRadiusInstanceHandler(w http.ResponseWriter, r*http.Request){
	vars := mux.Vars(r)
	radiusInstId := vars["serverinstid"]
	tenantId, err := strconv.Atoi(r.Header.Get("tenantid"))
	if (err != nil) {
		log.Error("Error while reading tenantid", err)
	}
	errOnDelete := radius.DeleteRadiusInstance(tenantId,radiusInstId)
	if errOnDelete != nil {
		checkErr(errOnDelete,"Error while deleting radius server.")
	}
	w.WriteHeader(http.StatusOK)
}

/**
* PUT
* @path /radius/server
*
*/
func UpdateRadiusServerHandler(w http.ResponseWriter, r*http.Request)  {
	decoder := json.NewDecoder(r.Body)
	var radServerConfig dao.RadiusServer
	decoder.Decode(&radServerConfig)

	tenantId, err := strconv.Atoi(r.Header.Get("tenantid"))
	if (err != nil) {
		log.Error("Error while reading tenantid", err)
	}
	errOnUpdate := radius.UpdateRadiusServerInstance(radServerConfig, tenantId)
	if errOnUpdate != nil{
		checkErr(errOnUpdate,"Error while updating radius server instance")
	}
	w.WriteHeader(http.StatusOK)
}

/**
* POST
* @path /radius/server/client
*
*/
func CreateNASClientOnServerHandler(w http.ResponseWriter, r*http.Request){
	decoder := json.NewDecoder(r.Body)
	var nasClientServerInfo dao.NASClientDBServer
	err := decoder.Decode(&nasClientServerInfo)
	if err != nil {
		panic("Error while decoding json")
	}
	var radiusServerConfigs dao.RadiusServer
	var nasClientInfo dao.NasClient
	radiusServerConfigs =  nasClientServerInfo.RadiusServerInfo;
	nasClientInfo =  nasClientServerInfo.NASClientInfo

	err_db := radius.CreateNASClientOnServer(radiusServerConfigs, nasClientInfo);
	if err_db!= nil {
		checkErr(err_db,"Error happening while create radius server")
	}
	w.WriteHeader(http.StatusOK)
}

/**
* GET
* @path /radius/users/{username}
*
*/
func WifiUserValidInRadiusHanlder(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	username := vars["username"]
	tenantId, err := strconv.Atoi(r.Header.Get("tenantid"))
	if (err != nil) {
		log.Error("Error while reading tenantid", err)
	}
	var isValid int
	isValid, err = radius.IsWifiUserValidInRadius(tenantId, username);
	if err != nil {
		checkErr(err,"Error happening while Wifi user is valid in Radius")
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(isValid); err != nil {
		checkErr(err, "Error happening while JSON encoding.")
	}
}

/**
* GET
* @path /radius/server/{instanceid}/validnas/{ipstring}/{rangesize}
*
*/
func NASIpExistInRadiusHandler(w http.ResponseWriter, r *http.Request){
	var isValid bool
	vars := mux.Vars(r)
	ipAddress := vars["ipstring"]
	lengthRange := vars["rangesize"]
	instanceId, err := strconv.Atoi(vars["instanceid"])
	if(err!= nil){
		log.Error("Error while reading instanceid ", err)
	}

	tenantId, err := strconv.Atoi(r.Header.Get("tenantid"))
	if (err != nil) {
		log.Error("Error while reading tenantid", err)
	}

	radiusConfig, err := radius.GetInstanceConfigsById(instanceId, tenantId)
	if err != nil {
		checkErr(err,"Error while get server instance configs by ID")
	}

	isValid, err = radius.IsNASIpExistsInRadius(radiusConfig, ipAddress,lengthRange);
	if err != nil {
		checkErr(err,"Error happening while Wifi user is valid in Radius")
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(isValid); err != nil {
		checkErr(err, "Error happening while JSON encoding.")
	}
}


