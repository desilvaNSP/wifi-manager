package routes

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Secured     bool
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var ApplicationRoutes Routes

func init() {
	routes := []Routes{
		WifiUserRoutes,
		WifiUsageRoutes,
		WifiLocationRoutes,
		DashboardUserRoutes,
		DashoardAppRoutes,
		RadiusRoutes,
		AuthenticationRoutes,
	}

	for _, r := range routes {
		ApplicationRoutes = append(ApplicationRoutes, r...)
	}
}

