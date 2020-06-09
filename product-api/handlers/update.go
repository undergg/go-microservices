package handlers

import (
	"net/http"

	"github.com/undergg/go-microservices-tutorial/product-api/data"
)

// swagger:route PUT /products products putProduct
// Updates a product from the database.
// responses:
// 201: noContentResponse
// 404: errorResponse

// UpdateProduct updates a product from the database.
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	p.logger.Println("Handle PUT Products request")

	prod := r.Context().Value(data.Product{}).(*data.Product)

	err := data.UpdateProduct(prod)

	if err == data.ErrProductNotFound {
		p.logger.Println("Unable to update record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.logger.Println("Unable to update record")

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
