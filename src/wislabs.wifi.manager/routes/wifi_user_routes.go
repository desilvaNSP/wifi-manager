package routes

import (
	dashboard_handlers "wislabs.wifi.manager/handlers"
)

var WifiUserRoutes = Routes{
	Route{
		"Create WIFI user",
		"POST",
		"/dashboard/wifi/users",
		true,
		dashboard_handlers.AddUserHandler,
	},
	Route{
		"GetWiFiUsers",
		"GET",
		"/dashboard/wifi/users",
		true,
		dashboard_handlers.GetUsersHandler,
	},
	Route{
		"UpdateWiFiUsers",
		"PUT",
		"/dashboard/wifi/users",
		true,
		dashboard_handlers.UpdateUserHandler,
	},
	Route{
		"DeleteWiFiUser",
		"DELETE",
		"/dashboard/wifi/{tenantid}/users/{username}/{groupname}",
		true,
		dashboard_handlers.DeleteUserHandler,
	},
	Route{
		"Check Wifi user Exists In Group",
		"GET",
		"/dashboard/wifi/users/{groupname}/{username}",
		true,
		dashboard_handlers.WifiUserExistInGroupNameHanlder,
	},
	Route{
		"GetUsersCountFromToLocation",
		"POST",
		"/dashboard/wifi/users/count",
		true,
		dashboard_handlers.GetUsersCountFromToHandler,
	},
	Route{
		"Get Returning Users",
		"POST",
		"/dashboard/wifi/users/returncount",
		true,
		dashboard_handlers.GetReturningUsersCountFromToHandler,
	},
	Route{
		"GetUsersCountSeriesFromTo",
		"POST",
		"/dashboard/wifi/users/dailycountseries",
		true,
		dashboard_handlers.GetDailyUsersCountSeriesFromToHandler,
	},
	Route{
		"GetUsersCountFromToLocation",
		"POST",
		"/dashboard/wifi/users/countbydownlods/{threshold}",
		true,
		dashboard_handlers.GetUserCountOfDownloadsOverHandler,
	},
}
