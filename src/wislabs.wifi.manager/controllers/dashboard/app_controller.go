package dashboard

import (
	"wislabs.wifi.manager/utils"
	"wislabs.wifi.manager/dao"
	"wislabs.wifi.manager/commons"
	"fmt"
	"strconv"
)

func CreateNewDashboardApp(dashboardAppInfo dao.DashboardAppInfo) {
	appId, err := AddDashboardApp(&dashboardAppInfo)
	if (err == nil) {
		AddDashboardAppUsers(&dashboardAppInfo.Users, appId)
		AddDashboardAppGroups(&dashboardAppInfo.Groups, appId)
		AddDashboardAppMetrics(&dashboardAppInfo.Metrics, appId)
		AddDashboardAppAcls(dashboardAppInfo.Acls,appId)
	}
}

func UpdateDashBoardSettings(dashboardAppInfo dao.DashboardAppInfo) {

	UpadateDashboardAppUsers(&dashboardAppInfo);
	UpadateDashboardAppGroups(&dashboardAppInfo);
	UpadateDashboardAppMetrics(&dashboardAppInfo);
	UpadateDashboardAppAcls(&dashboardAppInfo);
	UpadateDashboardAppAggregateValue(&dashboardAppInfo);



}

func GetAllDashboardAppsOfUser(username string, tenantId int) []dao.DashboardApp {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var apps []dao.DashboardApp
	_, err := dbMap.Select(&apps, commons.GET_DASHBOARD_USER_APPS, username, tenantId)
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		return apps
	}
	return apps
}

func GetDashboardUsersOfApp(appId int) []dao.DashboardAppUser {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var users []dao.DashboardAppUser
	_, err := dbMap.Select(&users, commons.GET_DASHBOARD_APP_USERS, appId)
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		return users
	}
	return users
}

func GetDashboardMetricsOfApp(appId int) []dao.DashboardMetric {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var metrics []dao.DashboardMetric
	_, err := dbMap.Select(&metrics, commons.GET_DASHBOARD_APP_METRICS, appId)
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		return metrics
	}
	return metrics
}

func GetDashboardGroupsOfApp(appId int) []dao.DashboardAppGroup {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var groups []dao.DashboardAppGroup
	_, err := dbMap.Select(&groups, commons.GET_DASHBOARD_APP_GROUPS, appId)
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		return groups
	}
	return groups
}
func GetDashboardAclsOfApp(appId int) string {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var acls []dao.DashboardAppAcls
	_, err := dbMap.Select(&acls, commons.GET_DASHBOARD_APP_ACLS, appId)
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		return acls[0].Acls
	}
	return acls[0].Acls
}

func GetDashboardAggregateOfApp(appId int) string{
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var strAggregate []string
	_, err := dbMap.Select(&strAggregate, commons.GET_DASHBOARD_APP_AGGREGATE, appId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return strAggregate[0]
}

func GetAllDashboardMetrics(tenantId int) []dao.DashboardMetric {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var metrics []dao.DashboardMetric
	_, err := dbMap.Select(&metrics, commons.GET_ALL_DASHBOARD_METRICS, tenantId)
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		return metrics
	}
	return metrics
}


func GetAllDashboardAclTypes( ) []string {
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()
	var aclsTypes []string

	_, err := dbMap.Select(&aclsTypes, commons.GET_ALL_DASHBOARD_ACLS)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return aclsTypes
}

func AddDashboardApp(app *dao.DashboardAppInfo) (int64, error) {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.ADD_DASHBOARD_APP)
	defer stmtIns.Close()

	if err != nil {
		return 0, err
	}
	result, err := stmtIns.Exec(app.TenantId, app.Name, app.Aggregate)
	if err != nil {
		return 0, err
	} else {
		id, err := result.LastInsertId()
		return id, err
	}
}

func AddDashboardAppMetrics(appMetrics *[]dao.DashboardAppMetric, appId int64) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.ADD_DASHBOARD_APP_METRIC)
	defer stmtIns.Close()

	if err != nil {
		return err
	}
	for i := 0; i < len(*appMetrics); i++ {
		_, err = stmtIns.Exec(appId, (*appMetrics)[i].MetricId)
	}
	return err
}

func AddDashboardAppGroups(appGroup *[]dao.DashboardAppGroup, appId int64) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.ADD_DASHBOARD_APP_GROUP)
	defer stmtIns.Close()

	if err != nil {
		return err
	}
	for i := 0; i < len(*appGroup); i++ {
		_, err = stmtIns.Exec(appId,i+1, (*appGroup)[i].GroupName)
	}
	return err
}

func AddDashboardAppUsers(appUsers *[]dao.DashboardAppUser, appId int64) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.ADD_DASHBOARD_APP_USER)
	defer stmtIns.Close()

	if err != nil {
		return err
	}
	for i := 0; i < len(*appUsers); i++ {
		fmt.Printf((*appUsers)[i].UserName)
		_, err = stmtIns.Exec((*appUsers)[i].TenantId, appId, (*appUsers)[i].UserName)
	}
	return err
}

func AddDashboardAppAcls(acls string, appId int64) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.ADD_DASHBOARD_ACLS)
	defer stmtIns.Close()

	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(appId, acls)
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
	return err
}

func DeleteDashboardApp(appId int, tenantId int) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.DELETE_DASHBOARD_APP)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmtIns.Exec(appId, tenantId)
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
	return err
}


func UpadateDashboardAppAggregateValue(dashboardAppInfo  *dao.DashboardAppInfo){
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.UPDATE_DB_APP_AGGREGATE_VALUE)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmtIns.Exec(dashboardAppInfo.Aggregate, dashboardAppInfo.AppId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()

}


func UpadateDashboardAppGroups(dashboardAppInfo  *dao.DashboardAppInfo){
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var appGroups *[]dao.DashboardAppGroup
	appGroups = &dashboardAppInfo.Groups
	var groups []dao.DashboardAppGroup
	_, err := dbMap.Select(&groups, commons.GET_DASHBOARD_APP_GROUPS, dashboardAppInfo.AppId)
	if err != nil {
		panic(err.Error())
	}

	stmtInsDelete, err := dbMap.Db.Prepare(commons.DELETE_OLD_DB_APP_GROUPS)
	defer stmtInsDelete.Close()
	stmtInsAdd, err := dbMap.Db.Prepare(commons.ADD_NEW_DB_APP_GROUPS)
	defer stmtInsAdd.Close()
	if err != nil {
		panic(err.Error())
	}

	var Length = (len(*appGroups))
	var countID = len(groups)
	if( Length <= countID){
		for i := 0; i < countID; i++ {
			if !(checkContainsGroups(groups[i].GroupName ,(*appGroups))) {
				_,err = stmtInsDelete.Exec(&dashboardAppInfo.AppId, groups[i].GroupName)
			}
		}
	}else{
		for j := 0; j < Length; j++ {
			if !(checkContainsGroups((*appGroups)[j].GroupName ,groups)) {
				_, err = stmtInsAdd.Exec(&dashboardAppInfo.AppId, (*appGroups)[j].GroupName)
			}
		}
	}

	if err != nil {
		panic(err.Error())
	}
}

func UpadateDashboardAppMetrics(dashboardAppInfo  *dao.DashboardAppInfo){
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var updatedMetrics *[]dao.DashboardAppMetric
	updatedMetrics = &dashboardAppInfo.Metrics;
	var metrics []dao.DashboardAppMetric
	_, err := dbMap.Select(&metrics, commons.GET_EXIST_DASHBOARD_APP_METRICS, dashboardAppInfo.AppId)
	if err != nil {
		panic(err.Error())
	}

	stmtInsDelete, err := dbMap.Db.Prepare(commons.DELETE_OLD_DB_APP_METRICS)
	defer stmtInsDelete.Close()
	stmtInsAdd, err := dbMap.Db.Prepare(commons.ADD_NEW_DB_APP_METRICS)
	defer stmtInsAdd.Close()
	if err != nil {
		panic(err.Error())
	}

	var Length = (len(*updatedMetrics))
	var countID = len(metrics)
	if( Length <= countID){
		for i := 0; i < countID; i++ {
			if !(checkContainsMetrics(metrics[i].MetricId ,(*updatedMetrics))) {
				_,err = stmtInsDelete.Exec(&dashboardAppInfo.AppId, metrics[i].MetricId)
			}
		}
	}else{
		for j := 0; j < Length; j++ {
			if !(checkContainsMetrics((*updatedMetrics)[j].MetricId ,metrics)) {
				_, err = stmtInsAdd.Exec(&dashboardAppInfo.AppId, (*updatedMetrics)[j].MetricId)
			}
		}
	}
	if err != nil {
		panic(err.Error())
	}
}


func UpadateDashboardAppUsers(dashboardAppInfo  *dao.DashboardAppInfo){
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var updatedUsers *[]dao.DashboardAppUser
	updatedUsers = &dashboardAppInfo.Users;
	var users []dao.DashboardAppUser
	_, err := dbMap.Select(&users, commons.GET_DASHBOARD_APP_USERS, dashboardAppInfo.AppId)
	if err != nil {
		panic(err.Error())
	}

	stmtInsDelete, err := dbMap.Db.Prepare(commons.DELETE_OLD_DB_APP_USERS)
	defer stmtInsDelete.Close()
	stmtInsAdd, err := dbMap.Db.Prepare(commons.ADD_NEW_DB_APP_USERS)
	defer stmtInsAdd.Close()
	if err != nil {
		panic(err.Error())
	}

	var Length = (len(*updatedUsers))
	var countUsers = len(users)e
	if( Length <= countUsers){
		for i := 0; i < countUsers; i++ {
			if !(checkContainsUsers(users[i].UserName ,(*updatedUsers))) {
				_,err = stmtInsDelete.Exec(&dashboardAppInfo.AppId, users[i].UserName)
			}
		}
	}else{
		for j := 0; j < countUsers; j++ {
			if !(checkContainsUsers((*updatedUsers)[j].UserName ,users)) {
				_, err = stmtInsAdd.Exec(1, 2, (*updatedUsers)[j].UserName)
			}
		}
	}

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

func UpadateDashboardAppAcls(dashboardAppInfo  *dao.DashboardAppInfo){
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.UPDATE_DB_APP_ACLS)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmtIns.Exec(dashboardAppInfo.Acls, dashboardAppInfo.AppId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
}

func checkContainsGroups(group string, groups []dao.DashboardAppGroup) bool {
	for _, v := range groups {
		if v.GroupName == group {
			return true
		}
	}
	return false
}

func checkContainsMetrics(metricid int, groups []dao.DashboardAppMetric) bool {
	for _, v := range groups {
		if v.MetricId == metricid {
			return true
		}
	}
	return false
}

func checkContainsUsers(username string,users []dao.DashboardAppUser) bool{
	for _, v := range users {
		if v.UserName == username {
			return true
		}
	}
	return false
}