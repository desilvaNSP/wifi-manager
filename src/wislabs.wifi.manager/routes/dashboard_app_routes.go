package routes

import (
	dashboard_handlers "wislabs.wifi.manager/handlers"
)

var DashoardAppRoutes = Routes{
	Route{
		"Get Dashboard App Users ",
		"GET",
		"/dashboard/apps/{appid}/users",
		true,
		dashboard_handlers.GetUsersOfApp,
	},
	Route{
		"Get Dashboard App Metrics",
		"GET",
		"/dashboard/apps/{appid}/metrics",
		true,
		dashboard_handlers.GetMetricsOfApp,
	},
	Route{
		"Get Dashboard App Groups",
		"GET",
		"/dashboard/apps/{appid}/groups",
		true,
		dashboard_handlers.GetGroupsOfApp,
	},
	Route{
		"Get Dashboard App Acl",
		"GET",
		"/dashboard/apps/{appid}/acl",
		true,
		dashboard_handlers.GetAclsOfApp,
	},
	Route{
		"Get Dashboard App Acl",
		"GET",
		"/dashboard/apps/{appid}/aggregate",
		true,
		dashboard_handlers.GetAggregateValueOfApp,
	},
	Route{
		"Get Dashboard App Filter parameters",
		"GET",
		"/dashboard/apps/{appid}/filterparameters",
		true,
		dashboard_handlers.GetAppFilterParameters,
	},
	Route{
		"Get All Dashboard app settings",
		"GET",
		"/dashboard/apps/{appid}/appsettings",
		true,
		dashboard_handlers.GetDashboardAppSettings,
	},
	Route{
		"Delete Dashboard User App",
		"DELETE",
		"/dashboard/{tenantid}/apps/{appid}",
		true,
		dashboard_handlers.DeleteDashboardApp,
	},
	Route{
		"Add Dashboard User App",
		"POST",
		"/dashboard/apps",
		true,
		dashboard_handlers.CreateDashboardApp,
	},
	Route{
		"Update Dashboard User App Settings",
		"PUT",
		"/dashboard/apps",
		true,
		dashboard_handlers.UpdateDashBoardSettingsHandler,
	},
	Route{
		"Get All Dashboard Metrics ",
		"GET",
		"/dashboard/{tenantid}/metrics",
		true,
		dashboard_handlers.GetAllDashboardMetrics,
	},
	Route{
		"Get All Dashboard ACL types ",
		"GET",
		"/dashboard/acltypes",
		true,
		dashboard_handlers.GetAclTypes,
	},
}
