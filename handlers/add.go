package handlers

import (
	"github.com/Erickype/GoMicroservices/data"
	"net/http"
)

// AddProduct creates a new data.Product in the data source
//
//	@Summary		Creates a product
//	@Description	Creates a product by passing the "data.Product" model
//	@Tags			products
//	@Accept			json
//	@Param			product	body	data.Product	true	"New product"
//	@Success		202
//	@Failure		400
//	@Failure		500
//	@Router			/products [post]
func (p *Products) AddProduct(_ http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST product")
	product := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(product)
}
