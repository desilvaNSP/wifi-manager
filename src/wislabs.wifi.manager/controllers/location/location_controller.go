package location

import (
	"wislabs.wifi.manager/utils"
	log "github.com/Sirupsen/logrus"
	"wislabs.wifi.manager/dao"
)

func GetLocationFromIP(ip string) string{
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()
	var locationId string
	var query string

	query = "SELECT locationid FROM aplocation WHERE INET_ATON(ipfrom)<= INET_ATON('" + ip + "') AND INET_ATON(ipto) >=INET_ATON('" + ip + "');"

	locationId, err := dbMap.SelectStr(query)
	checkErr(err, "Select failed")
	return locationId
}

func ReCalculateLocationFromIP(constrains dao.Constrains) string{
	dbMap := utils.GetDBConnection("radsummary");
	defer dbMap.Db.Close()
	var locationId string
	var query string
 //UPDATE radsummary.dailyacct SET location = (SELECT locationid FROM portal.aplocation WHERE INET_ATON(portal.aplocation.ipfrom)<= INET_ATON(framedipaddress) AND INET_ATON(portal.aplocation.ipto) >=INET_ATON(framedipaddress));
	query = "UPDATE dailyacct SET location = (SELECT locationid FROM portal.aplocation WHERE INET_ATON(portal.aplocation.ipfrom)<= INET_ATON(framedipaddress) AND INET_ATON(portal.aplocation.ipto) >=INET_ATON(framedipaddress) )"

	locationId, err := dbMap.SelectStr(query)
	checkErr(err, "Select failed")
	return locationId
}

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

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
