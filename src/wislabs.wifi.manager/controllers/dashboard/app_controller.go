package dashboard

import (
	"wislabs.wifi.manager/utils"
	"wislabs.wifi.manager/dao"
	"wislabs.wifi.manager/commons"
//	log "github.com/Sirupsen/logrus"
//"database/sql"
	"fmt"
)

func CreateNewDashboardApp(dashboardAppInfo dao.DashboardAppInfo) {
	appId, err := AddDashboardApp(&dashboardAppInfo)
	fmt.Printf("ACL : @")
	fmt.Printf(dashboardAppInfo.Acls)
	if (err == nil) {
		AddDashboardAppUsers(&dashboardAppInfo.Users, appId)
		AddDashboardAppGroups(&dashboardAppInfo.Groups, appId)
		AddDashboardAppMetrics(&dashboardAppInfo.Metrics, appId)
		AddDashboardAppAcls(dashboardAppInfo.Acls,appId)
	}
}

func UpdateDashBoardSettings(dashboardAppInfo dao.DashboardAppInfo) {

	//UpadateDashboardAppUsers(&dashboardAppInfo);
	UpadateDashboardAppGroups(&dashboardAppInfo);
	//UpadateDashboardAppMetrics(&dashboardAppInfo);
	UpadateDashboardAppAcls(&dashboardAppInfo);



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
	fmt.Printf("sasasas @ 2")
	dbMap := utils.GetDBConnection("portal");
	defer dbMap.Db.Close()
	fmt.Printf("sasasas @ 3")
	var aclsTypes []string

	_, err := dbMap.Select(&aclsTypes, commons.GET_ALL_DASHBOARD_ACLS)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf(aclsTypes[0])
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
		_, err = stmtIns.Exec(appId, (*appGroup)[i].GroupName)
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

/*func UpadateDashboardAppUsers(dashboardAppInfo  dao.DashboardAppInfo){
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.UPDATE_DB_APP_USERS)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmtIns.Exec(tenantId, username)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
}*/
/*func UpadateDashboardAppGroups(dashboardAppInfo  dao.DashboardAppInfo){
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var appGroups *[]dao.DashboardAppGroup
	appGroups = dashboardAppInfo.Groups




	stmtIns, err := dbMap.Db.Prepare(commons.UPDATE_DB_APP_GROUPS)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmtIns.Exec(tenantId, username)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
}*/

/*func UpadateDashboardAppMetrics(dashboardAppInfo  dao.DashboardAppInfo){
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.UPDATE_DB_APP_METRICS)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmtIns.Exec(tenantId, username)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
}*/

func UpadateDashboardAppAcls(dashboardAppInfo  *dao.DashboardAppInfo){
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	fmt.Printf("sandun")
	fmt.Printf(dashboardAppInfo.Acls);


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