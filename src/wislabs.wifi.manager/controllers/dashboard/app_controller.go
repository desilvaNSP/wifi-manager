package dashboard

import (
	"wislabs.wifi.manager/utils"
	"wislabs.wifi.manager/dao"
	"wislabs.wifi.manager/commons"
	"database/sql"
	log "github.com/Sirupsen/logrus"
)

func CreateNewDashboardApp(dashboardAppInfo dao.DashboardAppInfo) {
	appId, err := AddDashboardApp(&dashboardAppInfo)
	if (err == nil) {
		switch (dashboardAppInfo.FilterCriteria){
		case "groupname" :
			AddDashboardAppGroups(&dashboardAppInfo.Groups, appId)
		case "ssid" :
			AddDashboardAppFilterParams(appId, dashboardAppInfo.Parameters)
		}
		AddDashboardAppUsers(&dashboardAppInfo.Users, appId)
		AddDashboardAppMetrics(&dashboardAppInfo.Metrics, appId)
		AddDashboardAppAcls(dashboardAppInfo.Acls,appId)
	}
}

func UpdateDashBoardAppSettings(dashboardAppInfo dao.DashboardAppInfo) {
	UpdateDashboardAppFilterCriteria(&dashboardAppInfo)
	UpdateDashboardAppUsers(&dashboardAppInfo);
	UpdateDashboardAppGroups(&dashboardAppInfo);
	UpdateDashboardAppMetrics(&dashboardAppInfo);
	UpdateDashboardAppAcls(&dashboardAppInfo);
	UpdateDashboardAppAggregateValue(&dashboardAppInfo);
}

func GetAllDashboardAppsOfUser(username string, tenantId int) []dao.DashboardApp {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var apps []dao.DashboardApp
	_, err := dbMap.Select(&apps, commons.GET_DASHBOARD_USER_APPS, username, tenantId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return apps
}

func GetDashboardUsersInGroups(tenantid int,appGroups []dao.DashboardAppGroup) [][]string{
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	usersInGroups := make([][]string,len(appGroups))
	for i := 0; i < len(appGroups); i++ {

		var users []dao.DashboardAppUser
		_, err := dbMap.Select(&users, commons.GET_DASHBOARD_USERS_IN_GROUP, tenantid,GetApGroupId(tenantid,appGroups[i].GroupName))
		if err != nil {
			checkErr(err,"Error happening while get dashboard users in group") // proper error handling instead of panic in your app
		}
		usersInGroup := make([]string,len(users))
		for j := 0; j < len(users); j++ {
			usersInGroup[j] = (users[j].UserName)
		}
		usersInGroups[i]= usersInGroup
	}
	return  usersInGroups
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

func GetDashboardMetricsOfApp(appId int) []dao.DashboardAppMetric {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var metrics []dao.DashboardAppMetric
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

	var acls dao.DashboardAppAcls
	err := dbMap.SelectOne(&acls, commons.GET_DASHBOARD_APP_ACLS, appId)
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		return acls.Acls
	}
	return acls.Acls
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

func GetFilterParamsOfApp(appId int) []string{
	dbMap := utils.GetDBConnection(commons.DASHBOARD_DB);
	defer dbMap.Db.Close()

	var filterParams []string
	_, err := dbMap.Select(&filterParams, commons.GET_DASHBOARD_APP_FILTER_PARAMS, appId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return filterParams
}

func GetFilterCriteriaOfApp(appId int) string{
	dbMap := utils.GetDBConnection(commons.DASHBOARD_DB);
	defer dbMap.Db.Close()

	var filterCriteria sql.NullString
	err := dbMap.SelectOne(&filterCriteria, commons.GET_DASHBOARD_APP_CRITERIA, appId)
	if err != nil {
		log.Error(err.Error())
	}
	if filterCriteria.Valid{
		return filterCriteria.String
	}
	return "none"
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
	dbMap := utils.GetDBConnection(commons.DASHBOARD_DB);
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.ADD_DASHBOARD_APP)
	defer stmtIns.Close()

	if err != nil {
		return 0, err
	}
	result, err := stmtIns.Exec(app.TenantId, app.Name, app.Aggregate, app.FilterCriteria)
	if err != nil {
		return 0, err
	} else {
		id, err := result.LastInsertId()
		return id, err
	}
}

func AddDashboardAppFilterParams(appId int64, filterParams []string) error {
	dbMap := utils.GetDBConnection(commons.DASHBOARD_DB);
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.ADD_DASHBOARD_APP_FILTER_PARAMS)
	defer stmtIns.Close()

	if err != nil {
		return err
	}
	for i := 0; i < len(filterParams); i++ {
		_, err = stmtIns.Exec(appId, filterParams[i])
	}
	return err
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
		_, err = stmtIns.Exec(appId,(*appGroup)[i].GroupName)
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

func UpdateDashboardAppAggregateValue(dashboardAppInfo  *dao.DashboardAppInfo){
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

func UpdateDashboardAppFilterCriteria(dashboardAppInfo  *dao.DashboardAppInfo){
	dbMap := utils.GetDBConnection(commons.DASHBOARD_DB);
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.UPDATE_DB_APP_FILTER_CRITERIA)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmtIns.Exec(dashboardAppInfo.FilterCriteria, dashboardAppInfo.AppId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
}

func UpdateDashboardAppGroups(dashboardAppInfo  *dao.DashboardAppInfo){
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

func UpdateDashboardAppMetrics(dashboardAppInfo  *dao.DashboardAppInfo){
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

func UpdateDashboardAppUsers(dashboardAppInfo  *dao.DashboardAppInfo){
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
	var countUsers = len(users)
	if( Length <= countUsers){
		for i := 0; i < countUsers; i++ {
			if !(checkContainsUsers(users[i].UserName ,(*updatedUsers))) {
				_,err = stmtInsDelete.Exec(dashboardAppInfo.AppId, users[i].UserName)
			}
		}
	}else{
		for j := 0; j < Length; j++ {
			if !(checkContainsUsers((*updatedUsers)[j].UserName ,users)) {
				_, err = stmtInsAdd.Exec(dashboardAppInfo.TenantId,dashboardAppInfo.AppId, (*updatedUsers)[j].UserName)
			}
		}
	}

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

func UpdateDashboardAppAcls(dashboardAppInfo  *dao.DashboardAppInfo){
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