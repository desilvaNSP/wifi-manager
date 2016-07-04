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
	dbMap := utils.GetDBConnection(commons.DASHBOARD_DB);
	defer dbMap.Db.Close()
	var apLocations []dao.ApLocation
	_, err := dbMap.Select(&apLocations, commons.GET_ALL_AP_LOCATIONS, tenantid)
	checkErr(err, "Error occured while getting AP locations")
	return apLocations
}

func GetAllAPs(tenantid int) ([]dao.APs, error) {
	dbMap := utils.GetDBConnection(commons.DASHBOARD_DB);
	defer dbMap.Db.Close()
	var aps []dao.APs
	_, err := dbMap.Select(&aps, commons.GET_ALL_APS, tenantid)
	if err != nil{
		return nil, err
	}
	return aps, nil
}

func GetAllLocationGroups(tenantid int) []string {
	dbMap := utils.GetDBConnection(commons.DASHBOARD_DB);
	defer dbMap.Db.Close()
	var apLocationGroups []string
	_, err := dbMap.Select(&apLocationGroups, commons.GET_ALL_AP_GROUPS, tenantid)
	checkErr(err, "Error occured while getting AP location groups")
	return apLocationGroups
}

func GetAPsOnLocation(mac string, tenantid int) (dao.APs , error){
	dbMap := utils.GetDBConnection(commons.DASHBOARD_DB);
	defer dbMap.Db.Close()
	var apsOnLocation dao.APs
	err := dbMap.SelectOne(&apsOnLocation, commons.GET_APS_ON_LOCATION, mac, tenantid)
	if err != nil{
		return  apsOnLocation , err
	}
	return apsOnLocation, nil
}

func GetMACsOnAllLocations(tenantid int) ([]string , error){
	dbMap := utils.GetDBConnection(commons.DASHBOARD_DB);
	defer dbMap.Db.Close()
	var allMACS []string
	_, err := dbMap.Select(&allMACS, commons.GET_ALL_APS_MACS, tenantid)
	if err != nil {
		return nil, err
	}
	return allMACS, nil
}


func GetSSIDsOfLocationGroups(groupnames []string, tenantId int) []string {
	dbMap := utils.GetDBConnection(commons.DASHBOARD_DB);
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

func IsSSIDExistsOnMac(macaddress string, ssid string, tenantId int) (int ,error) {
	dbMap := utils.GetDBConnection(commons.DASHBOARD_DB);
	defer dbMap.Db.Close()
	var valid int
	err := dbMap.SelectOne(&valid, commons.IS_EXISTS_SSID_ON_MAC, macaddress, ssid, tenantId)
	if err != nil {
		return valid, err
	}
	return valid, nil
}

func AddWiFiLocation(location *dao.ApLocationSSIDs) error {
	dbMap := utils.GetDBConnection(commons.DASHBOARD_DB);
	defer dbMap.Db.Close()
	stmtInsAps, err  := dbMap.Db.Prepare(commons.ADD_APS)
	defer stmtInsAps.Close()
	stmtInsApLocation, err  := dbMap.Db.Prepare(commons.ADD_AP_LOCATION)
	defer stmtInsApLocation.Close()
	if err != nil {
		return err
	}

	if  IsMacExists(location.TenantId, location.MAC) == 0 {
		_, err = stmtInsAps.Exec(location.TenantId, location.APs.APName, location.APs.Address, location.APs.Longitude, location.APs.Latitude, location.MAC)
		if err != nil {
			return  err
		}
	}
	for i := 0; i < len(location.SSID); i++ {
		_, err = stmtInsApLocation.Exec(location.TenantId, location.SSID[i], location.MAC, location.BSSID, dashboard.GetApGroupId(location.TenantId, location.GroupName), location.GroupName)
		if err != nil {
			return  err
		}
	}
	return nil
}

func UpdateWifiLocation(location *dao.ApLocationSSIDs) error{
	dbMap := utils.GetDBConnection(commons.DASHBOARD_DB);
	defer dbMap.Db.Close()
	stmtInsApLocation, err := dbMap.Db.Prepare(commons.UPDATE_AP_LOCATION)
	defer stmtInsApLocation.Close()
	if err != nil {
		return err
	}
	_, err = stmtInsApLocation.Exec(location.BSSID, location.SSID[0], dashboard.GetApGroupId(location.TenantId, location.GroupName), location.GroupName, location.MAC, location.SSID[1], location.TenantId)
	if err != nil {
		return  err
	}
	return nil
}

func UpdateAPs(aps *dao.APs) error{
	dbMap := utils.GetDBConnection(commons.DASHBOARD_DB);
	defer dbMap.Db.Close()
	stmtInsAps, err := dbMap.Db.Prepare(commons.UPDATE_APS)
	defer stmtInsAps.Close()
	if err != nil {
		return err
	}
	_, err = stmtInsAps.Exec(aps.APName, aps.Address, aps.Longitude, aps.Latitude, aps.MAC, aps.TenantId)
	if err != nil {
		return  err
	}
	return nil
}


func AddWiFiGroup(group *dao.ApGroup) {
	dbMap := utils.GetDBConnection(commons.DASHBOARD_DB);
	defer dbMap.Db.Close()
	stmtIns, err := dbMap.Db.Prepare(commons.ADD_AP_GROUP)
	_, err = stmtIns.Exec(group.TenantId, group.GroupName, group.GroupSymbol)
	checkErr(err, "Error occured while adding AP group")
	defer stmtIns.Close()
}

func DeleteApLocation(tenantid int, ssid string, mac string, groupName string) error {
	dbMap := utils.GetDBConnection(commons.DASHBOARD_DB);
	defer dbMap.Db.Close()
	_, err := dbMap.Exec(commons.DELETE_AP_LOCATION, ssid, mac, groupName, tenantid)
	if (err != nil) {
		return err
	} else {
		return nil
	}
}

func DeleteApGroup(groupName string, tenantid int) error {
	dbMap := utils.GetDBConnection(commons.DASHBOARD_DB);
	defer dbMap.Db.Close()
	_, err := dbMap.Exec(commons.DELETE_AP_GROUP, groupName, tenantid)
	if (err != nil) {
		return err
	} else {
		return nil
	}
}

func DeleteAccessPoint(mac string, tenantid int) error {
	dbMap := utils.GetDBConnection(commons.DASHBOARD_DB);
	defer dbMap.Db.Close()
	_, err := dbMap.Exec(commons.DELETE_AP, mac, tenantid)
	if (err != nil) {
		return err
	} else {
		return nil
	}
}

func GetAccessPointFeatureDetails(constraints dao.AccessPointConstraints) (int, error){
	dbMap := utils.GetDBConnection(commons.SUMMARY_DB);
	defer dbMap.Db.Close()
	var selectedResult []utils.NullString
	var err error;
	if constraints.Threshold != 0 {
		_, err = dbMap.Select(&selectedResult,constraints.Query, constraints.From, constraints.To, constraints.TenantId, constraints.Threshold)
	}else {
		_, err = dbMap.Select(&selectedResult,constraints.Query, constraints.From, constraints.To, constraints.TenantId)
	}
	if err != nil {
		return len(selectedResult), errors.New(err.Error())
	}
	return len(selectedResult), nil
}

func IsMacExists(tenantId int, macAddress string) int{
	dbMap := utils.GetDBConnection(commons.DASHBOARD_DB);
	defer dbMap.Db.Close()
	var checkmac int
	err := dbMap.SelectOne(&checkmac, commons.IS_MAC_EXISTS, macAddress, tenantId)
	if err != nil {
		checkErr(err, "Error occur while checking exists user");
	}
	return checkmac
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Error(msg, err)
	}
}

