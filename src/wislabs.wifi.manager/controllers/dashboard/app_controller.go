package dashboard

import (
	"wislabs.wifi.manager/utils"
	"wislabs.wifi.manager/dao"
	"wislabs.wifi.manager/common"
//	log "github.com/Sirupsen/logrus"
	//"database/sql"
)

func CreateNewDashboardApp(dashboardAppInfo dao.DashboardAppInfo) {
	AddDashboardApp(&dashboardAppInfo)
}

func GetAllDashboardAppsOfUser(username string, tenantId int) []dao.DashboardApp {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var apps []dao.DashboardApp
	_, err := dbMap.Select(&apps, common.GET_DASHBOARD_USER_APPS,username, tenantId)
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		return apps
	}
	return apps
}

func GetAllDashboardUsersOfApp(appId int) []dao.DashboardAppUser {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var users []dao.DashboardAppUser
	_, err := dbMap.Select(&users, common.GET_DASHBOARD_APP_USERS, appId)
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		return users
	}
	return users
}

func GetDashboardAppMetrics(appId int) []dao.DashboardMetric {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var metrics []dao.DashboardMetric
	_, err := dbMap.Select(&metrics, common.GET_DASHBOARD_APP_METRICS, appId)
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		return metrics
	}
	return metrics
}

func GetAllDashboardMetrics(tenantId int) []dao.DashboardMetric {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var metrics []dao.DashboardMetric
	_, err := dbMap.Select(&metrics, common.GET_ALL_DASHBOARD_METRICS, tenantId)
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		return metrics
	}
	return metrics
}

func GetDashboardAppGroups(appId int) []dao.DashboardAppGroup {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var groups []dao.DashboardAppGroup
	_, err := dbMap.Select(&groups, common.GET_DASHBOARD_APP_GROUPS, appId)
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		return groups
	}
	return groups
}

func AddDashboardApp(app *dao.DashboardAppInfo) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(common.ADD_DASHBOARD_APP)
	defer stmtIns.Close()

	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(app.TenantId, app.Name)
	return err
}

func AddDashboardAppMetric(appMetric *dao.DashboardAppMetric) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(common.ADD_DASHBOARD_APP_METRIC)
	defer stmtIns.Close()

	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(appMetric.AppId, appMetric.MetricId)
	return err
}

func AddDashboardAppGroup(appGroup *dao.DashboardAppGroup) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(common.ADD_DASHBOARD_APP_GROUP)
	defer stmtIns.Close()

	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(appGroup.AppId, appGroup.GroupName)
	return err
}

func AddDashboardAppUser(appUser *dao.DashboardAppUser) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(common.ADD_DASHBOARD_APP_USER)
	defer stmtIns.Close()

	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(appUser.TenantId, appUser.AppId, appUser.UserName)
	return err
}

func DeleteDashboardApp(appId int, tenantId int) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(common.DELETE_DASHBOARD_APP)
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
