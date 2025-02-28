package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	AuthenticationRequired bool
}

func Configure(router *mux.Router) *mux.Router {
	routes := routesLogin
	routes = append(routes, userRoutes...)

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
