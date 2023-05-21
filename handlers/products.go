package handlers

import (
	"context"
	"fmt"
	"github.com/Erickype/GoMicroservices/data"
	"log"
	"net/http"
)

// Products struct that describes a handler that have a *log.Logger reference for log information
type Products struct {
	logger *log.Logger
}

// NewProducts function that creates a Products struct that acts as a handler
func NewProducts(logger *log.Logger) *Products {
	return &Products{
		logger: logger,
	}
}

// KeyProduct struct to create a context with value
type KeyProduct struct{}

// MiddlewareProductValidation is a middleware that validates a product
// checking if it can be deserialized and if it has valid fields.
func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		product := &data.Product{}
		err := product.FromJSON(r.Body)
		if err != nil {
			p.logger.Println("Error deserializing product!", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		err = product.Validate()
		if err != nil {
			p.logger.Println("Error validating product!", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
