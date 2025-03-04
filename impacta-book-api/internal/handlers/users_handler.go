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

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var reqBody models.User

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
	}

	utils.JSON(w, http.StatusOK, reqBody)

}
