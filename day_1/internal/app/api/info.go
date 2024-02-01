package api

import (
	"encoding/json"
	"net/http"

	"day_1/internal/models"
	"day_1/internal/utils"
)

func InfoHandler(info *models.Info) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.SetOKResult(w)
		err := json.NewEncoder(w).Encode(info)
		if err != nil {
			return
		}
	}
}
