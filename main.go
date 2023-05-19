package main

import (
	"context"
	"github.com/Erickype/GoMicroservices/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "coffee-api", log.LstdFlags)

	helloHandler := handlers.NewHello(logger)
	productsHandler := handlers.NewProducts(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/hello", helloHandler)
	serveMux.Handle("/products", productsHandler)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	sig := <-signalChannel
	log.Println("Received terminated, graceful shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	_ = server.Shutdown(timeoutContext)
}
