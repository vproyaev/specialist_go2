package builders

import (
	"github.com/gorilla/mux"
	"specialist/internal/app/api/task_store"
)

func TagStoreBuilderResource(apiRouter *mux.Router, apiPrefix string) {
	// GET
	apiRouter.HandleFunc(
		apiPrefix,
		task_store.GetTagsHandler,
	).Methods("GET")
	apiRouter.HandleFunc(
		apiPrefix+`/{tag_id:[0-9]+}`,
		task_store.GetTagHandler,
	).Methods("GET")

	//// POST
	apiRouter.HandleFunc(
		apiPrefix+"/",
		task_store.PostTagHandler,
	).Methods("POST")
}
