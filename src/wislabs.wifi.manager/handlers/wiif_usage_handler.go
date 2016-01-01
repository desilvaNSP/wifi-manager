package handlers

import (
	"wislabs.wifi.manager/controllers/wifi"
	"wislabs.wifi.manager/dao"
	"encoding/json"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)


/**
* POST /wifi/usage/dailytotaldownloads
 */
func GetAgregatedDownloadsFromToHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var constrains dao.Constrains
	decoder.Decode(&constrains)
	count := wifi.GetAggregatedDownloadsFromTo(constrains)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(count); err != nil {
		panic(err)
	}
}

/**
* POST /wifi/usage/dailytotaluploads
 */
func GetAgregatedUploadsFromToHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var constrains dao.Constrains
	decoder.Decode(&constrains)
	count := wifi.GetAggregatedUploadsFromTo(constrains)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(count); err != nil {
		panic(err)
	}
}


/**
* POST
* @path /wifi/usage/downloads
*
*/
func GetDownloadsFromToHandler(w http.ResponseWriter, r *http.Request) {
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
* @path /wifi/usage/uploads
*
*/
func GetUploadsFromToHandler(w http.ResponseWriter, r *http.Request) {
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

/**
* POST
* @path /wifi/sessions/avg
*
*/
func GetAvgSessoinTimeFromToHandler(w http.ResponseWriter, r *http.Request) {
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

/**
* POST
* @path /wifi/sessions/count
*
*/
func GetTotalSessionCountTimeFromToHandler(w http.ResponseWriter, r *http.Request) {
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
