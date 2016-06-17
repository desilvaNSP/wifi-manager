package routes

import (
	dashboard_handlers "wislabs.wifi.manager/handlers"
)

var WifiUsageRoutes = Routes{
	Route{
		"Get OS Stats ",
		"POST",
		"/wifi/devices/osstats",
		true,
		dashboard_handlers.GetOSStatsHandler,
	},
	Route{
		"Get Device Stats ",
		"POST",
		"/wifi/devices/devicestats",
		true,
		dashboard_handlers.GetDeviceStatsHandler,
	},
	Route{
		"Get Browser Stats ",
		"POST",
		"/wifi/devices/browserstats",
		true,
		dashboard_handlers.GetBrowserStatsHandler,
	},
	Route{
		"Download Summary Details Dashboard",
		"POST",
		"/wifi/summary/downloadrawdata",
		true,
		dashboard_handlers.DownlaodCSVSummaryDetailsDashboard,
	},
	Route{
		"Get Access Point Details Dashboard",
		"POST",
		"/wifi/summary/accespoint",
		true,
		dashboard_handlers.GetAccessPointAggregatedDataFromToHandler,
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
		"Get daily total WIFI Downloads ",
		"POST",
		"/wifi/usage/dailyavguserdownloads",
		true,
		dashboard_handlers.GetAvgUserDownloadsFromToHandler,
	},
	Route{
		"Get daily AVG WIFI Downloads ",
		"POST",
		"/wifi/usage/dailytotaluploads",
		true,
		dashboard_handlers.GetAgregatedUploadsFromToHandler,
	},
	Route{
		"Get daily AVG user session time ",
		"POST",
		"/wifi/usage/dailyavgusersessiontime",
		true,
		dashboard_handlers.GetAvgUserSessionTimeFromToHandler,
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
		dashboard_handlers.GetTotalSessionCountTimeFromToHandler,
	},
}
