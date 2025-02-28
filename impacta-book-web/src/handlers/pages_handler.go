package handlers

import (
	"impacta-book/src/utils"
	"net/http"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "login.html", nil)

}

func LoadUserRegisterPage(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "user-register.html", nil)
}
