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

	})

	return nil
}
