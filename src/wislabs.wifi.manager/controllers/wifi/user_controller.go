package wifi

import (
	"wislabs.wifi.manager/utils"
	"wislabs.wifi.manager/dao"
	"wislabs.wifi.manager/common"
	log "github.com/Sirupsen/logrus"
	"database/sql"
)

func GetUserCountOfDownloadsOver(constrains dao.Constrains, threshold int) int64 {
	dbMap := utils.GetDBConnection("radsummary");
	defer dbMap.Db.Close()
	var err error
	var count sql.NullInt64
	query := "SELECT count(DISTINCT username) FROM dailyacct where date >= ? AND date < ? AND outputoctets >= ?"

	if len(constrains.GroupNames) > 0 {
		args := getArgs(&constrains)
		query = query + " AND groupname=? "
		for i := 1; i < len(constrains.GroupNames); i++ {
			query = query + " OR groupname=? "
		}
		smtOut, err := dbMap.Db.Prepare(query)
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

	stmtIns, err := dbMap.Db.Prepare(common.ADD_WIFI_USER_SQL)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmtIns.Exec(user.TenantId, user.Username, user.GroupName.String, user.ACL.String)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
}

func UpdateWiFiUser(user *dao.PortalUser) {
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(common.UPDATE_WIFI_USER)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmtIns.Exec(user.ACL.String, user.Username, user.TenantId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
}

func GetAllWiFiUsers(tenantid int) []dao.PortalUser {
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()
	var users []dao.PortalUser
	_, err := dbMap.Select(&users, common.GET_ALL_WIFI_USERS, tenantid)
	checkErr(err, "Select failed")
	return users
}

func DeleteUserAccountingSession(username string, tenantid int) error {
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()

	_, err := dbMap.Exec(common.DELETE_WIFI_USER, username, tenantid)
	return err
}

func DeleteUserFromRadCheck(username string, tenantid int) error {
	dbMap := utils.GetDBConnection("radius");
	defer dbMap.Db.Close()

	_, err := dbMap.Exec(common.DELETE_RADCHECk_USER, username)

	return err
}

func DeleteUserFromRadAcct(username string, tenantid int) error {
	dbMap := utils.GetDBConnection("radius");
	defer dbMap.Db.Close()

	_, err := dbMap.Exec(common.DELETE_RADACCT_USER, username)

	return err
}


func GetUsersCountFromTo(constrains dao.Constrains) int64 {
	dbMap := utils.GetDBConnection("radsummary");
	defer dbMap.Db.Close()
	var err error
	var count sql.NullInt64
	query := "SELECT COUNT(DISTINCT username) FROM dailyacct where date >= ? AND date < ? AND tenantid=? "

	if len(constrains.GroupNames) > 0 {
		args := getArgs(&constrains)
		query = query + " AND groupname=? "
		for i := 1; i < len(constrains.GroupNames); i++ {
			query = query + " OR groupname=? "
		}
		smtOut, err := dbMap.Db.Prepare(query)
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
	query := "SELECT COUNT(DISTINCT username) FROM accounting where acctstarttime >= ? AND acctstarttime < ? AND tenantid=? AND visits > 1"

	if len(constrains.GroupNames) > 0 {
		args := getArgs(&constrains)
		query = query + " AND groupname=? "
		for i := 1; i < len(constrains.GroupNames); i++ {
			query = query + " OR groupname=? "
		}
		smtOut, err := dbMap.Db.Prepare(query)
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

