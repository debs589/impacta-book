package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	URI                string
	Method             string
	Function           func(http.ResponseWriter, *http.Request)
	NeedAuthentication bool
}

func Configure(router *mux.Router) *mux.Router {
	routes := routesLogin

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return router
}
