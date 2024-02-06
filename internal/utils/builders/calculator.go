package builders

import (
	"github.com/gorilla/mux"
	"specialist/internal/app/api/calculator"
)

func CalculatorBuilderResource(apiRouter *mux.Router, apiPrefix string) {
	calculatorPrefixes := []string{
		"/first",
		"/second",
		"/add",
		"/sub",
		"/mul",
		"/div",
	}

	for _, uri := range calculatorPrefixes {
		apiRouter.HandleFunc(
			apiPrefix+uri,
			calculator.Handler,
		).Methods("GET")
	}
}
