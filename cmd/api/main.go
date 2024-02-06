package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"specialist/internal/db"
	"specialist/internal/middleware"
	"specialist/internal/utils/builders"
)

const serverPort = ":8080"
const apiPrefix string = "/api/v1"
const infoPrefix string = apiPrefix + "/info"
const calculatorPrefix string = apiPrefix + "/calculator"
const taskStorePrefix string = apiPrefix + "/task"
const dueTasksStorePrefix string = taskStorePrefix + "/due"
const tagStorePrefix string = apiPrefix + "/tag"

func main() {

	log.Println("Trying to connect to database")
	connStr := "postgres://admin:admin@localhost:5432/task_store?sslmode=disable"
	connector := db.Connector{}
	connector.Create(connStr)

	err := connector.Db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected")

	log.Println("Trying to start server")
	router := mux.NewRouter().StrictSlash(true)

	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.ErrorHandlingMiddleware)
	router.Use(
		func(next http.Handler) http.Handler {
			return middleware.DBMiddleware(&connector, next)
		},
	)

	builders.InfoBuilderResource(router, infoPrefix)
	builders.CalculatorBuilderResource(router, calculatorPrefix)
	builders.TaskStoreBuilderResource(router, taskStorePrefix)
	builders.TagStoreBuilderResource(router, tagStorePrefix)
	builders.DueTasksStoreBuilderResource(router, dueTasksStorePrefix)

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
