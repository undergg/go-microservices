package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/undergg/go-microservices-tutorial/product-api/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product from the database.
// responses:
// 201: noContentResponse
// 404: errorResponse
// 501: errorResponse

// DeleteProduct deletes a product from the database.
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	// Get the ID from the URL parameters.
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	err = data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		p.logger.Println("Unable to delete record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.logger.Println("Unable to delete record")

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
