package wifi

import (
	"wislabs.wifi.manager/utils"
	"wislabs.wifi.manager/dao"
	"database/sql"
)

func GetAggregatedDownloadsFromTo(constrains dao.Constrains) [] dao.NameValue {
	dbMap := utils.GetDBConnection("radsummary");
	defer dbMap.Db.Close()
	var totalDailyDownloads[] dao.NameValue
	query := "SELECT SUM(inputoctets) as value ,date as name FROM dailyacct where date >= ? AND date < ? AND tenantid=? "

	if len(constrains.GroupNames) > 0 {
		args := getArgs(&constrains)
		query = query + " AND groupname=? "
		for i := 1; i< len(constrains.GroupNames); i++ {
   			query = query + " OR groupname=? "
		}
		query = query + " group by date"
		_, err := dbMap.Select(&totalDailyDownloads, query, args...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic
		}
	}
	return totalDailyDownloads
}

func GetAggregatedUploadsFromTo(constrains dao.Constrains) [] dao.NameValue {
	dbMap := utils.GetDBConnection("radsummary");
	defer dbMap.Db.Close()
	var totalDailyDownloads[] dao.NameValue
	query := "SELECT SUM(outputoctets) as value ,date as name FROM dailyacct where date >= ? AND date < ? AND tenantid=? "

	if len(constrains.GroupNames) > 0 {
		args := getArgs(&constrains)
		query = query + " AND groupname=? "
		for i := 1; i< len(constrains.GroupNames); i++ {
			query = query + " OR groupname=? "
		}
		query = query + " group by date"
		_, err := dbMap.Select(&totalDailyDownloads, query, args...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic
		}
	}
	return totalDailyDownloads
}

func GetAvgDailyDownloadsPerUserFromTo(constrains dao.Constrains) [] dao.NameValue {
	dbMap := utils.GetDBConnection("radsummary");
	defer dbMap.Db.Close()
	var totalDailyDownloads[] dao.NameValue
	query := "SELECT SUM(inputoctets)/COUNT(DISTINCT username) as value ,date as name FROM dailyacct where date >= ? AND date < ? AND tenantid=? "

	if len(constrains.GroupNames) > 0 {
		args := getArgs(&constrains)
		query = query + " AND groupname=? "
		for i := 1; i< len(constrains.GroupNames); i++ {
			query = query + " OR groupname=? "
		}
		query = query + " group by date"
		_, err := dbMap.Select(&totalDailyDownloads, query, args...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic
		}
	}
	return totalDailyDownloads
}

func GetDownloadsFromTo(constrains dao.Constrains) int64 {
	dbMap := utils.GetDBConnection("radsummary");
	defer dbMap.Db.Close()
	var err error
	var count sql.NullInt64
	query := "SELECT SUM(inputoctets) FROM dailyacct where date >= ? AND date < ? AND tenantid = ? "

	if len(constrains.GroupNames) > 0 {
		args := getArgs(&constrains)
		query = query + " AND groupname=? "
		for i := 1; i< len(constrains.GroupNames); i++ {
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

func GetUploadsFromTo(constrains dao.Constrains) int64 {
	dbMap := utils.GetDBConnection("radsummary");
	defer dbMap.Db.Close()
	var err error
	var count sql.NullInt64
	query := "SELECT SUM(outputoctets) FROM dailyacct where date >= ? AND date < ? AND tenantid = ? "

	if len(constrains.GroupNames) > 0 {
		args := getArgs(&constrains)
		query = query + " AND groupname=? "
		for i := 1; i< len(constrains.GroupNames); i++ {
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

func GetTotalSessionsCountFromTo(constrains dao.Constrains) int64 {
	dbMap := utils.GetDBConnection("radsummary");
	defer dbMap.Db.Close()
	var err error
	var count sql.NullInt64
	query := "SELECT SUM(noofsessions) FROM dailyacct where date >= ? AND date < ? AND tenantid = ? "

	if len(constrains.GroupNames) > 0 {
		args := getArgs(&constrains)
		query = query + " AND groupname=? "
		for i := 1; i< len(constrains.GroupNames); i++ {
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

func GetAvgSessionsFromTo(constrains dao.Constrains) float64 {
	dbMap := utils.GetDBConnection("radsummary");
	defer dbMap.Db.Close()
	var err error
	var count sql.NullFloat64
	query := "SELECT SUM(sessionavgduration)/COUNT(DISTINCT username) FROM dailyacct where date >= ? AND date < ? AND tenantid = ? "

	if len(constrains.GroupNames) > 0 {
		args := getArgs(&constrains)
		query = query + " AND groupname=? "
		for i := 1; i< len(constrains.GroupNames); i++ {
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
		return count.Float64
	}else {
		return 0
	}
}

func GetAvgDailySessionTimePerUserFromTo(constrains dao.Constrains) [] dao.NameValue {
	dbMap := utils.GetDBConnection("radsummary");
	defer dbMap.Db.Close()
	var totalDailyDownloads[] dao.NameValue
	query := "SELECT SUM(sessionavgduration)/COUNT(DISTINCT username) as value ,date as name FROM dailyacct where date >= ? AND date < ? AND tenantid = ?"

	if len(constrains.GroupNames) > 0 {
		args := getArgs(&constrains)
		query = query + " AND groupname=? "
		for i := 1; i< len(constrains.GroupNames); i++ {
			query = query + " OR groupname=? "
		}
		query = query + " group by date"
		_, err := dbMap.Select(&totalDailyDownloads, query, args...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic
		}
	}
	return totalDailyDownloads
}

func getArgs(constrains *dao.Constrains) []interface{}{
	args := make([]interface{}, len(constrains.GroupNames)+3)
	args[0] = constrains.From
	args[1] = constrains.To
	args[2] = constrains.TenantId
	for index, value := range constrains.GroupNames { args[index+3] = value }
	return args
}