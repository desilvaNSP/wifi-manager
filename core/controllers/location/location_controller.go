package location

import (
	"dashboard-core/utils"
	log "github.com/Sirupsen/logrus"
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

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
