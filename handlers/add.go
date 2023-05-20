package handlers

import (
	"github.com/Erickype/GoMicroservices/data"
	"net/http"
)

func (p *Products) AddProduct(_ http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST product")
	product := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(product)
}
