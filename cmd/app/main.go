package main

import (
	"fmt"
	"github.com/Mishanki/specialist-dz-1/internal/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/items", handlers.GetItems).Methods("GET")
	router.HandleFunc("/api/v1/item/{id}", handlers.GetItem).Methods("GET")
	router.HandleFunc("/api/v1/item/{id}", handlers.CreateItem).Methods("POST")
	router.HandleFunc("/api/v1/item/{id}", handlers.UpdateItem).Methods("PUT")
	router.HandleFunc("/api/v1/item/{id}", handlers.DeleteItem).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		panic("Param PORT is not found in .env")
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
