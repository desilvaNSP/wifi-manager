package routes

import (
	dashboard_handlers "wislabs.wifi.manager/handlers"
)

var WifiAPRoutes = Routes{
	Route{
		"GetCountActiveInactiveAccessPoints",
		"GET",
		"/dashboard/wifi/ap/activecount",
		true,
		dashboard_handlers.GetActiveAPHandler,
	},
	Route{
		"GetCountActiveInactiveAccessPoints",
		"GET",
		"/dashboard/wifi/ap/inactivecount",
		true,
		dashboard_handlers.GetInactiveAPHandler,
	},
	Route{
		"GetCountActiveInactiveAccessPoints",
		"GET",
		"/dashboard/wifi/ap/distinctcount",
		true,
		dashboard_handlers.GetDistinctMacCountHandler,
	},
}
