package handlers

import (
	"api/internal/authentication"
	"api/internal/models"
	"api/internal/utils"
	"encoding/json"
	"errors"
	"fmt"
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

func (h *PublicationHandler) UpdatePublication(w http.ResponseWriter, r *http.Request) {
	userIDToken, err := authentication.ExtractUserID(r)
	if err != nil {
		utils.Error(w, http.StatusUnauthorized, err)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	publication, err := h.service.GetPublication(id)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}

	if publication.AuthorID != userIDToken {
		utils.Error(w, http.StatusUnauthorized, fmt.Errorf("It's not possible to edit a publication from another user"))
		return
	}

	var reqBody models.Publication
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		utils.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = h.service.UpdatePublication(id, reqBody)
	if err != nil {
		if errors.Is(err, utils.ErrInvalidArguments) {
			utils.Error(w, http.StatusBadRequest, err)
			return
		}
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	utils.JSON(w, http.StatusNoContent, nil)
}

func (h *PublicationHandler) DeletePublication(w http.ResponseWriter, r *http.Request) {
	userIDToken, err := authentication.ExtractUserID(r)
	if err != nil {
		utils.Error(w, http.StatusUnauthorized, err)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	publication, err := h.service.GetPublication(id)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}

	if publication.AuthorID != userIDToken {
		utils.Error(w, http.StatusUnauthorized, fmt.Errorf("It's not possible to delete a publication from another user"))
		return
	}

	err = h.service.DeletePublication(id)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	utils.JSON(w, http.StatusNoContent, nil)
}

func (h *PublicationHandler) GetPublicationsByUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
	}

	publications, err := h.service.GetPublications(userID)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, publications)
}
