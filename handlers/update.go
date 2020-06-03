package handlers

import (
	"net/http"
	"simple-microservice1/data"
)

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle PUT Products request")

	prod := r.Context().Value(data.Product{}).(*data.Product)

	err := data.UpdateProduct(prod)

	if err != nil {
		http.Error(rw, "The product was not updated", http.StatusBadRequest)
		return
	}
}
