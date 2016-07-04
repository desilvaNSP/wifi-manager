package routes

import (
	dashboard_handlers "wislabs.wifi.manager/handlers"
)

var WifiUserRoutes = Routes{
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
		"/wifi/{tenantid}/users/{username}/{groupname}",
		true,
		dashboard_handlers.DeleteUserHandler,
	},
	Route{
		"Check Wifi user Exists In Group",
		"GET",
		"/wifi/users/{groupname}/{username}",
		true,
		dashboard_handlers.WifiUserExistInGroupNameHanlder,
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
}
