package wifi

import (
	"wislabs.wifi.manager/utils"
	"wislabs.wifi.manager/dao"
	"strconv"
	"wislabs.wifi.manager/commons"
)

func SummaryDetailsFromTo(constrains dao.Constrains) [][]string {

	dbMap := utils.GetDBConnection("summary");
	defer dbMap.Db.Close()
	var dailyAccData[] dao.SummaryDailyAcctAll
	query := "SELECT * FROM dailyacct where date >= ? AND date <= ? AND tenantid=? "
	if len(constrains.ACL) > 0 {
		query = query + " AND acl=? "
	}

	args := getArgs(&constrains)
	filterQuery := buildQueryComponent(&constrains)
	query = query + filterQuery

	_, err := dbMap.Select(&dailyAccData, query, args...)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic
	}

	CSVcontent := make([][]string, len(dailyAccData) + 1)

	CSVcontent[0] = make([]string, 17)
	CSVcontent[0][0] = "TenantId"
	CSVcontent[0][1] = "Username"
	CSVcontent[0][2] = "Date"
	CSVcontent[0][3] = "NumOfSessions"
	CSVcontent[0][4] = "TotalSessionDuration"
	CSVcontent[0][5] = "MaxSessionDuration"
	CSVcontent[0][6] = "MinSessionDuration"
	CSVcontent[0][7] = "AvgSessionDuration"
	CSVcontent[0][8] = "Downloads"
	CSVcontent[0][9] = "Uploads"
	CSVcontent[0][10] = "NASIPAddress"
	CSVcontent[0][11] = "FramedIPAddress"
	CSVcontent[0][12] = "CalledStationId"
	CSVcontent[0][13] = "SSID"
	CSVcontent[0][14] = "CalledStatoinMAC"
	CSVcontent[0][15] = "GroupName"

	for i := 1; i < len(CSVcontent) - 1; i++ {
		CSVcontent[i] = make([]string, 17)
		CSVcontent[i][0] = strconv.Itoa(dailyAccData[i].Tenantid)
		CSVcontent[i][1] = dailyAccData[i].Username
		CSVcontent[i][2] = dailyAccData[i].Date.String
		CSVcontent[i][3] = strconv.Itoa(dailyAccData[i].Noofsessions)
		CSVcontent[i][4] = strconv.Itoa(dailyAccData[i].Totalsessionduration)
		CSVcontent[i][5] = strconv.Itoa(dailyAccData[i].Sessionmaxduration)
		CSVcontent[i][6] = strconv.Itoa(dailyAccData[i].Sessionminduration)
		CSVcontent[i][7] = strconv.Itoa(dailyAccData[i].Sessionavgduration)
		CSVcontent[i][8] = strconv.FormatInt(dailyAccData[i].Inputoctets, 10)
		CSVcontent[i][9] = strconv.FormatInt(dailyAccData[i].Outputoctets, 10)
		CSVcontent[i][10] = dailyAccData[i].Nasipaddress
		CSVcontent[i][11] = dailyAccData[i].Framedipaddress
		CSVcontent[i][12] = dailyAccData[i].Calledstationid
		CSVcontent[i][13] = dailyAccData[i].Ssid.String
		CSVcontent[i][14] = dailyAccData[i].Calledstationmac.String
		CSVcontent[i][15] = dailyAccData[i].Groupname.String
	}
	return CSVcontent
}

func GetAccessPointAggregatedDataFromTo(constrains dao.Constrains) [] dao.AccessPoint {
	dbMap := utils.GetDBConnection(commons.SUMMARY_DB);
	defer dbMap.Db.Close()
	var accessPointData[] dao.AccessPoint

	query := "SELECT calledstationmac as calledstationmac," +
	"SUM(outputoctets) as totaloutputoctets," +
	"SUM(inputoctets) as totalinputoctets," +
	"SUM(noofsessions) as totalsessions ," +
	"COUNT(DISTINCT username) as totalusers," +
	"SUM(inputoctets)/COUNT(DISTINCT username) as avgdataperuser," +
	"SUM(totalsessionduration)/SUM(noofsessions) as avgdatapersessiontime " +
	"FROM dailyacct " +
	"WHERE date >= ? AND date <= ? AND tenantid=? "

	if len(constrains.ACL) > 0 {
		query = query + " AND acl=? "
	}
	args := getArgs(&constrains)
	filterQuery := buildQueryComponent(&constrains)
	query = query + filterQuery + " GROUP BY calledstationmac"
	_, err := dbMap.Select(&accessPointData, query, args...)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic
	}
	return accessPointData
}

func GetLongLatLocationByMacAddress(mac string) dao.LongLatMac {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	var longlatbymac dao.LongLatMac

	query := "SELECT longitude as longitude, latitude as latitude, mac as mac, apname as apname from aplocations where mac=?"
	err := dbMap.SelectOne(&longlatbymac, query, mac)
	if err != nil {
		return longlatbymac
	}
	return longlatbymac
}

func GetAggregatedDownloadsFromTo(constrains dao.Constrains) [] dao.NameValue {
	dbMap := utils.GetDBConnection("summary");
	defer dbMap.Db.Close()
	var totalDailyDownloads[] dao.NameValue
	query := "SELECT SUM(inputoctets) as value, date as name FROM dailyacct where date >= ? AND date <= ? AND tenantid=? "
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

func GetAggregatedUploadsFromTo(constrains dao.Constrains) [] dao.NameValue {
	dbMap := utils.GetDBConnection("summary");
	defer dbMap.Db.Close()
	var totalDailyDownloads[] dao.NameValue
	query := "SELECT SUM(outputoctets) as value ,date as name FROM dailyacct where date >= ? AND date <= ? AND tenantid=? "
	if len(constrains.ACL) > 0 {
		query = query + " AND acl=? "
	}

	args := getArgs(&constrains)
	filterQuery := buildQueryComponent(&constrains)
	query = query + filterQuery + " GROUP BY date"

	_, err := dbMap.Select(&totalDailyDownloads, query, args...)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic
	}
	return totalDailyDownloads
}

func GetAvgDailyDownloadsPerUserFromTo(constrains dao.Constrains) [] dao.NameValue {
	dbMap := utils.GetDBConnection("summary");
	defer dbMap.Db.Close()
	var totalDailyDownloads[] dao.NameValue
	query := "SELECT SUM(inputoctets)/COUNT(DISTINCT username) as value ,date as name FROM dailyacct where date >= ? AND date <= ? AND tenantid=? "
	if len(constrains.ACL) > 0 {
		query = query + " AND acl=? "
	}

	args := getArgs(&constrains)
	filterQuery := buildQueryComponent(&constrains)
	query = query + filterQuery + " GROUP BY date"
	_, err := dbMap.Select(&totalDailyDownloads, query, args...)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic
	}
	return totalDailyDownloads
}

func GetDownloadsFromTo(constrains dao.Constrains) (int64, int64) {
	dbMap := utils.GetDBConnection("summary");
	defer dbMap.Db.Close()
	var err error
	query := "SELECT SUM(inputoctets) FROM dailyacct where date >= ? AND date <= ? AND tenantid = ? "
	if len(constrains.ACL) > 0 {
		query = query + " AND acl=? "
	}
	args := getArgs(&constrains)
	filterQuery := buildQueryComponent(&constrains)
	query = query + filterQuery

	count, err := dbMap.SelectNullInt(query, args...)
	argsPast := getArgsPast(&constrains)

	countPre, err := dbMap.SelectNullInt(query, argsPast...)
	checkErr(err, "Select failed on Get downloads")

	if count.Valid {
		return count.Int64, countPre.Int64
	} else {
		if countPre.Valid {
			return 0, countPre.Int64
		} else {
			return 0, 0
		}
	}
}

func GetUploadsFromTo(constrains dao.Constrains) (int64, int64) {
	dbMap := utils.GetDBConnection("summary");
	defer dbMap.Db.Close()
	var err error
	query := "SELECT SUM(outputoctets) FROM dailyacct where date >= ? AND date <= ? AND tenantid = ? "
	if len(constrains.ACL) > 0 {
		query = query + " AND acl=? "
	}
	args := getArgs(&constrains)
	filterQuery := buildQueryComponent(&constrains)
	query = query + filterQuery

	count, err := dbMap.SelectNullInt(query, args...)
	argsPast := getArgsPast(&constrains)

	countPre, err := dbMap.SelectNullInt(query, argsPast...)

	checkErr(err, "Select failed on Get uploads")
	if count.Valid {
		return count.Int64, countPre.Int64
	} else {
		if countPre.Valid {
			return 0, countPre.Int64
		} else {
			return 0, 0
		}
	}
}

func GetTotalSessionsCountFromTo(constrains dao.Constrains) (int64, int64) {
	dbMap := utils.GetDBConnection("summary");
	defer dbMap.Db.Close()
	var err error
	query := "SELECT SUM(noofsessions) FROM dailyacct where date >= ? AND date <= ? AND tenantid = ? "
	if len(constrains.ACL) > 0 {
		query = query + " AND acl=? "
	}
	args := getArgs(&constrains)
	filterQuery := buildQueryComponent(&constrains)
	query = query + filterQuery

	count, err := dbMap.SelectNullInt(query, args...)
	argsPast := getArgsPast(&constrains)

	countPre, err := dbMap.SelectNullInt(query, argsPast...)

	checkErr(err, "Select failed on getting total session count")
	if count.Valid {
		return count.Int64, countPre.Int64
	} else {
		if countPre.Valid {
			return 0, countPre.Int64
		} else {
			return 0, 0
		}
	}
}

func GetAvgSessionsFromTo(constrains dao.Constrains) (float64, float64) {
	dbMap := utils.GetDBConnection("summary");
	defer dbMap.Db.Close()
	var err error
	query := "SELECT SUM(totalsessionduration)/SUM(noofsessions) FROM dailyacct where date >= ? AND date <= ? AND tenantid = ? "
	if len(constrains.ACL) > 0 {
		query = query + " AND acl=? "
	}
	args := getArgs(&constrains)
	filterQuery := buildQueryComponent(&constrains)
	query = query + filterQuery

	count, err := dbMap.SelectNullFloat(query, args...)
	argsPast := getArgsPast(&constrains)

	countPre, err := dbMap.SelectNullFloat(query, argsPast...)

	checkErr(err, "Select failed on getting avg session count")
	if count.Valid {
		return count.Float64, countPre.Float64
	} else {
		if countPre.Valid {
			return 0, countPre.Float64
		} else {
			return 0, 0
		}
	}
}

func GetAvgDailySessionTimePerUserFromTo(constrains dao.Constrains) [] dao.NameValue {
	dbMap := utils.GetDBConnection("summary");
	defer dbMap.Db.Close()
	var totalDailyDownloads[] dao.NameValue
	query := "SELECT SUM(totalsessionduration)/SUM(noofsessions) as value ,date as name FROM dailyacct where date >= ? AND date <= ? AND tenantid = ? "
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

func buildQueryComponent(constrains *dao.Constrains) (string) {
	query := " "
	if len(constrains.Parameters) > 0 {
		query = " AND ( " + constrains.Criteria + "=? "
		for i := 1; i < len(constrains.Parameters); i++ {
			query = query + " OR " + constrains.Criteria + "=? "
		}
		query = query + ") "
	}
	return query
}

func getArgs(constrains *dao.Constrains) []interface{} {
	var startSize = 3;
	var startIndex = 3
	if len(constrains.ACL) > 0 {
		startSize = 4
	}
	args := make([]interface{}, len(constrains.Parameters) + startSize)
	args[0] = constrains.From
	args[1] = constrains.To
	args[2] = constrains.TenantId
	if len(constrains.ACL) > 0 {
		args[3] = constrains.ACL
		startIndex++;
	}
	for index, value := range constrains.Parameters {
		args[index + startIndex] = value
	}
	return args
}

func getArgsPast(constrains *dao.Constrains) []interface{} {
	var startSize = 3;
	var startIndex = 3
	if len(constrains.ACL) > 0 {
		startSize = 4
	}
	args := make([]interface{}, len(constrains.Parameters) + startSize)
	args[0] = constrains.PreFrom
	args[1] = constrains.PreTo
	args[2] = constrains.TenantId
	if len(constrains.ACL) > 0 {
		args[3] = constrains.ACL
		startIndex++;
	}
	for index, value := range constrains.Parameters {
		args[index + startIndex] = value
	}
	return args
}

func getArgs2(constrains *dao.Constrains, threshold int) []interface{} {
	var arraySize = 4;
	if len(constrains.ACL) > 0 {
		arraySize++
	}
	args := make([]interface{}, len(constrains.Parameters) + arraySize)
	args[0] = constrains.From
	args[1] = constrains.To
	args[2] = constrains.TenantId
	args[3] = threshold
	if len(constrains.ACL) > 0 {
		args[4] = constrains.ACL
	}
	for index, value := range constrains.Parameters {
		args[index + arraySize] = value
	}
	return args
}

func getArgs2Past(constrains *dao.Constrains, threshold int) []interface{} {
	var arraySize = 4;
	if len(constrains.ACL) > 0 {
		arraySize++
	}
	args := make([]interface{}, len(constrains.Parameters) + arraySize)
	args[0] = constrains.PreFrom
	args[1] = constrains.PreTo
	args[2] = constrains.TenantId
	args[3] = threshold
	if len(constrains.ACL) > 0 {
		args[4] = constrains.ACL
	}
	for index, value := range constrains.Parameters {
		args[index + arraySize] = value
	}
	return args
}

func getArgs3(constrains *dao.Constrains) []interface{} {
	args := make([]interface{}, len(constrains.Parameters) + 3)
	args[0] = constrains.From
	args[1] = constrains.To
	args[2] = constrains.TenantId
	for index, value := range constrains.Parameters {
		args[index + 3] = value
	}
	return args
}

