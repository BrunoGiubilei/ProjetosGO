package main

import (
	"log"
	"net/http"

	"github.com/brunogiubilei/go-rest-api/handlers"
	"github.com/brunogiubilei/go-rest-api/middleware"
	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.Use(middleware.Cors)

	route.HandleFunc("/api", handlers.GetAPI).Methods("GET")
	route.HandleFunc("/api/greeting", handlers.GetGreeting).Methods("GET")

	log.Println("Starting server on :8000")
	log.Fatal(http.ListenAndServe(":8000", route))
}
