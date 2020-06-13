// Package classification of Product API
//
// Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import (
	"github.com/undergg/go-microservices-tutorial/product-api/data"
)

// A list of products.
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the database.
	// in: body
	Body data.Products
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentsResponseWrapper struct {
	// No content.
}

// swagger:parameters deleteProduct listSingleProduct
type productIDParameterWrapper struct {
	// The id of the product to delete from the database.
	// in: path
	// required: true
	ID int `json:"id"`
}

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}
