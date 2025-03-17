package middlewares

import (
	"api/internal/authentication"
	"api/internal/utils"
	"net/http"
)

func Authenticate(nextFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			utils.Error(w, http.StatusUnauthorized, err)
			return
		}
		nextFunc.ServeHTTP(w, r)
	})
}
