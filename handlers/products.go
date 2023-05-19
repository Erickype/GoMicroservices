package handlers

import (
	"github.com/Erickype/GoMicroservices/data"
	"log"
	"net/http"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{
		logger: logger,
	}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal", http.StatusInternalServerError)
		return
	}
}
