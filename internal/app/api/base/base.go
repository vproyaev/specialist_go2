package base

import (
	"encoding/json"
	"net/http"

	"specialist/internal/models"
	"specialist/internal/utils"
)

func Handler(info *models.Info) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.SetOKResult(w)
		err := json.NewEncoder(w).Encode(info)
		if err != nil {
			return
		}
	}
}
