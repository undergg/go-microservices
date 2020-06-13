package handlers

import (
	"log"

	protos "github.com/undergg/go-microservices-tutorial/currency/protos/currency"
)

// Products handler for interacting with the products database.
type Products struct {
	logger *log.Logger
	cc     protos.CurrencyClient
}

// NewProducts returns a new handler of type Products.
func NewProducts(logger *log.Logger, cc protos.CurrencyClient) *Products {
	return &Products{logger, cc}
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string
}

// ValidationError is used for Validation errors.
type ValidationError struct {
	Messages []string
}
