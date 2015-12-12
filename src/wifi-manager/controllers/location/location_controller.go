package location

import (
	"wifi-manager/utils"
	log "github.com/Sirupsen/logrus"
	"wifi-manager/dao"
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

//func IsOverlappingRange(from string, to string) string{
//	dbMap := utils.GetDBConnection("portal");
//	defer dbMap.Db.Close()
//	var locationId string
//	var query string
//
//	query = "SELECT locationid FROM aplocation WHERE INET_ATON(ipfrom)<= INET_ATON('" + ip + "') AND INET_ATON(ipto) >=INET_ATON('" + ip + "');"
//
//	locationId, err := dbMap.SelectStr(query)
//	checkErr(err, "Select failed")
//	return locationId
//}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
