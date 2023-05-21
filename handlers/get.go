package handlers

import (
	"github.com/Erickype/GoMicroservices/data"
	"net/http"
)

// GetProducts returns the list of products from the data store
// @Summary	get products
// @Description  Returns all products from data source
// @Tags         products
// @Produce      json
// @Success      200 {object}   data.Products
// @Router       /products [get]
func (p *Products) GetProducts(rw http.ResponseWriter, _ *http.Request) {
	p.logger.Println("Handle GET products")

	products := data.GetProducts()
	err := products.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal", http.StatusInternalServerError)
		return
	}
}
