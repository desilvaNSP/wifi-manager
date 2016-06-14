package location

import (
	"wislabs.wifi.manager/utils"
	log "github.com/Sirupsen/logrus"
	"wislabs.wifi.manager/dao"
	"wislabs.wifi.manager/commons"
	"wislabs.wifi.manager/controllers/dashboard"
	"errors"
)

func GetAllLocations(tenantid int) []dao.ApLocation {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	var apLocations []dao.ApLocation
	_, err := dbMap.Select(&apLocations, commons.GET_ALL_AP_LOCATIONS, tenantid)
	checkErr(err, "Error occured while getting AP locations")
	return apLocations
}

func GetAllLocationGroups(tenantid int) []string {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	var apLocationGroups []string
	_, err := dbMap.Select(&apLocationGroups, commons.GET_ALL_AP_GROUPS, tenantid)
	checkErr(err, "Error occured while getting AP location groups")
	return apLocationGroups
}

func AddWiFiLocation(location *dao.ApLocation) {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.ADD_AP_LOCATION)
	_, err = stmtIns.Exec(location.TenantId, location.SSID, location.MAC, location.APName, location.BSSID, location.Longitude, location.Latitude, dashboard.GetApGroupId(location.TenantId, location.GroupName), location.GroupName)
	checkErr(err, "Error occured while adding AP location")
	defer stmtIns.Close()
}

func UpdateWifiLocation(location *dao.ApLocation){
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	stmtIns, err := dbMap.Db.Prepare(commons.UPDATE_AP_LOCATION)
	_, err = stmtIns.Exec(location.SSID, location.APName, location.BSSID, location.Longitude, location.Latitude, dashboard.GetApGroupId(location.TenantId, location.GroupName), location.GroupName, location.LocationId,location.TenantId)
	checkErr(err, "Error occured while updating AP location")
	defer stmtIns.Close()
}

func AddWiFiGroup(group *dao.ApGroup) {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.ADD_AP_GROUP)
	_, err = stmtIns.Exec(group.TenantId, group.GroupName, group.GroupSymbol)
	checkErr(err, "Error occured while adding AP group")
	defer stmtIns.Close()
}

func DeleteApLocation(tenantid int, ssid string, mac string, groupName string) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	_, err := dbMap.Exec(commons.DELETE_AP_LOCATION, ssid, mac, groupName, tenantid)
	if (err != nil) {
		return err
	}else {
		return nil
	}
}

func DeleteApGroup(groupName string, tenantid int) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	_, err := dbMap.Exec(commons.DELETE_AP_GROUP, groupName, tenantid)
	if (err != nil) {
		return err
	}else {
		return nil
	}
}

func DeleteAccessPoint(mac string, tenantid int) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	_, err := dbMap.Exec(commons.DELETE_AP, mac, tenantid)
	if (err != nil) {
		return err
	}else {
		return nil
	}
}

func GetActiveInactiveAccessPoint(tenantId int,activePeriodFrom string, activePeriodTo string, treshold int) (int, error){
	dbMap := utils.GetDBConnection(commons.SUMMARY_DB_NAME);
	defer dbMap.Db.Close()
	var activeCalledStations []string

	_, err := dbMap.Select(&activeCalledStations, commons.GET_ACTIVE_APS_COUNT,"2016-01-10 00:00:00", "2016-01-25 00:00:00", tenantId,  treshold)
	if err != nil {
		return len(activeCalledStations), errors.New(err.Error())
	}
	return len(activeCalledStations), nil
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Error(msg, err)
	}
}

