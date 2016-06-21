package routes

import (
	dashboard_handlers "wislabs.wifi.manager/handlers"
)

var RadiusRoutes = Routes{
	Route{
		"Check Wifi user Valid In Radius",
		"GET",
		"/radius/users/{username}",
		true,
		dashboard_handlers.WifiUserValidInRadiusHanlder,
	},
	Route{
		"Radius Server Authentication Testing",
		"POST",
		"/radius/user/connection",
		true,
		dashboard_handlers.TestRadiusAuthConnection,
	},
	Route{
		"Create Radius Server",
		"POST",
		"/radius/server",
		true,
		dashboard_handlers.CreateRadiusServerHandler,
	},
	Route{
		"Updating Radius Server",
		"PUT",
		"/radius/server",
		true,
		dashboard_handlers.UpdateRadiusServerHandler,
	},
	Route{
		"Get Radius Details of User",
		"GET",
		"/radius/radiusdetails",
		true,
		dashboard_handlers.GetRadiusServerDetailsHandler,
	},
	Route{
		"Check Is NAS Client exists in Radius",
		"GET",
		"/radius/{instanceid}/validnas",
		true,
		dashboard_handlers.NASIpExistInRadiusHandler,
	},
	Route{
		"Delete Radius Instance",
		"DELETE",
		"/radius/{instanceid}",
		true,
		dashboard_handlers.DeleteRadiusInstanceHandler,
	},
	Route{
		"Get Clients in radius server",
		"GET",
		"/radius/{instanceid}/clients",
		true,
		dashboard_handlers.GetRadiusServerClientsHanlder,
	},
	Route{
		"Add Clients in radius server",
		"POST",
		"/radius/server/client",
		true,
		dashboard_handlers.CreateNASClientHanlder,
	},
	Route{
		"Updating NAS Clients in radius server",
		"PUT",
		"/radius/server/client",
		true,
		dashboard_handlers.UpdateNASClientHanlder,
	},
	Route{
		"Updating NAS Clients in radius server",
		"DELETE",
		"/radius/{instanceid}/{nasclientid}",
		true,
		dashboard_handlers.DeleteNASClientHanlder,
	},
}