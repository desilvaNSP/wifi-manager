package handlers

import (
	"wislabs.wifi.manager/controllers/wifi"
	"wislabs.wifi.manager/dao"
	"wislabs.wifi.manager/authenticator"
	"encoding/json"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"os"
	"bytes"
	"encoding/csv"
	"log"
	"wislabs.wifi.manager/commons"
	"math"
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

	count, countpre := wifi.GetDownloadsFromTo(constrains)
	changePercentage := getChangePrecentageSummaryDetails(countpre,count)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(changePercentage); err != nil {
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

	count, countpre := wifi.GetUploadsFromTo(constrains)
	changePercentage := getChangePrecentageSummaryDetails(countpre,count)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(changePercentage); err != nil {
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

	count, countpre := wifi.GetAvgSessionsFromTo(constrains)
	changePercentage := getChangePrecentageSummaryDetails(int64(countpre),int64(count))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(changePercentage); err != nil {
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

	count, countpre := wifi.GetTotalSessionsCountFromTo(constrains)
	changePercentage := getChangePrecentageSummaryDetails(countpre,count)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(changePercentage); err != nil {
		panic(err)
	}
}

/**
* POST
* @path /wifi/summary/accespoint
*
*/

func GetAccessPointAggregatedDataFromToHandler(w http.ResponseWriter, r*http.Request){
	decoder := json.NewDecoder(r.Body)
	var constrains dao.Constrains
	decoder.Decode(&constrains)

	var accesPoint[] dao.AccessPoint
	accesPoint = wifi.GetAccessPointAggregatedDataFromTo(constrains)

	accessPointDataWithLocation := make([]dao.LocationAccessPoint, len(accesPoint))

	for index, point := range accesPoint {
		var longLatPoint dao.LongLatMac
		longLatPoint =  wifi.GetLongLatLocationByMacAddress(point.Calledstationmac.String)
		point.APName = longLatPoint.APName
		accessPointDataWithLocation[index].AccessPointData = point;
		accessPointDataWithLocation[index].LongLatMacData = longLatPoint
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(accessPointDataWithLocation); err != nil {
		panic(err)
	}
}


/**
* POST
* @path /wifi/summary/downloadrawdata
*
*/

func DownlaodCSVSummaryDetailsDashboard(w http.ResponseWriter,r *http.Request){
	if(!authenticator.IsAuthorized(authenticator.CSV_DOWNLOAD, authenticator.ACTION_READ,r)){
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var constrains dao.Constrains
	decoder.Decode(&constrains)
	//generate temp file
	CSVcontent := wifi.SummaryDetailsFromTo(constrains)
	serverHome := os.Getenv(commons.SERVER_HOME)
	tempfile, err := ioutil.TempFile(serverHome + "/tmp/","summary")
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


func getChangePrecentageSummaryDetails(countpre int64, count int64) dao.SummaryChangePercentage{
	var changePercentage dao.SummaryChangePercentage
	changePercentage.Value = count
	changePercentage.PreValue = countpre
	if countpre == 0 {
		changePercentage.Percentage = 100;
	}else{
		changePercentage.Percentage =  math.Abs(float64(countpre-count)/float64(countpre)*100)

	}
	if countpre < count {
		changePercentage.Status = "up";
	}else {
		changePercentage.Status = "down"
	}
	return  changePercentage
}
