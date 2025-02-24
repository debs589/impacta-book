package handlers

import "net/http"

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login Page"))

}
