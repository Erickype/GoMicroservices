package handlers

import (
	"github.com/Erickype/GoMicroservices/data"
	"net/http"
)

func (p *Products) GetProducts(rw http.ResponseWriter, _ *http.Request) {
	p.logger.Println("Handle GET products")

	products := data.GetProducts()
	err := products.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal", http.StatusInternalServerError)
		return
	}
}
