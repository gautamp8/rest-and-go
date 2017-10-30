// Entrypoint for API

package main

import (
 	"log"
 	"net/http"
 	"os"
	"github.com/gorilla/handlers"
	"Bingo/store"
)

func main() {
	router := store.NewRouter() // create routes

	// These two lines are important in order to allow access from the front-end side to the methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"}) 
 	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// Launch server with CORS validations
 	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(allowedOrigins, allowedMethods)(router)))
}