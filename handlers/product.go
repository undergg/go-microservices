package handlers

import (
	"context"
	"log"
	"net/http"
	"simple-microservice1/data"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle GET Products request")

	lp := data.GetProducts()
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Products request")

	prod := r.Context().Value(data.Product{}).(*data.Product)

	data.AddProduct(prod)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle PUT Products request")

	// Get the ID from the URL parameters.
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	prod := r.Context().Value(data.Product{}).(*data.Product)

	err = data.UpdateProduct(id, prod)

	if err != nil {
		http.Error(rw, "The product was not updated", http.StatusBadRequest)
		return
	}
}

func (p *Products) ValidateProductMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		p.logger.Println("Enter ValidateProductMiddleware")

		prod := &data.Product{}
		// In the r.body there is a JSON object of type Product.
		err := prod.FromJSON(r.Body)

		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		context := context.WithValue(r.Context(), data.Product{}, prod)
		r = r.WithContext(context)

		// If you don't call this, then it will not call the next http handler.
		next.ServeHTTP(rw, r)

	})
}
