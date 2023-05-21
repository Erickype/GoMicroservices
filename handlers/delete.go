package handlers

import (
	"github.com/Erickype/GoMicroservices/data"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// DeleteProduct deletes a product from data source based on its ID
// @Summary      Delete a product
// @Description  Delete a product based on its ID
// @Tags         products
// @Param        id   path      int  true  "Account ID"
// @Success      202
// @Router       /products/{id} [delete]
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
