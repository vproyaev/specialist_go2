package task_store

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"specialist/internal/db"
	"specialist/internal/models"
	"specialist/internal/utils"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	utils.SetOKResult(w)

	connector, ok := r.Context().Value("connector").(*db.Connector)
	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	result, err := connector.GetTasks()
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
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	utils.SetOKResult(w)

	connector, ok := r.Context().Value("connector").(*db.Connector)
	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	taskID, ok := vars["task_id"]
	taskError := ""

	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	result, err := connector.GetTask(taskID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	} else if result.IsNull() {
		taskError = "Task not found"
	}

	if taskError != "" {
		err = json.NewEncoder(w).Encode(
			models.APIResponse{Err: taskError},
		)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	} else {
		err = json.NewEncoder(w).Encode(
			models.APIResponse{Result: result},
		)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
}

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	utils.SetPostOKResult(w)

	connector, ok := r.Context().Value("connector").(*db.Connector)
	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var postData models.InputTaskData
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	err = connector.CreateTask(postData)
	if err != nil {
		err = json.NewEncoder(w).Encode(
			models.APIResponse{Err: err.Error()},
		)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		return
	}

	err = json.NewEncoder(w).Encode(
		models.APIResponse{Result: "OK"},
	)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	utils.SetDeleteOKResult(w)

	connector, ok := r.Context().Value("connector").(*db.Connector)
	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	taskID, ok := vars["task_id"]

	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	err := connector.DeleteTask(taskID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(
		models.APIResponse{Result: "OK"},
	)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func DeleteTasksHandler(w http.ResponseWriter, r *http.Request) {
	utils.SetDeleteOKResult(w)

	connector, ok := r.Context().Value("connector").(*db.Connector)
	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	err := connector.DeleteTasks()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(
		models.APIResponse{Result: "OK"},
	)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
