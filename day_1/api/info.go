package api

import (
	"encoding/json"
	"net/http"

	"day_1/models"
)

func InfoHandler(info *models.Info) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(info)
		if err != nil {
			return
		}
	}
}
