package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func Error(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}

func JSON(w http.ResponseWriter, code int, body any) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if body != nil {
		if err := json.NewEncoder(w).Encode(body); err != nil {
			log.Fatal(err)
		}
	}

}
