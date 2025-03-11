package handlers

import (
	"api/internal/models"
	"api/internal/utils"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"strings"
)

func NewUserHandler(service models.UserService) *UserHandler {
	return &UserHandler{service}
}

type UserHandler struct {
	service models.UserService
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	id, err := h.service.CreateUser(user, "register")
	if err != nil {
		if errors.Is(err, utils.ErrInvalidArguments) {
			utils.Error(w, http.StatusBadRequest, err)
			return
		}
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}

	user.ID = id

	utils.JSON(w, http.StatusCreated, user)

}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	users, err := h.service.GetUsers(nameOrNick)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, users)

}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, user)

}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	var reqBody models.User
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		utils.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = h.service.UpdateUser(id, reqBody, "update")
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

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	err = h.service.DeleteUser(id)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	utils.JSON(w, http.StatusNoContent, nil)
}
