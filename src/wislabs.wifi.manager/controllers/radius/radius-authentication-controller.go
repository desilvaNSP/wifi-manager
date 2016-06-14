package radius

import (
	"wislabs.wifi.manager/dao"
	"github.com/kirves/goradius"
	"wislabs.wifi.manager/utils"
	"wislabs.wifi.manager/commons"
	log "github.com/Sirupsen/logrus"
	"gopkg.in/gorp.v1"
	"database/sql"
	"strconv"
	"errors"
	"strings"
	"net"
	"bytes"
)

func TestAuthenticationOnUser(nasClientInfo dao.NasClientTestInfo) bool{
	auth := goradius.Authenticator(nasClientInfo.ServerIp , nasClientInfo.AuthPort , nasClientInfo.Secret)
	authenticateStatus, err := auth.Authenticate(nasClientInfo.UserName, nasClientInfo.Password, nasClientInfo.NASClientName)

	if err != nil {
		return authenticateStatus;
	}
	return authenticateStatus
}
func CreateRadiusServer(rServer dao.RadiusServer) error{
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.ADD_RADIUS_SERVER)
	defer stmtIns.Close()
	if err != nil {
		return errors.New("Error occourred while insert radius server into radiusservers table| Stack : " + err.Error() )
	}
	_, err = stmtIns.Exec(rServer.TenantId, rServer.DBHostName, rServer.DBHostIp, rServer.DBSchemaName, rServer.DBHostPort, rServer.DBUserName, rServer.DBPassword);
	return err
}

func GetAllRadiusDetails(tenantId int) ([]dao.RadiusServer, error) {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	var radiusConfigs []dao.RadiusServer
	_, err := dbMap.Select(&radiusConfigs, commons.GET_ALL_RADIUS_CONFIGS, tenantId)
	if err != nil{
		return nil,	errors.New("Error occourred while getting all radius server configs| Stack : " + err.Error() )
	}
	return radiusConfigs, err
}

func GetInstanceConfigsById(instanceId int, tenantId int) (dao.RadiusServer, error){
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var radiusServerConfigs dao.RadiusServer

	err := dbMap.SelectOne(&radiusServerConfigs, commons.GET_SERVERCONFIGS_BY_INSTANCEID, instanceId, tenantId)
	if err != nil {
		return radiusServerConfigs, errors.New("Error occourred while getting server configs by id | Stack : " + err.Error() )
	}
	return radiusServerConfigs, err
}

func CreateNASClientOnServer(radiusServerConfigs dao.RadiusServer, nasClientInfo dao.NasClient) error{
	dbMap, errDb := GetRadiusServerDBConnection(radiusServerConfigs);
	defer dbMap.Db.Close()
	if errDb != nil{
		return errors.New("Error occourred while connection to DB server |"+" DB Server IP :"+radiusServerConfigs.DBHostIp +":"+strconv.Itoa(radiusServerConfigs.DBHostPort)+" | Stack : " + errDb.Error() )
	}
	stmtIns, err := dbMap.Db.Prepare(commons.ADD_NAS_CLIENT)
	defer stmtIns.Close()
	if err != nil {
		return errors.New("Error occourred while insert radius server into radiusservers table| Stack : " + err.Error() )
	}
	_, err = stmtIns.Exec(nasClientInfo.NasName, nasClientInfo.ShortName, nasClientInfo.NasType, nasClientInfo.NasPorts, nasClientInfo.Secret, nasClientInfo.NasServer, nasClientInfo.Community, nasClientInfo.Description);
	return err
}


func GetRadiusServerClients(radiusServerConfig dao.RadiusServer) ([]dao.NasClient, error){
	dbMap, errDb := GetRadiusServerDBConnection(radiusServerConfig);
	defer dbMap.Db.Close()
	var nasClients []dao.NasClient
	if errDb != nil{
		return nil, errors.New("Error occourred while connection to DB server |"+" DB Server IP :"+radiusServerConfig.DBHostIp +":"+strconv.Itoa(radiusServerConfig.DBHostPort)+" | Stack : " + errDb.Error() )

	}
	_, err := dbMap.Select(&nasClients, commons.GET_NAS_CLIENTS_INSERVER)
	if err != nil {
		return nil, errors.New("Error occourred while getting NAS clients On "+"DB Server IP :"+radiusServerConfig.DBHostIp +":"+strconv.Itoa(radiusServerConfig.DBHostPort)+" | Stack : " + err.Error() )
	}
	return nasClients, err
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

func UpdateRadiusServerInstance(config dao.RadiusServer, tenantId int) error{
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	_, err := dbMap.Exec(commons.UPDATE_RADIUS_SERVER_INST, config.DBHostName, config.DBHostPort, config.DBSchemaName, config.DBUserName, config.DBPassword, tenantId, config.InstanceId)
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

func IsNASIpExistsInRadius(radiusServerConfigs dao.RadiusServer, ipAddress string, length string) (bool, error)  {
	dbMap, errDb := GetRadiusServerDBConnection(radiusServerConfigs);
	defer dbMap.Db.Close()
	if errDb != nil{
		return false, errors.New("Error occourred while connection to DB server |"+" DB Server IP :"+radiusServerConfigs.DBHostIp +":"+strconv.Itoa(radiusServerConfigs.DBHostPort)+" | Stack : " + errDb.Error() )
	}
	rangeSize, _ := strconv.Atoi(length)
	var allNasNames []string
	_, err := dbMap.Select(&allNasNames, commons.GET_ALL_NASNAMES)
	if err != nil {
		return false, err
	}
	for _, nasName := range allNasNames {
		result, err := checkIpBetweenIPRange(ipAddress,rangeSize,nasName)
		if err != nil {
		}
		if(result){
			return true, nil
		}
	}
	return false, nil
}

// Get Database connection for each radius server.That database connection for get NAS clients of each radius server
func GetRadiusServerDBConnection(radiusServerConfig dao.RadiusServer) (*gorp.DbMap, error){
	var connectionUrl string
	connectionUrl = radiusServerConfig.DBUserName + ":" + radiusServerConfig.DBPassword + "@tcp(" + radiusServerConfig.DBHostIp + ":" + strconv.Itoa(radiusServerConfig.DBHostPort)  + ")/" + radiusServerConfig.DBSchemaName
	db, err := sql.Open("mysql", connectionUrl)
	if err != nil {
		return nil, errors.New("Error occourred while conecting to radius server ip:"+ radiusServerConfig.DBHostIp + "stack : " + err.Error() )
	}
	dbmap := &gorp.DbMap{Db: db, Dialect:gorp.MySQLDialect{"InnoDB", "UTF8"}}
	return dbmap, err
}


func checkIpBetweenIPRange(ip string,rangeSize int, iprange string) (bool, error){
	var slices []string
	slices = strings.Split(iprange, "/")
	var endLessDBIp string
	var err error
	if len(slices)>1 {
		netMask, _ := strconv.Atoi(slices[1])
		endLessDBIp, err = getEndPointOfIpRange(slices[0],netMask);
		if err != nil{
			return false, err
		}
		var (
			ip1 = net.ParseIP(slices[0])
			ip2 = net.ParseIP(endLessDBIp)
		)
		trial := net.ParseIP(ip)
		if trial.To4() == nil {
			return false, nil
		}
		if rangeSize > 0{
			var endLessUserEnterIp string
			endLessUserEnterIp, err = getEndPointOfIpRange(ip,rangeSize);
			if err != nil{
				return false, err
			}
			var (
				ip3 = net.ParseIP(ip)
				ip4 = net.ParseIP(endLessUserEnterIp)
			)
			if ( bytes.Compare(ip1, ip4) > 0)  || (bytes.Compare(ip2, ip3) < 0 ){
				return false, nil
			}
			return true, nil
		}else{
			if bytes.Compare(trial, ip1) >= 0 && bytes.Compare(trial, ip2) <= 0 {
				return true, nil
			}
			return false, nil
		}
	}else {
		var ip1 = net.ParseIP(slices[0])
		trial := net.ParseIP(ip)
		if trial.To4() == nil {
			return false, nil
		}
		if rangeSize > 0{
			var endLessUserEnterIp string
			endLessUserEnterIp, err = getEndPointOfIpRange(ip,rangeSize);
			if err != nil{
				return false, err
			}
			var (
				ip3 = net.ParseIP(ip)
				ip4 = net.ParseIP(endLessUserEnterIp)
			)
			if ( bytes.Compare(ip1, ip3) >= 0)  && (bytes.Compare(ip4, ip1) >= 0 ){
				return true, nil
			}
			return false, nil
		}else if bytes.Compare(trial, ip1) == 0 {
				return true, nil
		}
		return false, nil
	}
}


func getEndPointOfIpRange(ip string, netMask int) (string, error) {
	ipaddress := net.ParseIP(ip)
	ipaddress = ipaddress.To4()
	if ipaddress == nil {
		log.Error("non ipv4 address")
		return "", errors.New("Non ipv4 address")
	}
	ipaddress[3] += byte(netMask)
	return ipaddress.String(), nil
}