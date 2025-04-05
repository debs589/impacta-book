package handlers

import (
	"api/internal/authentication"
	"api/internal/models"
	"api/internal/utils"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func NewPublicationHandler(service models.PublicationService) *PublicationHandler {
	return &PublicationHandler{service}
}

type PublicationHandler struct {
	service models.PublicationService
}

func (h *PublicationHandler) CreatePublication(w http.ResponseWriter, r *http.Request) {
	userIDToken, err := authentication.ExtractUserID(r)
	if err != nil {
		utils.Error(w, http.StatusUnauthorized, err)
		return
	}

	var publication models.Publication

	if err := json.NewDecoder(r.Body).Decode(&publication); err != nil {
		utils.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	publication.AuthorID = userIDToken

	id, err := h.service.CreatePublication(publication)
	if err != nil {
		if errors.Is(err, utils.ErrInvalidArguments) {
			utils.Error(w, http.StatusBadRequest, err)
			return
		}
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}

	publication.ID = id

	utils.JSON(w, http.StatusCreated, publication)

}

func (h *PublicationHandler) GetPublication(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
	}

	publication, err := h.service.GetPublication(id)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, publication)
}

func (h *PublicationHandler) GetPublications(w http.ResponseWriter, r *http.Request) {
	userIDToken, err := authentication.ExtractUserID(r)
	if err != nil {
		utils.Error(w, http.StatusUnauthorized, err)
		return
	}

	publications, err := h.service.GetPublications(userIDToken)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, publications)
}
