package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking application health")
	response := map[string]string{
		"status" : "UP",
		"timestamp" : time.Now().String(),
	}

	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving home page")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application up and running")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler)

	r.HandleFunc("/", rootHandler)

	log.Fatal(http.ListenAndServe(":80", r))
}
