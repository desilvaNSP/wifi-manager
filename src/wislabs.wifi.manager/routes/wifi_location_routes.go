package routes

import (
	dashboard_handlers "wislabs.wifi.manager/handlers"
)

var WifiLocationRoutes = Routes{
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
		"GET ssids associated with the locations",
		"GET",
		"/wifi/locations/ssids",
		true,
		dashboard_handlers.GetSSIDsOfAPGroups,
	},
	Route{
		"GetCountActiveInactiveAccessPoints",
		"GET",
		"/wifi/locations/activeapcounts",
		true,
		dashboard_handlers.GetActiveAPHandler,
	},
	Route{
		"GetCountActiveInactiveAccessPoints",
		"GET",
		"/wifi/locations/inactiveapcounts",
		true,
		dashboard_handlers.GetInactiveAPHandler,
	},
	Route{
		"GetCountActiveInactiveAccessPoints",
		"GET",
		"/wifi/locations/distinctmaccount",
		true,
		dashboard_handlers.GetDistinctMacCountHandler,
	},
}
