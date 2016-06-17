package routes

import (
	dashboard_handlers "wislabs.wifi.manager/handlers"
)

var DashboardUserRoutes = Routes{
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
		dashboard_handlers.UpdateUserbyAdminHandler,
	},
	Route{
		"Update user profile",
		"PUT",
		"/dashboard/user",
		true,
		dashboard_handlers.UpdateUser,
	},
	Route{
		"Update Dashboard user",
		"POST",
		"/dashboard/users/changepassword",
		true,
		dashboard_handlers.UpdateUserPasswordHandler,
	},
	Route{
		"Check Dashboard user Exists",
		"GET",
		"/dashboard/checkuser/{username}",
		true,
		dashboard_handlers.UserExistInTenantHandler,
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
		"Get Users in App groups",
		"POST",
		"/dashboard/usersingroups",
		true,
		dashboard_handlers.GetUsersOfGroups,
	},
	Route{
		"Get Users SSIDS",
		"GET",
		"/dashboard/users/{username}/ssids",
		true,
		dashboard_handlers.GetSSIDsOfUser,
	},
	Route{
		"Add user SSIDS",
		"POST",
		"/dashboard/users/{username}/ssids",
		true,
		dashboard_handlers.AddSSIDsOfUser,
	},
	Route{
		"Update user SSIDS",
		"PUT",
		"/dashboard/users/{username}/ssids",
		true,
		dashboard_handlers.UpdateSSIDsOfUser,
	},
	Route{
		"GET usernames allowed for the ssid list",
		"GET",
		"/dashboard/users/ssids",
		true,
		dashboard_handlers.GetUsersOfSSIDs,
	},
}
