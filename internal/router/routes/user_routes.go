package routes

import (
	"api/internal/handlers"
	"api/internal/middlewares"
	"api/internal/models"
	"github.com/go-chi/chi/v5"
)

func NewUserRoutes(mux *chi.Mux, service models.UserService) error {
	userHandler := handlers.NewUserHandler(service)

	mux.Route("/user", func(router chi.Router) {
		router.Post("/", userHandler.CreateUser)
		router.With(middlewares.Authenticate).Get("/{id}", userHandler.GetUser)
		router.With(middlewares.Authenticate).Get("/", userHandler.GetUsers)
		router.With(middlewares.Authenticate).Put("/{id}", userHandler.UpdateUser)
		router.With(middlewares.Authenticate).Delete("/{id}", userHandler.DeleteUser)
		router.With(middlewares.Authenticate).Post("/{id}/follow", userHandler.FollowUser)
		router.With(middlewares.Authenticate).Post("/{id}/unfollow", userHandler.UnfollowUser)
		router.With(middlewares.Authenticate).Post("/{id}/followers", userHandler.GetFollowers)
		router.With(middlewares.Authenticate).Post("/{id}/following", userHandler.GetFollowing)
	})

	return nil
}
