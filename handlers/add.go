package handlers

import (
	"net/http"
	"simple-microservice1/data"
)

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Products request")

	prod := r.Context().Value(data.Product{}).(*data.Product)

	data.AddProduct(prod)
}
