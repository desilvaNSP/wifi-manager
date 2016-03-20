package handlers

import (
	"wislabs.wifi.manager/controllers/wifi"
	"wislabs.wifi.manager/dao"
	"encoding/json"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"os"
	"bytes"
	"encoding/csv"
	"log"
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
* POST /wifi/usage/dailyavguserdownloads
 */
func GetAvgUserDownloadsFromToHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var constrains dao.Constrains
	decoder.Decode(&constrains)
	count := wifi.GetAvgDailyDownloadsPerUserFromTo(constrains)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(count); err != nil {
		panic(err)
	}
}

/**
* POST /wifi/usage/dailyavgusersessiontime
 */
func GetAvgUserSessionTimeFromToHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var constrains dao.Constrains
	decoder.Decode(&constrains)
	count := wifi.GetAvgDailySessionTimePerUserFromTo(constrains)

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

/**
* POST
* @path /wifi/summary/downloadrawdata
*
*/

func DownlaodCSVSummaryDetailsDashboard(w http.ResponseWriter,r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var constrains dao.Constrains
	decoder.Decode(&constrains)

	//generate temp file
	CSVcontent := wifi.SummaryDetailsFromTo(constrains)
	tempfile, err := ioutil.TempFile("/","summary")
	if err != nil {
		log.Fatal("Cannot create temp file ", err)
	}
	defer os.Remove(tempfile.Name())

	writer := csv.NewWriter(tempfile)

	for _, value := range CSVcontent {
		err := writer.Write(value)
		if err != nil {
			log.Fatal("Cannot write to file ", err)
		}
	}
	writer.Flush()

	if err := tempfile.Close(); err != nil {
		log.Fatal(err)
	}
	//read content in temp file then send to client side
	streamCSVbytes, err := ioutil.ReadFile(tempfile.Name());

	if err != nil {
		os.Exit(1)
	}
	buffer := bytes.NewBuffer(streamCSVbytes)

	w.Header().Set("Content-type", "application/csv")
	if _, err := buffer.WriteTo(w); err != nil {
		panic(err)
	}

}
