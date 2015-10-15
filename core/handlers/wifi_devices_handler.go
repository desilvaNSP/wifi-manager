package handlers

import (
	"wifi-manager/core/dao"
	"wifi-manager/core/controllers/wifi"
	"encoding/json"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

/**
* GET
* @path /locatoins
* return [{"id":0,"username":"anu","password":"","acctstarttime":"2015-09-20 18:49:32",
*         "acctlastupdatedtime":"2015-09-20 18:49:32","acctactivationtime":"","acctstoptime":"2015-09-20 19:49:32"}]
*/
func GetUsersByOSHandler(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var constrains dao.Constrains
	decoder.Decode(&constrains)

	usersByOS := wifi.GetUsersByOS(constrains)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(usersByOS); err != nil {
		panic(err)
	}
}

func GetUsersByDeviceTypeHandler(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var constrains dao.Constrains
	decoder.Decode(&constrains)

	usersByDevice := wifi.GetUsersByDevice(constrains)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(usersByDevice); err != nil {
		panic(err)
	}
}
