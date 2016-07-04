package routes

import (
	dashboard_handlers "wislabs.wifi.manager/handlers"
)

var AuthenticationRoutes = Routes{
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
		"Logout",
		"POST",
		"/dashboard/logout",
		true,
		dashboard_handlers.Logout,
	},
}
