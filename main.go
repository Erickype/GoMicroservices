package main

import (
	"context"
	"github.com/Erickype/GoMicroservices/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "products-api", log.LstdFlags)

	productsHandler := handlers.NewProducts(logger)

	serveMux := mux.NewRouter()

	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", productsHandler.GetProducts)

	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", productsHandler.UpdateProduct)
	putRouter.Use(productsHandler.MiddlewareProductValidation)

	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", productsHandler.AddProduct)
	postRouter.Use(productsHandler.MiddlewareProductValidation)

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

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_ = server.Shutdown(timeoutContext)
}
