package main

import (
	"net/http"

	"day_1/api"
	"day_1/models"
)

func main() {
	info := &models.Info{
		Name:    "API",
		Version: "1.0.0",
	}

	http.HandleFunc(
		"/info",
		api.InfoHandler(info),
	)
	http.HandleFunc(
		"/second", api.CalculatorHandler,
	)
	http.HandleFunc(
		"/first", api.CalculatorHandler,
	)
	http.HandleFunc(
		"/add", api.CalculatorHandler,
	)
	http.HandleFunc(
		"/sub", api.CalculatorHandler,
	)
	http.HandleFunc(
		"/mul", api.CalculatorHandler,
	)
	http.HandleFunc(
		"/div", api.CalculatorHandler,
	)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
