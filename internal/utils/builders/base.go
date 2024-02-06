package builders

import (
	"github.com/gorilla/mux"
	"specialist/internal/app/api/base"
	"specialist/internal/models"
)

func InfoBuilderResource(apiRouter *mux.Router, apiPrefix string) {
	info := &models.Info{
		Name:    "API",
		Version: "1.0.0",
	}

	apiRouter.HandleFunc(
		apiPrefix+"/",
		base.Handler(info),
	).Methods("GET")
}
