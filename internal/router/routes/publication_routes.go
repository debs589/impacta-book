package routes

import (
	"api/internal/handlers"
	"api/internal/middlewares"
	"api/internal/models"
	"github.com/go-chi/chi/v5"
)

func NewPublicationRoutes(mux *chi.Mux, service models.PublicationService) error {
	publicationHandler := handlers.NewPublicationHandler(service)

	mux.Route("/publication", func(router chi.Router) {
		router.With(middlewares.Authenticate).Post("/", publicationHandler.CreatePublication)
		router.With(middlewares.Authenticate).Get("/{id}", publicationHandler.GetPublication)
		router.With(middlewares.Authenticate).Get("/", publicationHandler.GetPublications)
		router.With(middlewares.Authenticate).Put("/{id}", publicationHandler.UpdatePublication)
		router.With(middlewares.Authenticate).Delete("/{id}", publicationHandler.DeletePublication)
		router.With(middlewares.Authenticate).Get("/{user_id}/publication", publicationHandler.GetPublicationsByUser)
		router.With(middlewares.Authenticate).Post("/{id}/like", publicationHandler.LikePublication)
		router.With(middlewares.Authenticate).Post("/{id}/unlike", publicationHandler.UnlikePublication)
	})

	return nil
}
