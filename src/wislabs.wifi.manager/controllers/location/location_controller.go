package location

import (
	"wislabs.wifi.manager/utils"
	log "github.com/Sirupsen/logrus"
	"wislabs.wifi.manager/dao"
	"wislabs.wifi.manager/common"
)

func GetAllLocations(tenantid int) []dao.ApLocation {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	var apLocations []dao.ApLocation
	_, err := dbMap.Select(&apLocations, common.GET_ALL_AP_LOCATIONS, tenantid)
	checkErr(err, "Error occured while getting AP locations")
	return apLocations
}

func AddLocation(location *dao.ApLocation) {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns , err := dbMap.Db.Prepare(common.ADD_AP_LOCATION)
	_, err = stmtIns.Exec(location.TenantId, location.SSID, location.MAC, location.Longitude, location.Latitude, location.GroupName)
	checkErr(err, "Error occured while adding AP location")
	defer stmtIns.Close()
}

func DeleteApLocation(tenantid int, ssid string, mac string, groupName string) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	_, err := dbMap.Exec(common.DELETE_AP_LOCATION, ssid, mac, groupName, tenantid)
	if(err != nil){
		return err
	}else {
		return nil
	}
}

func DeleteApGroup(groupName string,tenantid int) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	_, err := dbMap.Exec(common.DELETE_AP_GROUP, groupName, tenantid)
	if(err != nil){
		return err
	}else {
		return nil
	}
}

func DeleteAccessPoint(mac string, tenantid int) error{
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	_, err := dbMap.Exec(common.DELETE_AP, mac, tenantid)
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

