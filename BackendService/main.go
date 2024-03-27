package main

import (
	"fmt"
	"net/http"

	handler "github.com/ali-ammar-kazmi/backend-service/handler"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Register the Environment variables of .env file with in go environment
	godotenv.Load()

	// Create a New Router Instance
	router := mux.NewRouter()

	// Register your API handlers with the CORS middleware
	headerOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "PUT", "OPTIONS"})

	// API request handler
	router.HandleFunc("/weather", handler.GetWeather).Methods("GET")
	router.HandleFunc("/news", handler.GetNews).Methods("GET")

	fmt.Println("Server listening at port - http://localhost:8080")

	// Initializing Server
	if err := http.ListenAndServe("localhost:8080", handlers.CORS(originOk, headerOk, methodsOk)(router)); err != nil {
		panic(err.Error())
	}
}
