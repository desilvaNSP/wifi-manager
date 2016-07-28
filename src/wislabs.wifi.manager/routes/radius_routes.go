package routes

import (
	dashboard_handlers "wislabs.wifi.manager/handlers"
)

var RadiusRoutes = Routes{
	Route{
		"Check Wifi user Valid In Radius",
		"GET",
		"/dashboard/radius/users/{username}",
		true,
		dashboard_handlers.WifiUserValidInRadiusHanlder,
	},
	Route{
		"Radius Server Authentication Testing",
		"POST",
		"/dashboard/radius/user/connection",
		true,
		dashboard_handlers.TestRadiusAuthConnection,
	},
	Route{
		"Create Radius Server",
		"POST",
		"/dashboard/radius/server",
		true,
		dashboard_handlers.CreateRadiusServerHandler,
	},
	Route{
		"Updating Radius Server",
		"PUT",
		"/dashboard/radius/server",
		true,
		dashboard_handlers.UpdateRadiusServerHandler,
	},
	Route{
		"Get Radius Details of User",
		"GET",
		"/dashboard/radius/radiusdetails",
		true,
		dashboard_handlers.GetRadiusServerDetailsHandler,
	},
	Route{
		"Check Is NAS Client exists in Radius",
		"GET",
		"/dashboard/radius/{instanceid}/validnas",
		true,
		dashboard_handlers.NASIpExistInRadiusHandler,
	},
	Route{
		"Delete Radius Instance",
		"DELETE",
		"/dashboard/radius/{instanceid}",
		true,
		dashboard_handlers.DeleteRadiusInstanceHandler,
	},
	Route{
		"Get Clients in radius server",
		"GET",
		"/dashboard/radius/{instanceid}/clients",
		true,
		dashboard_handlers.GetRadiusServerClientsHanlder,
	},
	Route{
		"Add Clients in radius server",
		"POST",
		"/dashboard/radius/server/client",
		true,
		dashboard_handlers.CreateNASClientHanlder,
	},
	Route{
		"Updating NAS Clients in radius server",
		"PUT",
		"/dashboard/radius/server/client",
		true,
		dashboard_handlers.UpdateNASClientHanlder,
	},
	Route{
		"Updating NAS Clients in radius server",
		"DELETE",
		"/dashboard/radius/{instanceid}/{nasclientid}",
		true,
		dashboard_handlers.DeleteNASClientHanlder,
	},
}