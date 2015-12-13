package wifi

import (
	"wislabs.wifi.manager/utils"
	"wislabs.wifi.manager/dao"
)

func GetUsersByOS(constrains dao.Constrains) []dao.NameValue{
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()
	var usersByOS[] dao.NameValue
	var query string

	query = "SELECT count(username) as value,os as name from useragentinfo WHERE locationid=? AND date >= ? AND date < ? group by os"

	_, err := dbMap.Select(&usersByOS, query, constrains.LocationId, constrains.From, constrains.To)
	checkErr(err, "Select failed")
	return usersByOS
}


func GetUsersByDevice(constrains dao.Constrains) []dao.NameValue{
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()
	var usersByOS[] dao.NameValue
	var query string

	query = "SELECT count(username) as value, device as name from useragentinfo WHERE locationid=? AND date >= ? AND date < ? group by os"

	_, err := dbMap.Select(&usersByOS, query, constrains.LocationId, constrains.From, constrains.To)
	checkErr(err, "Select failed")
	return usersByOS
}
