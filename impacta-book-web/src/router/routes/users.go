package routes

import (
	"impacta-book/src/handlers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:                    "/create-user",
		Method:                 http.MethodGet,
		Function:               handlers.LoadUserRegisterPage,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/user",
		Method:                 http.MethodPost,
		Function:               handlers.CreateUser,
		AuthenticationRequired: false,
	},
}
