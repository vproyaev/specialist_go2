package main

import (
	"log"
	"net/http"

	"day_1/api"
	"day_1/models"
)

func main() {
	defaultHandler := http.DefaultServeMux
	customHandler := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("Got request:", r.RequestURI)
			defaultHandler.ServeHTTP(w, r)
		},
	)

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

	const serverPort = ":8080"
	go func() {
		err := http.ListenAndServe(serverPort, customHandler)
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Server started: http://localhost" + serverPort)
	log.Println("Press Ctrl+C to stop the server")
	select {}
}
