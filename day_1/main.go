package main

import (
	"log"
	"net/http"

	"day_1/api"
	"day_1/middleware"
	"day_1/models"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Trying to start server")
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.ErrorHandlingMiddleware)

	info := &models.Info{
		Name:    "API",
		Version: "1.0.0",
	}

	router.HandleFunc(
		"/info",
		api.InfoHandler(info),
	).Methods("GET")

	router.HandleFunc(
		"/second", api.CalculatorHandler,
	).Methods("GET")

	router.HandleFunc(
		"/first", api.CalculatorHandler,
	).Methods("GET")

	router.HandleFunc(
		"/add", api.CalculatorHandler,
	).Methods("GET")

	router.HandleFunc(
		"/sub", api.CalculatorHandler,
	).Methods("GET")

	router.HandleFunc(
		"/mul", api.CalculatorHandler,
	).Methods("GET")

	router.HandleFunc(
		"/div", api.CalculatorHandler,
	).Methods("GET")

	const serverPort = ":8080"
	go func() {
		err := http.ListenAndServe(serverPort, router)
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Server started: http://localhost" + serverPort)
	log.Println("Press Ctrl+C to stop the server")
	select {}
}
