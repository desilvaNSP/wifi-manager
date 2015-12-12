package handlers

import (
	"wifi-manager/dao"
	"wifi-manager/utils"
	"encoding/json"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
)

/**
* GET
* @path /locatoins
* return [{"id":0,"username":"anu","password":"","acctstarttime":"2015-09-20 18:49:32",
*         "acctlastupdatedtime":"2015-09-20 18:49:32","acctactivationtime":"","acctstoptime":"2015-09-20 19:49:32"}]
*/
func GetLocations(w http.ResponseWriter, r *http.Request){
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()
	var locations []dao.Location

	_, err := dbMap.Select(&locations, "SELECT locationid, locationname, nasip, ipfrom, ipto FROM aplocation")
	checkErr(err, "Select failed")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(locations); err != nil {
		panic(err)
	}
}

/**
* POST
* @path /locations
* return [{"id":0,"username":"anu","password":"","acctstarttime":"2015-09-20 18:49:32",
*         "acctlastupdatedtime":"2015-09-20 18:49:32","acctactivationtime":"","acctstoptime":"2015-09-20 19:49:32"}]
*/
func AddLocation(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var location dao.Location
	err := decoder.Decode(&location)

	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare("INSERT INTO aplocation (locationid, locationname, nasip, ipfrom, ipto) VALUES( ?, ?, ?, ? , ? )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmtIns.Exec(location.LocationId, location.LocationName, location.NasIP ,location.IPFrom , location.IPTo )
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

/**
* DELETE
* @path /locations/{locationid}
* return [{"id":0,"username":"anu","password":"","acctstarttime":"2015-09-20 18:49:32",
*         "acctlastupdatedtime":"2015-09-20 18:49:32","acctactivationtime":"","acctstoptime":"2015-09-20 19:49:32"}]
*/
func DeleteLocation(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	locationid := vars["locationid"]
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()

	_, err := dbMap.Exec("DELETE FROM aplocation where locationid=?", locationid)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		log.Fatalln("Error while deleting user " + locationid +" from DB ", err)
		w.WriteHeader(http.StatusNotFound)
	}else{
		w.WriteHeader(http.StatusOK)
	}
}

/**
* DELETE
* @path /locations/{locationid}/{mac}
* return [{"id":0,"username":"anu","password":"","acctstarttime":"2015-09-20 18:49:32",
*         "acctlastupdatedtime":"2015-09-20 18:49:32","acctactivationtime":"","acctstoptime":"2015-09-20 19:49:32"}]
*/
func DeleteLocationAccessPoint(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	mac := vars["mac"]
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()

	_, err := dbMap.Exec("DELETE FROM location where apmacaddress=?", mac)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		log.Fatalln("Error while deleting user " + mac +" from DB ", err)
		w.WriteHeader(http.StatusNotFound)
	}else{
		w.WriteHeader(http.StatusOK)
	}
}