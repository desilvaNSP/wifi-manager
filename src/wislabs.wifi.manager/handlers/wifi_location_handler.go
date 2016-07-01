package handlers

import (
	"wislabs.wifi.manager/dao"
	"wislabs.wifi.manager/authenticator"
	"encoding/json"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"wislabs.wifi.manager/controllers/location"
	"strconv"
	"strings"
	"wislabs.wifi.manager/commons"
)

func GetLocations(w http.ResponseWriter, r *http.Request) {
	if (!authenticator.IsAuthorized(authenticator.WIFI_LOCATION, authenticator.ACTION_READ, r)) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	tenantId, err := strconv.Atoi(r.Header.Get("tenantid"))
	if (err != nil) {
		checkErr(err,"Error while getting getting tenantid")
	}
	locations := location.GetAllLocations(tenantId)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(locations); err != nil {
		panic(err)
	}
}

func GetAPs(w http.ResponseWriter, r *http.Request) {
	if (!authenticator.IsAuthorized(authenticator.WIFI_LOCATION, authenticator.ACTION_READ, r)) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	tenantId, err := strconv.Atoi(r.Header.Get("tenantid"))
	if (err != nil) {
		checkErr(err,"Error while getting getting tenantid")
	}

	aps, err := location.GetAllAPs(tenantId)

	checkErr(err, "Error while getting all aps")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(aps); err != nil {
		checkErr(err, "Error while encoding response")
	}
}

func GetLocationGroups(w http.ResponseWriter, r *http.Request) {
	if (!authenticator.IsAuthorized(authenticator.WIFI_LOCATION, authenticator.ACTION_READ, r)) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	tenantId, err := strconv.Atoi(r.Header.Get("tenantid"))
	if (err != nil) {
		checkErr(err,"Error while getting getting tenantid")
	}
	locationGroups := location.GetAllLocationGroups(tenantId)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(locationGroups); err != nil {
		panic(err)
	}
}

func GetSSIDsOfAPGroups(w http.ResponseWriter, r *http.Request) {
	tenantId, err := strconv.Atoi(r.Header.Get("tenantid"))
	if (err != nil) {
		log.Fatalln("Error while reading tenantid", err)
	}
	groupnames := strings.Split(r.FormValue("groupnames"), ",")
	ssids := location.GetSSIDsOfLocationGroups(groupnames, tenantId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ssids); err != nil {
		panic(err)
	}
}

func IsSSIDExist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	macaddress := vars["mac"]
	ssid := vars["ssid"]
	tenantId, err := strconv.Atoi(r.Header.Get("tenantid"))
	if (err != nil) {
		checkErr(err,"Error while getting getting tenantid")
	}
	valid , err := location.IsSSIDExistsOnMac(macaddress, ssid, tenantId)
	checkErr(err,"Error while getting SSIDs on MAC")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(valid); err != nil {
		checkErr(err,"Error while decoding SSIDs on MAC")
	}
}

func GetAPsByMac(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	macaddress := vars["mac"]
	tenantId, err := strconv.Atoi(r.Header.Get("tenantid"))
	if (err != nil) {
		checkErr(err,"Error while getting getting tenantid")
	}
	aps, err := location.GetAPsOnLocation(macaddress, tenantId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(aps); err != nil {
		checkErr(err,"Error while encoding response")
	}
}

func GetMACs(w http.ResponseWriter, r *http.Request) {
	tenantId, err := strconv.Atoi(r.Header.Get("tenantid"))
	if (err != nil) {
		checkErr(err,"Error while getting getting tenantid")
	}
	macs, err := location.GetMACsOnAllLocations(tenantId)
	if err != nil {
		checkErr(err, "Error while getting all macs over locations.")
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(macs); err != nil {
		panic(err)
	}
}

/**
* POST
* @path /wifi/locations
* return
*/
func AddWiFiLocationHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var apLocation dao.ApLocationSSIDs
	err := decoder.Decode(&apLocation)
	if (err != nil) {
		log.Fatalln("Error while decoding location json")
	}
	err2 := location.AddWiFiLocation(&apLocation)
	if err != nil {
		checkErr(err2, "Error while adding location")
	}
	w.WriteHeader(http.StatusOK)
}

/**
* PUT
* @path /wifi/locations
* return
*/
func UpdateWiFiLocationHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var apLocation dao.ApLocationSSIDs
	err := decoder.Decode(&apLocation)
	if (err != nil) {
		log.Fatalln("Error while decoding location json")
	}
	err2 := location.UpdateWifiLocation(&apLocation)
	checkErr(err2, "Error while updating wifi location")
	w.WriteHeader(http.StatusOK)
}


func UpdateAPsHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var aps dao.APs
	err := decoder.Decode(&aps)
	if (err != nil) {
		log.Fatalln("Error while decoding location json")
	}
	err2 := location.UpdateAPs(&aps)
	checkErr(err2, "Error while updating aps")
	w.WriteHeader(http.StatusOK)
}

/**
* POST
* @path /{tenantid}/locations
* return
*/
func AddWiFiGroupHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var apGroup dao.ApGroup
	err := decoder.Decode(&apGroup)
	if (err != nil) {
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
func DeleteLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenantid, err := strconv.Atoi(vars["tenantid"])
	if (err != nil) {
		log.Fatalln("Error while reading tenantid", err)
	}
	ssid := vars["ssid"]
	mac := vars["mac"]
	groupname := vars["groupname"]
	err = location.DeleteApLocation(tenantid, ssid, mac, groupname)

	if err != nil {
		log.Fatalln("Error while deleting location " + ssid + " from DB ", err)
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

/**
* DELETE
* @path /{tenantid}/locations/{groupname}
* return
*/
func DeleteLocationGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenantid, err := strconv.Atoi(vars["tenantid"])
	if (err != nil) {
		log.Fatalln("Error while reading tenantid", err)
	}
	groupname := vars["groupname"]
	err = location.DeleteApGroup(groupname, tenantid)

	if err != nil {
		log.Fatalln("Error while deleting location " + groupname + " from DB ", err)
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

/**
* DELETE
* @path /{tenantid}/locations/{mac}
* return
*/
func DeleteAccessPoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	tenantid, err := strconv.Atoi(vars["tenantid"])
	if (err != nil) {
		log.Fatalln("Error while reading tenantid", err)
	}
	mac := vars["mac"]

	err = location.DeleteAccessPoint(mac, tenantid)
	if err != nil {
		log.Fatalln("Error while deleting accesspoint : " + mac + " from DB ", err)
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

/**
* GET
* @path /wifi/ap/activecount
* return
*/

func GetActiveAPHandler(w http.ResponseWriter, r *http.Request) {
	constraints := getConstrains(r,commons.GET_ACTIVE_APS_COUNT)
	var countActiveAP int
	var err error
	countActiveAP, err = location.GetAccessPointFeatureDetails(constraints)
	checkErr(err, "Error occourred while getting active ap count ")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(countActiveAP); err != nil {
		panic(err)
	}
}

/**
* GET
* @path /wifi/ap/inactivecount
* return
*/

func GetInactiveAPHandler(w http.ResponseWriter, r *http.Request) {
	constraints := getConstrains(r,commons.GET_INACTIVE_APS_COUNT)
	var countInactiveAP int

	var err error
	countInactiveAP, err = location.GetAccessPointFeatureDetails(constraints)
	checkErr(err, "Error occourred while getting inactive ap count ")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(countInactiveAP); err != nil {
		panic(err)
	}
}

/**
* GET
* @path /wifi/ap/distinctcount
* return
*/

func GetDistinctMacCountHandler(w http.ResponseWriter, r *http.Request) {
	constraints := getConstrains(r,commons.GET_DISTINCT_MAC)
	var distinctMac int

	var err error
	distinctMac, err = location.GetAccessPointFeatureDetails(constraints)
	checkErr(err, "Error occourred while getting active ap count ")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(distinctMac); err != nil {
		panic(err)
	}
}

func getConstrains(r *http.Request,	query string) dao.AccessPointConstraints{
	var constraints dao.AccessPointConstraints
	constraints.To = r.URL.Query().Get("to")
	constraints.From = r.URL.Query().Get("from")
	threshold := r.URL.Query().Get("threshold")
	var err error
	if len(threshold) != 0  {
		constraints.Threshold, err = strconv.Atoi(r.URL.Query().Get("threshold"));
		checkErr(err, "Error while reading treshold")
	}
	constraints.TenantId, err = strconv.Atoi(r.Header.Get("tenantid"))
	checkErr(err, "Error while reading tenantid")
	constraints.Query = query

	return constraints
}

