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
		"Update WiFi Locations Instances",
		"PUT",
		"/wifi/locations",
		true,
		dashboard_handlers.UpdateWiFiLocationHandler,
	},
	Route{
		"Update WiFi APs",
		"PUT",
		"/wifi/updateaps",
		true,
		dashboard_handlers.UpdateAPsHandler,
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
		"Check ssids is exist with the mac",
		"GET",
		"/wifi/locations/{mac}/{ssid}",
		true,
		dashboard_handlers.IsSSIDExist,
	},
	Route{
		"GET aps details associated with the location",
		"GET",
		"/wifi/locations/{mac}",
		true,
		dashboard_handlers.GetAPsByMac,
	},
	Route{
		"GET all MACS",
		"GET",
		"/wifi/macs",
		true,
		dashboard_handlers.GetMACs,
	},
	Route{
		"GET all APS",
		"GET",
		"/wifi/aps",
		true,
		dashboard_handlers.GetAPs,
	},

}
