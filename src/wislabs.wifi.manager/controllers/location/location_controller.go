package location

import (
	"wislabs.wifi.manager/utils"
	log "github.com/Sirupsen/logrus"
	"wislabs.wifi.manager/dao"
	"wislabs.wifi.manager/common"
)

func GetAllLocations() []dao.ApLocation {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	var aplocations []dao.ApLocation
	_, err := dbMap.Select(&aplocations, common.GET_APLOCATIONS)
	checkErr(err, "Error occured while getting AP locations")
	return aplocations
}

func AddLocation(location *dao.ApLocation) {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns , err := dbMap.Db.Prepare(common.ADD_APLOCATION)
	_, err = stmtIns.Exec(location.SSID, location.MAC, location.Longitude, location.Latitude)
	checkErr(err, "Error occured while adding AP location")
	defer stmtIns.Close()
}

func DeleteApLocation(ssid string, mac string) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	_, err := dbMap.Exec(common.DELETE_APLOCATION, ssid, mac)
	if(err != nil){
		return err
	}else {
		return nil
	}
}

func DeleteAccessPoint(mac string) error{
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	_, err := dbMap.Exec(common.DELETE_AP, mac)
	if(err != nil){
		return err
	}else {
		return nil
	}
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

