package builders

import (
	"github.com/gorilla/mux"
	"specialist/internal/app/api/task_store"
)

func DueTasksStoreBuilderResource(apiRouter *mux.Router, apiPrefix string) {
	// GET
	apiRouter.PathPrefix(apiPrefix).Queries(
		"year", "{year:[0-9]+}",
		"month", "{month:[0-9]+}",
		"day", "{day:[0-9]+}",
	).HandlerFunc(
		task_store.GetDueTasksHandler,
	).Methods("GET")
}
