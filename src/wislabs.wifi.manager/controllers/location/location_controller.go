package location

import (
	"wislabs.wifi.manager/utils"
	log "github.com/Sirupsen/logrus"
	"wislabs.wifi.manager/dao"
	"wislabs.wifi.manager/commons"
	"wislabs.wifi.manager/controllers/dashboard"
	"strings"
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

func GetSSIDsOfLocationGroups(groupnames []string, tenantId int) []string {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	var ssids []string
	query := commons.GET_AP_GROUP_SSIDS + "("

	for index, value := range groupnames {
		aa := strings.Replace(value, "\"", "", -1)

		query += "'" + strings.Trim(aa, " ") + "'"
		if index < len(groupnames) - 1 {
			query += ","
		}
	}
	_, err := dbMap.Select(&ssids, query + ")", tenantId)
	checkErr(err, "Error occured while getting AP location groups")
	return ssids
}

func AddWiFiLocation(location *dao.ApLocation) {
	dbMap := utils.GetDBConnection(commons.DASHBOARD_DB);
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.ADD_AP_LOCATION)
	_, err = stmtIns.Exec(location.TenantId, location.SSID, location.MAC, location.APName, location.BSSID,  location.Address, location.Longitude, location.Latitude, dashboard.GetApGroupId(location.TenantId, location.GroupName), location.GroupName)
	if err != nil {
		println(err.Error())
	}
	checkErr(err, "Error occured while adding AP location")
	defer stmtIns.Close()
}

func UpdateWifiLocation(location *dao.ApLocation) {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	stmtIns, err := dbMap.Db.Prepare(commons.UPDATE_AP_LOCATION)
	_, err = stmtIns.Exec(location.SSID, location.APName, location.BSSID, location.Address, location.Longitude, location.Latitude, dashboard.GetApGroupId(location.TenantId, location.GroupName), location.GroupName, location.LocationId, location.TenantId)
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
	} else {
		return nil
	}
}

func DeleteApGroup(groupName string, tenantid int) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	_, err := dbMap.Exec(commons.DELETE_AP_GROUP, groupName, tenantid)
	if (err != nil) {
		return err
	} else {
		return nil
	}
}

func DeleteAccessPoint(mac string, tenantid int) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	_, err := dbMap.Exec(commons.DELETE_AP, mac, tenantid)
	if (err != nil) {
		return err
	} else {
		return nil
	}
}

func GetActiveAccessPoint(tenantId int, activePeriodFrom string, activePeriodTo string, treshold int) (int, error) {
	dbMap := utils.GetDBConnection(commons.SUMMARY_DB);
	defer dbMap.Db.Close()
	var activeCalledStations []utils.NullString
	_, err := dbMap.Select(&activeCalledStations, commons.GET_ACTIVE_APS_COUNT, activePeriodFrom, activePeriodTo, tenantId, treshold)
	if err != nil {
		return len(activeCalledStations), errors.New(err.Error())
	}
	return len(activeCalledStations), nil
}

func GetInactiveAccessPoint(tenantId int, activePeriodFrom string, activePeriodTo string, treshold int) (int, error) {
	dbMap := utils.GetDBConnection(commons.SUMMARY_DB);
	defer dbMap.Db.Close()
	var inactiveCalledStations []utils.NullString
	_, err := dbMap.Select(&inactiveCalledStations, commons.GET_INACTIVE_APS_COUNT, activePeriodFrom, activePeriodTo, tenantId, treshold)
	if err != nil {
		return len(inactiveCalledStations), errors.New(err.Error())
	}
	return len(inactiveCalledStations), nil
}

func GetDistinctMacCount(tenantId int, activePeriodFrom string, activePeriodTo string) (int, error) {
	dbMap := utils.GetDBConnection(commons.SUMMARY_DB);
	defer dbMap.Db.Close()
	var activeCalledStations []utils.NullString
	_, err := dbMap.Select(&activeCalledStations, commons.GET_DISTINCT_MAC, activePeriodFrom, activePeriodTo, tenantId)
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

