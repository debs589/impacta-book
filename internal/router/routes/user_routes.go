package routes

import (
	"api/internal/handlers"
	"api/internal/models"
	"github.com/go-chi/chi/v5"
)

func NewUserRoutes(mux *chi.Mux, service models.UserService) error {
	userHandler := handlers.NewUserHandler(service)

	mux.Route("/user", func(router chi.Router) {
		router.Post("/", userHandler.CreateUser)
		router.Get("/user", userHandler.GetUsers)
		router.Get("/{id}", userHandler.GetUser)
	})

	return nil

}
