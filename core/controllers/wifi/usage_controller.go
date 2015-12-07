package wifi

import (
	"wifi-manager/core/utils"
	"wifi-manager/core/dao"
	"database/sql"
)

func GetAgregatedDownloadsFromTo(constrains dao.Constrains) [] dao.NameValue{
	dbMap := utils.GetDBConnection("radsummary");
	defer dbMap.Db.Close()
	var totalDailyDownloads[] dao.NameValue
	if len(constrains.LocationId) >0 {
		query :="SELECT SUM(inputoctets) as value ,date as name FROM dailyacct where date >= ? AND date < ? AND location = ? group by date"
		_, err := dbMap.Select(&totalDailyDownloads, query, constrains.From, constrains.To, constrains.LocationId)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}else{
		query :="SELECT SUM(inputoctets) as value ,date as name FROM dailyacct where date >= ? AND date < ? group by date"
		_, err := dbMap.Select(&totalDailyDownloads, query, constrains.From, constrains.To)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}
	return totalDailyDownloads
}

func GetDownloadsFromTo(constrains dao.Constrains) int64{
	dbMap := utils.GetDBConnection("radsummary");
	defer dbMap.Db.Close()
    var err error
	var count sql.NullInt64
	if len(constrains.LocationId) >0 {
		smtOut, err := dbMap.Db.Prepare("SELECT SUM(inputoctets) FROM dailyacct where date >= ? AND date < ? AND locationid = ?")
		defer  smtOut.Close()
		err = smtOut.QueryRow( constrains.From, constrains.To, constrains.LocationId).Scan(&count) // WHERE number = 13
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}else{
		smtOut, err := dbMap.Db.Prepare("SELECT SUM(inputoctets) FROM dailyacct where date >= ? AND date < ? ")
		defer  smtOut.Close()
		err = smtOut.QueryRow( constrains.From, constrains.To).Scan(&count) // WHERE number = 13
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

func GetUploadsFromTo(constrains dao.Constrains) int64{
	dbMap := utils.GetDBConnection("radsummary");
	defer dbMap.Db.Close()
	var err error
	var count sql.NullInt64
	if len(constrains.LocationId) >0 {
		smtOut, err := dbMap.Db.Prepare("SELECT SUM(outputoctets) FROM dailyacct where date >= ? AND date < ? AND locationid = ?")
		defer  smtOut.Close()
		err = smtOut.QueryRow( constrains.From, constrains.To, constrains.LocationId).Scan(&count) // WHERE number = 13
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}else{
		smtOut, err := dbMap.Db.Prepare("SELECT SUM(outputoctets) FROM dailyacct where date >= ? AND date < ? ")
		defer  smtOut.Close()
		err = smtOut.QueryRow( constrains.From, constrains.To).Scan(&count) // WHERE number = 13
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

func GetTotalSessionsCountFromTo(constrains dao.Constrains) int64{
	dbMap := utils.GetDBConnection("radsummary");
	defer dbMap.Db.Close()
	var err error
	var count sql.NullInt64
	if len(constrains.LocationId) >0 {
		smtOut, err := dbMap.Db.Prepare("SELECT SUM(noofsessions) FROM dailyacct where date >= ? AND date < ? AND locationid = ?")
		defer  smtOut.Close()
		err = smtOut.QueryRow( constrains.From, constrains.To, constrains.LocationId).Scan(&count) // WHERE number = 13
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}else{
		smtOut, err := dbMap.Db.Prepare("SELECT SUM(noofsessions) FROM dailyacct where date >= ? AND date < ? ")
		defer  smtOut.Close()
		err = smtOut.QueryRow( constrains.From, constrains.To).Scan(&count) // WHERE number = 13
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

func GetAvgSessionsFromTo(constrains dao.Constrains) float64{
	dbMap := utils.GetDBConnection("radsummary");
	defer dbMap.Db.Close()
	var err error
	var count sql.NullFloat64
	if len(constrains.LocationId) >0 {
		smtOut, err := dbMap.Db.Prepare("SELECT AVG(sessionavgduration) FROM dailyacct where date >= ? AND date < ? AND locationid = ?")
		defer  smtOut.Close()
		err = smtOut.QueryRow( constrains.From, constrains.To, constrains.LocationId).Scan(&count) // WHERE number = 13
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}else{
		smtOut, err := dbMap.Db.Prepare("SELECT AVG(sessionavgduration) FROM dailyacct where date >= ? AND date < ? ")
		defer  smtOut.Close()
		err = smtOut.QueryRow( constrains.From, constrains.To).Scan(&count) // WHERE number = 13
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

