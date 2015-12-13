package wifi

import (
	"wislabs.wifi.manager/utils"
	"wislabs.wifi.manager/dao"
	"wislabs.wifi.manager/common"
	log "github.com/Sirupsen/logrus"
	"database/sql"
)

func GetUserCountOfDownloadsOver(constrains dao.Constrains, threshold int) int64{
	dbMap := utils.GetDBConnection("radsummary");
	defer dbMap.Db.Close()
	var count sql.NullInt64
	var err error
	if len(constrains.LocationId) >0 {
		smtOut, err := dbMap.Db.Prepare(common.GET_USER_COUNT_OF_DOWNLOADS_OVER_LOCATION)
		defer  smtOut.Close()
		err = smtOut.QueryRow( constrains.From, constrains.To, constrains.LocationId, threshold ).Scan(&count) // WHERE number = 13
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}else{
		smtOut, err := dbMap.Db.Prepare(common.GET_USER_COUNT_OF_DOWNLOADS_OVER)
		defer  smtOut.Close()
		err = smtOut.QueryRow( constrains.From, constrains.To, threshold).Scan(&count) // WHERE number = 13
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

func AddWiFiUser(user *dao.PortalUser){
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(common.ADD_WIFI_USER_SQL) // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmtIns.Exec(user.Username, user.Location.Int64)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
}

func UpdateWiFiUser(user *dao.PortalUser){
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(common.UPDATE_WIFI_USER_SQL) // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmtIns.Exec(user.ACL.String, user.Username)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
}

func GetAllWiFiUsers() []dao.PortalUser{
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()
	var users []dao.PortalUser

	_, err := dbMap.Select(&users,common.GET_ALL_WIFI_USER_SQL)
	checkErr(err, "Select failed")
	return users
}

func DeleteUserAccountingSession(username string) error{
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()

	_, err := dbMap.Exec("DELETE FROM accounting where username=?", username)
   return err
}

func DeleteUserFromRadCheck(username string) error{
	dbMap := utils.GetDBConnection("radius");
	defer dbMap.Db.Close()

	_, err := dbMap.Exec("DELETE FROM radcheck WHERE username = ?", username)

	return err
}

func DeleteUserFromRadAcct(username string) error{
	dbMap := utils.GetDBConnection("radius");
	defer dbMap.Db.Close()

	_, err := dbMap.Exec("DELETE FROM radacct WHERE username = ?", username)

	return err
}


func GetUsersCountFromTo(from string, to string, location string) int64{
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()
	var count int64
	count, err := dbMap.SelectInt("SELECT COUNT(DISTINCT username) FROM accounting where acctstarttime >= ? AND acctstarttime < ? AND locationid = ?", from, to, location)
	checkErr(err, "Select failed")
	return count
}

/*
* Users who visits more than once
*/
func GetReturningUsers(from string, to string, location string) int64{
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()
	var count int64
	count, err := dbMap.SelectInt("SELECT COUNT(DISTINCT username) FROM accounting where acctstarttime >= ? AND acctstarttime < ? AND locationid = ? AND visits > 1", from, to, location)
	checkErr(err, "Select failed")
	return count
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

