package main

import (
	"fmt"
	"net/http"

	"github.com/ali-ammar-kazmi/backend-service/handler"
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
	headerOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET"})

	// API request handler
	router.HandleFunc("/weather", handler.GetWeather).Methods("GET")
	router.HandleFunc("/news", handler.GetNews).Methods("GET")
	router.HandleFunc("/image", handler.GetImage).Methods("GET")

	fmt.Println("Server listening at port - http://localhost:8080")

	// Initializing Server
	if err := http.ListenAndServe("localhost:8080", handlers.CORS(originOk, headerOk, methodsOk)(router)); err != nil {
		panic(err.Error())
	}
}
