package routes

import (
	"net/http"
	dashboard_handlers "wislabs.wifi.manager/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Secured 	bool
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
		"Create WIFI user",
		"POST",
		"/wifi/users",
		true,
		dashboard_handlers.AddUserHandler,
	},
	Route{
		"GetWiFiUsers",
		"GET",
		"/wifi/{tenantid}/users",
		true,
		dashboard_handlers.GetUsersHandler,
	},
	Route{
		"UpdateWiFiUsers",
		"PUT",
		"/wifi/{tenantid}/users",
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
		dashboard_handlers.GetTotalSessoinCountTimeFromToHandler,
	},
	Route{
		"GetLocations",
		"GET",
		"/wifi/{tenantid}/locations",
		true,
		dashboard_handlers.GetLocations,
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
		dashboard_handlers.AddLocation,
	},
	Route{
		"Get Users By OS ",
		"POST",
		"/wifi/users/bydeviceos",
		true,
		dashboard_handlers.GetUsersByOSHandler,
	},
	Route{
		"Get Users By Device Type ",
		"POST",
		"/wifi/users/bydevicetype",
		true,
		dashboard_handlers.GetUsersByDeviceTypeHandler,
	},
}
