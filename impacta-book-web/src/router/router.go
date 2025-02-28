package router

import (
	"github.com/gorilla/mux"
	"impacta-book/src/router/routes"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
