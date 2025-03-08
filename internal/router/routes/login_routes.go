package routes

import (
	"api/internal/handlers"
	"api/internal/models"
	"github.com/go-chi/chi/v5"
)

func NewLoginRoutes(mux *chi.Mux, service models.UserService) error {
	loginHandler := handlers.NewLoginHandler(service)

	mux.Route("/login", func(router chi.Router) {
		router.Post("/", loginHandler.Login)
	})

	return nil

}
