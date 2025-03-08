package handlers

import (
	"api/internal/authentication"
	"api/internal/models"
	"api/internal/security"
	"api/internal/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func NewLoginHandler(service models.UserService) *LoginHandler {
	return (*LoginHandler)(&UserHandler{service})
}

type LoginHandler struct {
	service models.UserService
}

func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	userSaveAtDb, err := h.service.GetUserByEmail(user.Email)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.CheckPassword(userSaveAtDb.Password, user.Password); err != nil {
		utils.Error(w, http.StatusUnauthorized, err)
		return
	}
	userID := strconv.Itoa(user.ID)

	token, err := authentication.CreateToken(userSaveAtDb.ID)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, models.AuthenticationData{ID: userID, Token: token})

}
