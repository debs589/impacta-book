package routes

import (
	"api/src/handlers"
	"net/http"
)

var usersRoutes = []Route{
	{
		Uri:                    "/users",
		Method:                 http.MethodPost,
		Function:               handlers.CreateUser,
		AuthenticationRequired: false,
	},
	{
		Uri:                    "/users",
		Method:                 http.MethodGet,
		Function:               func(writer http.ResponseWriter, request *http.Request) {},
		AuthenticationRequired: false,
	},
	{
		Uri:                    "/users/{id}",
		Method:                 http.MethodGet,
		Function:               func(writer http.ResponseWriter, request *http.Request) {},
		AuthenticationRequired: false,
	},
	{
		Uri:                    "/users/{id}",
		Method:                 http.MethodPut,
		Function:               func(writer http.ResponseWriter, request *http.Request) {},
		AuthenticationRequired: false,
	},
	{
		Uri:                    "/users/{id}",
		Method:                 http.MethodDelete,
		Function:               func(writer http.ResponseWriter, request *http.Request) {},
		AuthenticationRequired: false,
	},
}
