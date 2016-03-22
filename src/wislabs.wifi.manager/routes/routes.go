package routes

import (
	"net/http"
	dashboard_handlers "wislabs.wifi.manager/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Secured     bool
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Login",
		"POST",
		"/dashboard/login",
		false,
		dashboard_handlers.Login,
	},
	Route{
		"GetLoginPage",
		"GET",
		"/dashboard/login",
		false,
		dashboard_handlers.GetLoginPage,
	},
	Route{
		"GetRegistrations",
		"GET",
		"/dashboard/register",
		false,
		dashboard_handlers.GetRegistrationPage,
	},
	Route{
		"Get dashboard users",
		"GET",
		"/dashboard/{tenantid}/users",
		true,
		dashboard_handlers.GetDashboardUsersHandler,
	},
	Route{
		"Get dashboard user roles",
		"GET",
		"/dashboard/{tenantid}/roles",
		true,
		dashboard_handlers.GetTenantRolesHandler,
	},
	Route{
		"Get all dashboard user permissions",
		"GET",
		"/dashboard/{tenantid}/permissions",
		true,
		dashboard_handlers.GetAllUserPermissionsHandler,
	},
	Route{
		"Login",
		"POST",
		"/dashboard/login2",
		false,
		dashboard_handlers.AuthenticateUser,
	},
	Route{
		"Register Dashboard user",
		"POST",
		"/dashboard/users",
		false,
		dashboard_handlers.RegisterUser,
	},
	Route{
		"Update Dashboard user",
		"PUT",
		"/dashboard/users",
		true,
		dashboard_handlers.UpdateUser,
	},
	Route{
		"Update user profile",
		"POST",
		"/dashboard/userupdate",
		true,
		dashboard_handlers.UpdateUserProfile,
	},
	Route{
		"Update Dashboard user",
		"POST",
		"/dashboard/userprofile/changepassword",
		true,
		dashboard_handlers.UpdateUserPasswordHandler,
	},
	Route{
		"Get Dashboard user Info",
		"GET",
		"/dashboard/{tenantid}/users/{username}",
		true,
		dashboard_handlers.GetUserProfile,
	},
	Route{
		"Delete Dashboard user",
		"DELETE",
		"/dashboard/{tenantid}/users/{username}",
		true,
		dashboard_handlers.DeleteDashboardUsersHandler,
	},
	Route{
		"Get Dashboard User Apps",
		"GET",
		"/dashboard/{tenantid}/apps/{username}",
		true,
		dashboard_handlers.GetAppsOfUser,
	},
	Route{
		"Get All Dashboard Metrics ",
		"GET",
		"/dashboard/{tenantid}/metrics",
		true,
		dashboard_handlers.GetAllDashboardMetrics,
	},
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
		"Create WIFI user",
		"POST",
		"/wifi/users",
		true,
		dashboard_handlers.AddUserHandler,
	},
	Route{
		"GetWiFiUsers",
		"GET",
		"/wifi/users",
		true,
		dashboard_handlers.GetUsersHandler,
	},
	Route{
		"UpdateWiFiUsers",
		"PUT",
		"/wifi/users",
		true,
		dashboard_handlers.UpdateUserHandler,
	},
	Route{
		"DeleteWiFiUser",
		"DELETE",
		"/wifi/{tenantid}/users/{username}",
		true,
		dashboard_handlers.DeleteUserHandler,
	},
	Route{
		"GetUsersCountFromToLocation",
		"POST",
		"/wifi/users/count",
		true,
		dashboard_handlers.GetUsersCountFromToHandler,
	},
	Route{
		"Get Returning Users",
		"POST",
		"/wifi/users/returncount",
		true,
		dashboard_handlers.GetReturningUsersCountFromToHandler,
	},
	Route{
		"GetUsersCountSeriesFromTo",
		"POST",
		"/wifi/users/dailycountseries",
		true,
		dashboard_handlers.GetDailyUsersCountSeriesFromToHandler,
	},
	Route{
		"GetUsersCountFromToLocation",
		"POST",
		"/wifi/users/countbydownlods/{threshold}",
		true,
		dashboard_handlers.GetUserCountOfDownloadsOverHandler,
	},
	Route{
		"Get WIFI Downloads",
		"POST",
		"/wifi/usage/downloads",
		true,
		dashboard_handlers.GetDownloadsFromToHandler,
	},
	Route{
		"Get daily total WIFI Downloads ",
		"POST",
		"/wifi/usage/dailytotaldownloads",
		true,
		dashboard_handlers.GetAgregatedDownloadsFromToHandler,
	},
	Route{
		"Get daily total WIFI Downloads ",
		"POST",
		"/wifi/usage/dailyavguserdownloads",
		true,
		dashboard_handlers.GetAvgUserDownloadsFromToHandler,
	},
	Route{
		"Get daily AVG WIFI Downloads ",
		"POST",
		"/wifi/usage/dailytotaluploads",
		true,
		dashboard_handlers.GetAgregatedUploadsFromToHandler,
	},
	Route{
		"Get daily AVG user session time ",
		"POST",
		"/wifi/usage/dailyavgusersessiontime",
		true,
		dashboard_handlers.GetAvgUserSessionTimeFromToHandler,
	},
	Route{
		"Get WIFI Uploads",
		"POST",
		"/wifi/usage/uploads",
		true,
		dashboard_handlers.GetUploadsFromToHandler,
	},
	Route{
		"Get Avg Session Time",
		"POST",
		"/wifi/sessions/avg",
		true,
		dashboard_handlers.GetAvgSessoinTimeFromToHandler,
	},
	Route{
		"Get Total Session Count",
		"POST",
		"/wifi/sessions/count",
		true,
		dashboard_handlers.GetTotalSessionCountTimeFromToHandler,
	},
	Route{
		"GetLocations",
		"GET",
		"/wifi/{tenantid}/locations",
		true,
		dashboard_handlers.GetLocations,
	},
	Route{
		"GetLocationGroups",
		"GET",
		"/wifi/{tenantid}/locations/groups",
		true,
		dashboard_handlers.GetLocationGroups,
	},
	Route{
		"GetLocationAccessPoints",
		"GET",
		"/wifi/{tenantid}/locations/{locationid}",
		true,
		dashboard_handlers.GetUsersHandler,
	},
	Route{
		"DeleteLocationAccessPoints",
		"DELETE",
		"/wifi/{tenantid}/locations/{mac}/{ssid}/{groupname}",
		true,
		dashboard_handlers.DeleteLocation,
	},
	Route{
		"DeleteLocationAccessPoints",
		"DELETE",
		"/wifi/{tenantid}/locations/{mac}",
		true,
		dashboard_handlers.DeleteAccessPoint,
	},
	Route{
		"AddLocationAccessPoints",
		"POST",
		"/wifi/locations",
		true,
		dashboard_handlers.AddWiFiLocationHandler,
	},
	Route{
		"Update WiFi Locations AccessPoints",
		"POST",
		"/wifi/locationsupdate",
		true,
		dashboard_handlers.UpdateWiFiLocationHandler,
	},
	Route{
		"Add WiFi Group",
		"POST",
		"/wifi/locations/groups",
		true,
		dashboard_handlers.AddWiFiGroupHandler,
	},
	Route{
		"Get OS Stats ",
		"POST",
		"/wifi/devices/osstats",
		true,
		dashboard_handlers.GetOSStatsHandler,
	},
	Route{
		"Get Device Stats ",
		"POST",
		"/wifi/devices/devicestats",
		true,
		dashboard_handlers.GetDeviceStatsHandler,
	},
	Route{
		"Get Browser Stats ",
		"POST",
		"/wifi/devices/browserstats",
		true,
		dashboard_handlers.GetBrowserStatsHandler,
	},
	Route{
		"Download Summary Details Dashboard",
		"POST",
		"/wifi/summary/downloadrawdata",
		true,
		dashboard_handlers.DownlaodCSVSummaryDetailsDashboard,
	},
	Route{
		"Upload files wifi manager",
		"POST",
		"/dashboard/upload",
		true,
		dashboard_handlers.UploadFilesHandler,
	},
}
