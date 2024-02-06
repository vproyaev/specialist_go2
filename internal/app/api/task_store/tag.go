package task_store

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"specialist/internal/db"
	"specialist/internal/models"
	"specialist/internal/utils"
)

func GetTagsHandler(w http.ResponseWriter, r *http.Request) {
	utils.SetOKResult(w)

	connector, ok := r.Context().Value("connector").(*db.Connector)
	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	result, err := connector.GetTags()
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

func GetTagHandler(w http.ResponseWriter, r *http.Request) {
	utils.SetOKResult(w)

	connector, ok := r.Context().Value("connector").(*db.Connector)
	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	tagID, ok := vars["tag_id"]
	tagError := ""

	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	result, err := connector.GetTag(tagID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	} else if result.IsNull() {
		tagError = "Tag not found"
	}

	if tagError != "" {
		err = json.NewEncoder(w).Encode(
			models.APIResponse{Err: tagError},
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

func PostTagHandler(w http.ResponseWriter, r *http.Request) {
	utils.SetPostOKResult(w)

	connector, ok := r.Context().Value("connector").(*db.Connector)
	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var postData models.InputTagData
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	err = connector.CreateTag(postData)
	if err != nil {
		err = json.NewEncoder(w).Encode(
			models.APIResponse{Err: err.Error()},
		)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

	err = json.NewEncoder(w).Encode(
		models.APIResponse{Result: "OK"},
	)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
