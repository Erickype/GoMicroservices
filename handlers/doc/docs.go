package handler

import "github.com/Erickype/GoMicroservices/data"

// A list of products return in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// Empty response that confirms the operation
// swagger:response noContent
type productsNoContent struct{}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The ID of the product to be deleted from data source
	// in: path
	// required: true
	ID int `json:"id"`
}
