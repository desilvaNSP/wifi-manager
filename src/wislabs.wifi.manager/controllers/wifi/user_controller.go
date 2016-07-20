package wifi

import (
	"wislabs.wifi.manager/utils"
	"wislabs.wifi.manager/dao"
	"wislabs.wifi.manager/commons"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strconv"
)

func GetDailyUserCountSeriesFromTo(constrains dao.Constrains) [] dao.NameValue {
	dbMap := utils.GetDBConnection("summary");
	defer dbMap.Db.Close()
	var totalDailyDownloads[] dao.NameValue
	query := "SELECT COUNT(DISTINCT username) as value ,date as name FROM dailyacct where date >= ? AND date <= ? AND tenantid=? "
	if len(constrains.ACL) > 0 {
		query = query + " AND acl=? "
	}

	args := getArgs(&constrains)
	filterQuery := buildQueryComponent(&constrains)
	query = query + filterQuery + " GROUP BY date"

	_, err := dbMap.Select(&totalDailyDownloads, query, args...)
	if err != nil {
		panic(err.Error())
	}
	return totalDailyDownloads
}

func GetUserCountOfDownloadsOver(constrains dao.Constrains, threshold int) (int64,int64) {
	dbMap := utils.GetDBConnection(commons.SUMMARY_DB);
	defer dbMap.Db.Close()
	var err error
	query := "SELECT count(DISTINCT username) FROM dailyacct where date >= ? AND date <= ? AND tenantid= ? AND outputoctets >= ?"

	if len(constrains.ACL) > 0 {
		query = query + " AND acl=? "
	}

	args := getArgs2(&constrains, threshold)
	filterQuery := buildQueryComponent(&constrains)
	query = query + filterQuery

	count, err := dbMap.SelectNullInt(query, args...)
	argsPast := getArgs2Past(&constrains, threshold)

	countPre, err := dbMap.SelectNullInt(query, argsPast...)
	checkErr(err, "Select failed while getting user count of downloads of over")
	if count.Valid {
		return count.Int64 , countPre.Int64
	}else {
		if countPre.Valid {
			return 0 , countPre.Int64
		}else{
			return 0,0
		}
	}
}

func  GetUsersCountFromTo(constrains dao.Constrains) (int64,int64) {
	dbMap := utils.GetDBConnection(commons.SUMMARY_DB);
	defer dbMap.Db.Close()
	query := "SELECT COUNT(DISTINCT username) FROM dailyacct where date >= ? AND date <= ? AND tenantid=? "
	if len(constrains.ACL) > 0 {
		query = query + " AND acl=? "
	}

	args := getArgs(&constrains)
	filterQuery := buildQueryComponent(&constrains)
	query = query + filterQuery

	count, err := dbMap.SelectNullInt(query, args...)
	argsPast := getArgsPast(&constrains)

	countPre, err := dbMap.SelectNullInt(query, argsPast...)
	checkErr(err, "Select failed on while getting user count from to")
	if count.Valid {
		return count.Int64 , countPre.Int64
	}else {
		if countPre.Valid {
			return 0 , countPre.Int64
		}else{
			return 0,0
		}
	}
}

/*
* Users who visits more than once
*/
func GetReturningUsersCount(constrains dao.Constrains) (int64,int64) {
	dbMap := utils.GetDBConnection(commons.PORTAL_DB);
	defer dbMap.Db.Close()
	query := "SELECT COUNT(DISTINCT username) FROM accounting where acctstarttime >= ? AND acctstarttime < ? AND tenantid=? AND visits > 0"

	if len(constrains.ACL) > 0 {
		query = query + " AND acl=? "
	}
	args := getArgs(&constrains)
	filterQuery := buildQueryComponent(&constrains)
	query = query + filterQuery

	count, err := dbMap.SelectNullInt(query, args...)
	argsPast := getArgsPast(&constrains)

	countPre, err := dbMap.SelectNullInt(query, argsPast...)

	checkErr(err, "Select failed while getting returning user count")
	if count.Valid {
		return count.Int64 , countPre.Int64
	}else {
		if countPre.Valid {
			return 0 , countPre.Int64
		}else{
			return 0,0
		}
	}
}

func AddWiFiUser(user *dao.PortalUser) error {
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()
	stmtIns, err := dbMap.Db.Prepare(commons.ADD_WIFI_USER_SQL)
	defer stmtIns.Close()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmtIns.Exec(user.TenantId, user.Username,user.MaxSessionDuration, user.GroupName.String, user.ACL.String, user.Accounting)
	if err != nil {
		return err
	} else {
		if user.GroupName.String =="Master" {
			AddRadiusUser(user)
		}
	}
	return err
}

func AddRadiusUser(user *dao.PortalUser) {
	dbMap := utils.GetDBConnection(commons.RADIUS_DB);
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.ADD_RADIUS_USER)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmtIns.Exec(user.Username, "Cleartext-Password", ":=", user.Password)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
}

func UpdateWiFiUser(user *dao.PortalUser) {
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.UPDATE_WIFI_USER)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmtIns.Exec(user.MaxSessionDuration, user.ACL.String, user.Accounting, user.Username, user.GroupName.String, user.TenantId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
}

func GetAllWiFiUsers(tenantId int, draw int, r *http.Request) dao.DataTablesResponse {
	var users []dao.PortalUser
	var response dao.DataTablesResponse
	columns := []string{"username", "acl", "groupname", "visits", "acctstarttime", "acctactivationtime", "maxsessionduration", "accounting"}
	totalRecordCountQuery := "SELECT COUNT(username) FROM accounting where tenantid=" + strconv.Itoa(tenantId)
	var err error
	response.RecordsFiltered, response.RecordsTotal, err = commons.Fetch(r, "portal", "accounting", totalRecordCountQuery, columns, &users)
	if( err!= nil){
		log.Error("")
	}
	response.Data = users
	response.Draw = draw
	return response
}

func DeleteUserAccountingSession(username string,groupname string, tenantid int) error {
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()

	_, err := dbMap.Exec(commons.DELETE_WIFI_USER, username, groupname, tenantid)
	return err
}

func DeleteUserFromRadCheck(username string, tenantid int) error {
	dbMap := utils.GetDBConnection("radius");
	defer dbMap.Db.Close()

	_, err := dbMap.Exec(commons.DELETE_RADCHECk_USER, username)

	return err
}

func DeleteUserFromRadAcct(username string, tenantid int) error {
	dbMap := utils.GetDBConnection("radius");
	defer dbMap.Db.Close()

	_, err := dbMap.Exec(commons.DELETE_RADACCT_USER, username)

	return err
}

func IsWifiUserExistInGroup(tenantId int, username string, groupname string) (int, error){
	dbMap := utils.GetDBConnection(commons.PORTAL_DB);
	defer dbMap.Db.Close()
	var checkUser int
	err := dbMap.SelectOne(&checkUser, commons.IS_EXISTS_USER_NAME_IN_GROUP, username, groupname, tenantId)
	if err != nil {
		return checkUser, err
	}
	return checkUser, nil
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

