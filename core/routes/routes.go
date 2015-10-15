package routes

import (
	"net/http"
	dashboard_handlers "wifi-manager/core/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetLoginPage",
		"GET",
		"/dashboard/login",
		dashboard_handlers.GetLoginPage,
	},
	Route{
		"Login",
		"POST",
		"/dashboard/login",
		dashboard_handlers.Login,
	},
	Route{
		"Register Dashboard user",
		"POST",
		"/dashboard/register",
		dashboard_handlers.RegisterUser,
	},
	Route{
		"Get Dashboard user Info",
		"GET",
		"/dashboard/users/{username}",
		dashboard_handlers.GetUserInfo,
	},
	Route{
		"Create WIFI user",
		"POST",
		"/wifi/users",
		dashboard_handlers.AddUserHandler,
	},
	Route{
		"GetWiFiUsers",
		"GET",
		"/wifi/users",
		dashboard_handlers.GetUsersHandler,
	},
	Route{
		"DeleteWiFiUser",
		"DELETE",
		"/wifi/users/{username}",
		dashboard_handlers.DeleteUserHandler,
	},
	Route{
		"GetUsersCountFromToLocation",
		"POST",
		"/wifi/users/count",
		dashboard_handlers.GetUsersCountFromToHandler,
	},
	Route{
		"Get Returning Users",
		"POST",
		"/wifi/users/returncount",
		 dashboard_handlers.GetReturningUsersCountFromToHandler,
	},
	Route{
		"GetUsersCountFromToLocation",
		"POST",
		"/wifi/users/countbydownlods/{threshold}",
		dashboard_handlers.GetUserCountOfDownloadsOverHandler,
	},
	Route{
		"Get WIFI Downloads",
		"POST",
		"/wifi/usage/downloads",
		dashboard_handlers.GetDownloadsFromToHandler,
	},
	Route{
		"Get daily total WIFI Downloads ",
		"POST",
		"/wifi/usage/dailytotaldownloads",
		dashboard_handlers.GetAgregatedDownloadsFromToHandler,
	},
	Route{
		"Get WIFI Uploads",
		"POST",
		"/wifi/usage/uploads",
		dashboard_handlers.GetUploadsFromToHandler,
	},
	Route{
		"Get Avg Session Time",
		"POST",
		"/wifi/sessions/avg",
		dashboard_handlers.GetAvgSessoinTimeFromToHandler,
	},
	Route{
		"Get Total Session Count",
		"POST",
		"/wifi/sessions/count",
		dashboard_handlers.GetTotalSessoinCountTimeFromToHandler,
	},
	Route{
		"GetLocations",
		"GET",
		"/wifi/locations",
		dashboard_handlers.GetLocations,
	},
	Route{
		"GetLocationAccessPoints",
		"GET",
		"/wifi/locations/{locationid}",
		dashboard_handlers.GetUsersHandler,
	},
	Route{
		"DeleteLocationAccessPoints",
		"DELETE",
		"/wifi/locations/{locationid}/{mac}",
		dashboard_handlers.DeleteLocation,
	},
	Route{
		"AddLocationAccessPoints",
		"POST",
		"/wifi/locations/{locationid}/{mac}",
		dashboard_handlers.AddLocation,
	},
	Route{
		"Get Users By OS ",
		"POST",
		"/wifi/users/bydeviceos",
		dashboard_handlers.GetUsersByOSHandler,
	},
	Route{
		"Get Users By Device Type ",
		"POST",
		"/wifi/users/bydevicetype",
		dashboard_handlers.GetUsersByDeviceTypeHandler,
	},
}
