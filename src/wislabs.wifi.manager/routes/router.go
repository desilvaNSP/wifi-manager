package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"wislabs.wifi.manager/authenticator"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	for _, route := range ApplicationRoutes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
        if(route.Secured) {
			handler = authenticator.RequireTokenAuthentication(handler)
		}
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
		 	Handler(handler)
	}

	return router
}
