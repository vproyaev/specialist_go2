package utils

import (
	"net/http"
)

func SetOKResult(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
