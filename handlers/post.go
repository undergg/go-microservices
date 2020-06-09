package handlers

import (
	"net/http"

	"github.com/undergg/go-microservices-tutorial/data"
)

// swagger:route POST /products products addProduct
// Add a product in the database.
// responses:
// 201: noContentResponse

// AddProduct adds a product in the database.
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Products request")

	prod := r.Context().Value(data.Product{}).(*data.Product)

	data.AddProduct(prod)
}
