package handlers

import (
	"context"
	"github.com/Erickype/GoMicroservices/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{
		logger: logger,
	}
}

func (p *Products) GetProducts(rw http.ResponseWriter, _ *http.Request) {
	p.logger.Println("Handle GET products")

	products := data.GetProducts()
	err := products.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal", http.StatusInternalServerError)
		return
	}
}

func (p *Products) AddProduct(_ http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST product")
	product := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(product)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle PUT product")
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Invalid ID", http.StatusBadRequest)
		return
	}

	product := r.Context().Value(KeyProduct{}).(*data.Product)

	err = data.UpdateProduct(id, product)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		product := &data.Product{}
		err := product.FromJSON(r.Body)
		if err != nil {
			p.logger.Println("Error deserializing product!", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
