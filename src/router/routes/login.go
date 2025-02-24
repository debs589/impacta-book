package routes

import (
	"impacta-book/src/handlers"
	"net/http"
)

var routesLogin = []Route{
	{
		URI:                "/",
		Method:             http.MethodGet,
		Function:           handlers.LoadLoginPage,
		NeedAuthentication: false,
	},
	{
		URI:                "/login",
		Method:             http.MethodGet,
		Function:           handlers.LoadLoginPage,
		NeedAuthentication: false,
	},
}
