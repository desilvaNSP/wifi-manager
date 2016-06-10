package radius

import (
	"wislabs.wifi.manager/dao"
	"github.com/kirves/goradius"
	"wislabs.wifi.manager/utils"
	"wislabs.wifi.manager/commons"
	log "github.com/Sirupsen/logrus"

)

func TestAuthenticationOnUser(radiusConfig dao.RadiusConfigsInfo) bool{
	auth := goradius.Authenticator(radiusConfig.ServerIP, radiusConfig.AuthPort, radiusConfig.SharedSecret)

	authenticateStatus, err := auth.Authenticate(radiusConfig.TestUsername, radiusConfig.Password,"test")
	if err != nil {
		return authenticateStatus;
	}
	return authenticateStatus
}

func CreateRadiusServer(rConfig dao.RadiusConfigsInfo) error{
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.ADD_RADIUS_SERVER)
	defer stmtIns.Close()
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(rConfig.TenantId, rConfig.Username,rConfig.ServerName,rConfig.ServerIP,rConfig.AuthPort,rConfig.Accounting,rConfig.SharedSecret);

	return err
}

func GetAllRadiusDetails(tenantid int,username string) []dao.RadiusConfigs  {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var radiusConfigs []dao.RadiusConfigs
	_, err := dbMap.Select(&radiusConfigs, commons.GET_ALL_RADIUS_CONFIGS, tenantid,username)
	checkErr(err, "Error occured while getting Radius Details")
	return radiusConfigs
}

func DeleteRadiusInstance(tenantid int,username string) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	_, err := dbMap.Exec(commons.DELETE_RADIUS_SERVER_INST, tenantid, username)
	if (err != nil) {
		return err
	}else {
		return nil
	}
}

func UpdateRadiusInstance(config dao.RadiusConfigs,tenantid int,username string) error{
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	_, err := dbMap.Exec(commons.UPDATE_RADIUS_SERVER_INST, config.ServerName, config.ServerIP, config.AuthPort, config.Accounting, config.SharedSecret, tenantid, username)
	if (err != nil) {
		return err
	}else {
		return nil
	}
}

func IsWifiUserValidInRadius(tenantId int, username string) (int, error) {
	dbMap := utils.GetDBConnection(commons.RADIUS_DB);
	defer dbMap.Db.Close()

	var checkUser int
	err := dbMap.SelectOne(&checkUser, commons.IS_VALID_USER_IN_RADIUS, username, username)
	if err != nil {
		return checkUser, err
	}
	return checkUser, nil
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Error(msg, err)
	}
}