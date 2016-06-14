package wifi

import (
	"wislabs.wifi.manager/utils"
	"wislabs.wifi.manager/dao"
	"database/sql"
	"wislabs.wifi.manager/commons"
)

func GetBrowserStats(constrains dao.Constrains) []dao.NameValue {
	dbMap := utils.GetDBConnection(commons.SUMMARY_DB);
	defer dbMap.Db.Close()
	browser := []string{"Chrome", "Firefox", "IE", "IE Mobile", "Kindle", "Safari", "Safari Mobile", "Opera", "WebKit", "Chrome Mobile", "Other"}
	usersByOS := make([]dao.NameValue, 11)
	values := make([]sql.NullFloat64, 11)
	var query string
	query = "SELECT sum(chrome) as chrome, sum(firefox) as firefox, sum(ie) as ie, sum(iemobile) as iemobile, sum(kindle) as kindle, sum(safari) as safari, sum(safarimobile) as safarimobile, sum(opera) as opera, sum(webkit) as webkit, sum(chromemobile) as chromemobile, sum(other) as other from browserstats WHERE date >= ? AND date <= ? AND tenantid=? "

	args := getArgs3(&constrains)
	filterQuery := buildQueryComponent(&constrains)
	query = query + filterQuery
	smtOut, err := dbMap.Db.Prepare(query)
	defer smtOut.Close()
	err = smtOut.QueryRow(args...).Scan(&values[0], &values[1], &values[2], &values[3], &values[4], &values[5], &values[6], &values[7], &values[8], &values[9], &values[10]) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for i := 0; i < len(values); i++ {
		usersByOS[i] = dao.NameValue{browser[i], values[i].Float64}
	}
	//checkErr(err, "Select failed")
	return usersByOS
}

func GetUsersByOS(constrains dao.Constrains) []dao.NameValue {
	dbMap := utils.GetDBConnection(commons.SUMMARY_DB);
	defer dbMap.Db.Close()
	os := []string{"Android", "iOS", "Windows", "Linux", "Mac OS", "Windows Mobile", "Other"}
	usersByOS := make([]dao.NameValue, 7)
	values := make([]sql.NullFloat64, 7)
	var query string
	query = "SELECT sum(android) as android, sum(ios) as ios, sum(windows) as windows, sum(linux) as linux, sum(macos) as macos, sum(windowsmobile) as windowsmobile, sum(other) as other from osstats WHERE date >= ? AND date <= ? AND tenantid=?"

	args := getArgs3(&constrains)
	filterQuery := buildQueryComponent(&constrains)
	query = query + filterQuery

	smtOut, err := dbMap.Db.Prepare(query)
	defer smtOut.Close()
	err = smtOut.QueryRow(args...).Scan(&values[0], &values[1], &values[2], &values[3], &values[4], &values[5], &values[6]) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for i := 0; i < len(values); i++ {
		usersByOS[i] = dao.NameValue{os[i], values[i].Float64}
	}
	//checkErr(err, "Select failed")
	return usersByOS
}

func GetUsersByDevice(constrains dao.Constrains) []dao.NameValue {
	dbMap := utils.GetDBConnection(commons.SUMMARY_DB);
	defer dbMap.Db.Close()
	device := []string{"Mobile", "Tablet", "Smart TV", "Wearable", "Embedded", "Other"}
	usersByDevice := make([]dao.NameValue, 6)
	values := make([]sql.NullFloat64, 6)
	var query string
	query = "SELECT sum(mobile) as mobile, sum(tablet) as tablet, sum(smarttv) as smarttv, sum(wearable) as wearable, sum(embedded) as embedded, sum(other) as other from devicestats WHERE date >= ? AND date <= ? AND tenantid=?"

	args := getArgs3(&constrains)
	filterQuery := buildQueryComponent(&constrains)
	query = query + filterQuery

	smtOut, err := dbMap.Db.Prepare(query)
	defer smtOut.Close()
	err = smtOut.QueryRow(args...).Scan(&values[0], &values[1], &values[2], &values[3], &values[4], &values[5]) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for i := 0; i < len(values); i++ {
		usersByDevice[i] = dao.NameValue{device[i], values[i].Float64}
	}
	//checkErr(err, "Select failed")
	return usersByDevice
}
