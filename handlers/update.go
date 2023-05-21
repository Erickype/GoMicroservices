package handlers

import (
	"github.com/Erickype/GoMicroservices/data"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// UpdateProduct updates a data.Product by passing its ID
//
//	@Summary		Updates a product
//	@Description	Updates a product by passing its ID
//	@Tags			products
//	@Accept			json
//	@Param			id		path	int				true	"Product ID"
//	@Param			product	body	data.Product	true	"Updated product"
//	@Success		202
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/products [put]
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
