package handlers

import (
	"wislabs.wifi.manager/controllers/wifi"
	"wislabs.wifi.manager/dao"
	"encoding/json"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)



func GetAgregatedDownloadsFromToHandler(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var constrains dao.Constrains
	decoder.Decode(&constrains)
	count := wifi.GetAgregatedDownloadsFromTo(constrains)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(count); err != nil {
		panic(err)
	}
}

/**
* GET
* @path /locatoins
* return [{"id":0,"username":"anu","password":"","acctstarttime":"2015-09-20 18:49:32",
*         "acctlastupdatedtime":"2015-09-20 18:49:32","acctactivationtime":"","acctstoptime":"2015-09-20 19:49:32"}]
*/
func GetDownloadsFromToHandler(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var constrains dao.Constrains
	decoder.Decode(&constrains)
	count := wifi.GetDownloadsFromTo(constrains)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(count); err != nil {
		panic(err)
	}
}

/**
* POST
* @path /locations
* return [{"id":0,"username":"anu","password":"","acctstarttime":"2015-09-20 18:49:32",
*         "acctlastupdatedtime":"2015-09-20 18:49:32","acctactivationtime":"","acctstoptime":"2015-09-20 19:49:32"}]
*/
func GetUploadsFromToHandler(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var constrains dao.Constrains
	decoder.Decode(&constrains)

	count := wifi.GetUploadsFromTo(constrains)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(count); err != nil {
		panic(err)
	}
}

func GetAvgSessoinTimeFromToHandler(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var constrains dao.Constrains
	decoder.Decode(&constrains)

	count := wifi.GetAvgSessionsFromTo(constrains)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(count); err != nil {
		panic(err)
	}
}

func GetTotalSessoinCountTimeFromToHandler(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var constrains dao.Constrains
	decoder.Decode(&constrains)

	count := wifi.GetTotalSessionsCountFromTo(constrains)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(count); err != nil {
		panic(err)
	}
}
