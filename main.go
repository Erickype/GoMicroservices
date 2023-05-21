package main

import (
	"context"
	_ "github.com/Erickype/GoMicroservices/docs"
	"github.com/Erickype/GoMicroservices/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//	@title			Products API
//	@version		1.0.0
//	@description	This is the documentation for Products API.

//	@contact.name	Erickype
//	@contact.email	erickype@hotmail.com

// @host		localhost:9090
// @BasePath	/
func main() {
	logger := log.New(os.Stdout, "products-api", log.LstdFlags)

	productsHandler := handlers.NewProducts(logger)

	serveMux := mux.NewRouter()

	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", productsHandler.GetProducts)

	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", productsHandler.UpdateProduct)
	putRouter.Use(productsHandler.MiddlewareProductValidation)

	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", productsHandler.AddProduct)
	postRouter.Use(productsHandler.MiddlewareProductValidation)

	deleteRouter := serveMux.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", productsHandler.DeleteProduct)

	getRouter.PathPrefix("/docs/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:9090/docs/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	))

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
