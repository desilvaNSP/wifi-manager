package routes

import (
	dashboard_handlers "wislabs.wifi.manager/handlers"
)

var WifiAPRoutes = Routes{
	Route{
		"GetCountActiveInactiveAccessPoints",
		"GET",
		"/wifi/ap/activecount",
		true,
		dashboard_handlers.GetActiveAPHandler,
	},
	Route{
		"GetCountActiveInactiveAccessPoints",
		"GET",
		"/wifi/ap/inactivecount",
		true,
		dashboard_handlers.GetInactiveAPHandler,
	},
	Route{
		"GetCountActiveInactiveAccessPoints",
		"GET",
		"/wifi/ap/distinctcount",
		true,
		dashboard_handlers.GetDistinctMacCountHandler,
	},
}
