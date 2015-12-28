package dashboard
//
//import (
//	"wislabs.wifi.manager/utils"
//	"wislabs.wifi.manager/dao"
//	"wislabs.wifi.manager/common"
////	log "github.com/Sirupsen/logrus"
//	"database/sql"
//)
//
//func CreateNewDashboard(constrains dao.Constrains, threshold int) int64 {
//	dbMap := utils.GetDBConnection("radsummary");
//	defer dbMap.Db.Close()
//	var count sql.NullInt64
//	var err error
//	if len(constrains.LocationId) > 0 {
//		smtOut, err := dbMap.Db.Prepare(common.GET_USER_COUNT_OF_DOWNLOADS_OVER_LOCATION)
//		defer smtOut.Close()
//		err = smtOut.QueryRow(constrains.From, constrains.To, constrains.LocationId, threshold).Scan(&count) // WHERE number = 13
//		if err != nil {
//			panic(err.Error()) // proper error handling instead of panic in your app
//		}
//	}else {
//		smtOut, err := dbMap.Db.Prepare(common.GET_USER_COUNT_OF_DOWNLOADS_OVER)
//		defer smtOut.Close()
//		err = smtOut.QueryRow(constrains.From, constrains.To, threshold).Scan(&count) // WHERE number = 13
//		if err != nil {
//			panic(err.Error()) // proper error handling instead of panic in your app
//		}
//	}
//
//	checkErr(err, "Select failed on Get downloads")
//	if count.Valid {
//		return count.Int64
//	}else {
//		return 0
//	}
//}