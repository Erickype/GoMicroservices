package main

import (
	"github.com/Erickype/GoMicroservices/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "bonsai-api", log.LstdFlags)

	helloHandler := handlers.NewHello(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/hello", helloHandler)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	_ = server.ListenAndServe()
}
