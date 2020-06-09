package handlers

import (
	"log"
)

type Products struct {
	logger *log.Logger
}

// NewProducts returns a new handler of type Products.
func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string
}

// ValidationError is used for Validation errors.
type ValidationError struct {
	Messages []string
}
