package builders

import (
	"github.com/gorilla/mux"
	"specialist/internal/app/api/task_store"
)

func TaskStoreBuilderResource(apiRouter *mux.Router, apiPrefix string) {
	// GET
	apiRouter.HandleFunc(
		apiPrefix,
		task_store.GetTasksHandler,
	).Methods("GET")
	apiRouter.HandleFunc(
		apiPrefix+`/{task_id:[0-9]+}`,
		task_store.GetTaskHandler,
	).Methods("GET")

	// POST
	apiRouter.HandleFunc(
		apiPrefix+"/",
		task_store.PostTaskHandler,
	).Methods("POST")

	// DELETE
	apiRouter.HandleFunc(
		apiPrefix+`/{task_id:[0-9]+}`,
		task_store.DeleteTaskHandler,
	).Methods("DELETE")
	apiRouter.HandleFunc(
		apiPrefix+"/",
		task_store.DeleteTasksHandler,
	).Methods("DELETE")
}
