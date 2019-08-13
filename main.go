package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func helloEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello!")
}

func main() {
	router := mux.NewRouter()

	// endpoints
	router.HandleFunc("/hello", helloEndpoint).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
