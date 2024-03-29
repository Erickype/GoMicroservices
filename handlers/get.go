package handlers

import (
	"github.com/Erickype/GoMicroservices/data"
	"net/http"
)

// GetProducts returns the list of products from the data store
//
//	@Summary		Returns the list of products
//	@Description	Returns all products from data source
//	@Tags			products
//	@Produce		json
//	@Success		200	{object}	data.Products
//	@Failure		500
//	@Router			/products [get]
func (p *Products) GetProducts(rw http.ResponseWriter, _ *http.Request) {
	p.logger.Println("Handle GET products")

	products := data.GetProducts()
	err := products.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal", http.StatusInternalServerError)
		return
	}
}
