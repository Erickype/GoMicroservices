package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Erickype/GoMicroservices/details"
	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking application health")
	response := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}

	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving home page")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application up and running")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching details")
	hostname, err := details.GetHostName()

	if err != nil {
		panic(err)
	}

	ip := details.GetLocalIP()
	fmt.Println(hostname, ip)

	response := map[string]string{
		"hostname": hostname,
		"ip":       ip,
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler)

	r.HandleFunc("/details", detailsHandler)

	r.HandleFunc("/", rootHandler)

	log.Println("Server has started!!")

	log.Fatal(http.ListenAndServe(":80", r))
}
