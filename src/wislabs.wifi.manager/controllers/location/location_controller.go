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
	var apLocations []dao.ApLocation
	_, err := dbMap.Select(&apLocations, common.GET_ALL_AP_LOCATIONS)
	checkErr(err, "Error occured while getting AP locations")
	return apLocations
}

func AddLocation(location *dao.ApLocation) {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns , err := dbMap.Db.Prepare(common.ADD_AP_LOCATION)
	_, err = stmtIns.Exec(location.SSID, location.MAC, location.Longitude, location.Latitude, location.GroupName)
	checkErr(err, "Error occured while adding AP location")
	defer stmtIns.Close()
}

func DeleteApLocation(ssid string, mac string, groupName string) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	_, err := dbMap.Exec(common.DELETE_AP_LOCATION, ssid, mac, groupName)
	if(err != nil){
		return err
	}else {
		return nil
	}
}

func DeleteApGroup(groupName string) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	_, err := dbMap.Exec(common.DELETE_AP_GROUP, groupName)
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

