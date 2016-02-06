package handlers

import (
	"wislabs.wifi.manager/dao"
	"wislabs.wifi.manager/controllers/wifi"
	"encoding/json"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

func GetBrowserStatsHandler2(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var constrains dao.Constrains
	decoder.Decode(&constrains)

	usersByOS := wifi.GetBrowserStats(constrains)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(usersByOS); err != nil {
		panic(err)
	}
}
