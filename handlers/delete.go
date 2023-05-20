package handlers

import (
	"github.com/Erickype/GoMicroservices/data"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product from data source based on its ID
// responses:
//  201: noContent

// DeleteProduct deletes a product from data source based on its ID
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle DELETE product")
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
