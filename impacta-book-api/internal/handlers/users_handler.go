package handlers

import (
	"api/internal/models"
	"api/internal/utils"
	"encoding/json"
	"net/http"
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
		utils.Error(w, http.StatusBadRequest, err.Error())
	}

	id, err := h.service.CreateUser(user)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err.Error())
	}

	user.ID = id

	utils.JSON(w, http.StatusOK, user)

}
