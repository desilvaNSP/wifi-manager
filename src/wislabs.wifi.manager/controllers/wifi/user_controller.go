package wifi

import (
	"wislabs.wifi.manager/utils"
	"wislabs.wifi.manager/dao"
	"wislabs.wifi.manager/commons"
	log "github.com/Sirupsen/logrus"
	"database/sql"
	"net/http"
	"strconv"
)

func GetDailyUserCountSeriesFromTo(constrains dao.Constrains) [] dao.NameValue {
	dbMap := utils.GetDBConnection("summary");
	defer dbMap.Db.Close()
	var totalDailyDownloads[] dao.NameValue
	query := "SELECT COUNT(DISTINCT username) as value ,date as name FROM dailyacct where date >= ? AND date < ? AND tenantid=? AND acl=? "

	if len(constrains.GroupNames) > 0 {
		args := getArgs(&constrains)
		query = query + " AND (groupname=? "
		for i := 1; i < len(constrains.GroupNames); i++ {
			query = query + " OR groupname=? "
		}
		query = query + " )group by date"
		_, err := dbMap.Select(&totalDailyDownloads, query, args...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic
		}
	}
	return totalDailyDownloads
}

func GetUserCountOfDownloadsOver(constrains dao.Constrains, threshold int) int64 {
	dbMap := utils.GetDBConnection("summary");
	defer dbMap.Db.Close()
	var err error
	var count sql.NullInt64
	query := "SELECT count(DISTINCT username) FROM dailyacct where date >= ? AND date < ? AND tenantid= ? AND inputoctets >= ?"

	if len(constrains.GroupNames) > 0 {
		args := getArgs2(&constrains, threshold)
		query = query + " AND (groupname=? "
		for i := 1; i < len(constrains.GroupNames); i++ {
			query = query + " OR groupname=? "
		}
		smtOut, err := dbMap.Db.Prepare(query + ")")
		defer smtOut.Close()
		err = smtOut.QueryRow(args...).Scan(&count) // WHERE number = 13
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}
	checkErr(err, "Select failed on Get downloads")
	if count.Valid {
		return count.Int64
	}else {
		return 0
	}
}

func AddWiFiUser(user *dao.PortalUser) {
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.ADD_WIFI_USER_SQL)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmtIns.Exec(user.TenantId, user.Username, user.GroupName.String, user.ACL.String)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	AddRadiusUser(user)
	defer stmtIns.Close()
}

func AddRadiusUser(user *dao.PortalUser) {
	dbMap := utils.GetDBConnection(commons.RADIUS_DB_NAME);
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
	_, err = stmtIns.Exec(user.ACL.String, user.Username, user.TenantId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
}

func GetAllWiFiUsers(tenantId int, draw int, r *http.Request) dao.DataTablesResponce {
	var users []dao.PortalUser
	var response dao.DataTablesResponce
	columns := []string{"username", "acl", "groupname", "visits", "acctstarttime", "acctstoptime", "acctlastupdatedtime"}
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

func DeleteUserAccountingSession(username string, tenantid int) error {
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()

	_, err := dbMap.Exec(commons.DELETE_WIFI_USER, username, tenantid)
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

func GetUsersCountFromTo(constrains dao.Constrains) int64 {
	dbMap := utils.GetDBConnection("summary");
	defer dbMap.Db.Close()
	var err error
	var count sql.NullInt64
	query := "SELECT COUNT(DISTINCT username) FROM dailyacct where date >= ? AND date < ? AND tenantid=? AND acl=?"

	if len(constrains.GroupNames) > 0 {
		args := getArgs(&constrains)
		query = query + " AND (groupname=? "
		for i := 1; i < len(constrains.GroupNames); i++ {
			query = query + " OR groupname=? "
		}
		smtOut, err := dbMap.Db.Prepare(query + ")")
		defer smtOut.Close()
		err = smtOut.QueryRow(args...).Scan(&count) // WHERE number = 13
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}
	checkErr(err, "Select failed on Get downloads")
	if count.Valid {
		return count.Int64
	}else {
		return 0
	}
}

/*
* Users who visits more than once
*/
func GetReturningUsers(constrains dao.Constrains) int64 {
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()
	var err error
	var count sql.NullInt64
	query := "SELECT COUNT(DISTINCT username) FROM accounting where acctstarttime >= ? AND acctstarttime < ? AND tenantid=? AND visits > 0"

	if len(constrains.GroupNames) > 0 {
		args := getArgs3(&constrains)
		query = query + " AND ( groupname=? "
		for i := 1; i < len(constrains.GroupNames); i++ {
			query = query + " OR groupname=? "
		}
		smtOut, err := dbMap.Db.Prepare(query + ")")
		defer smtOut.Close()
		err = smtOut.QueryRow(args...).Scan(&count) // WHERE number = 13
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}
	checkErr(err, "Select failed on Get downloads")
	if count.Valid {
		return count.Int64
	}else {
		return 0
	}
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

