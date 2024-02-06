package task_store

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"specialist/internal/db"
	"specialist/internal/models"
	"specialist/internal/utils"
)

func GetDueTasksHandler(w http.ResponseWriter, r *http.Request) {
	utils.SetOKResult(w)

	connector, ok := r.Context().Value("connector").(*db.Connector)
	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	year, yearOK := vars["year"]
	month, _ := vars["month"]
	day, _ := vars["day"]

	if !yearOK {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	yearINT, err := strconv.Atoi(year)
	if err != nil {
		http.Error(w, "Invalid year", http.StatusBadRequest)
		return
	}

	var monthINT *int
	if month != "" {
		monthVal, err := strconv.Atoi(month)
		if err != nil {
			http.Error(w, "Invalid month", http.StatusBadRequest)
			return
		}
		monthINT = &monthVal
	}

	var dayINT *int
	if day != "" {
		dayVal, err := strconv.Atoi(day)
		if err != nil {
			http.Error(w, "Invalid day", http.StatusBadRequest)
			return
		}
		dayINT = &dayVal
	}

	result, err := connector.GetDueTasks(yearINT, monthINT, dayINT)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(
		models.APIResponse{Result: result},
	)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
