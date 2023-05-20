package handlers

import (
	"github.com/Erickype/GoMicroservices/data"
	"net/http"
)

// swagger:route GET /products products listProducts
// Returns the list of products
// responses:
//  200: productsResponse

// GetProducts returns the list of products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, _ *http.Request) {
	p.logger.Println("Handle GET products")

	products := data.GetProducts()
	err := products.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal", http.StatusInternalServerError)
		return
	}
}
