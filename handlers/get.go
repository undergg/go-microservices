package handlers

import (
	"net/http"

	"github.com/undergg/go-microservices-tutorial/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products.
// responses:
// 200: productsResponse

// GetProducts returns a list of products from the database.
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	// In our Swagger documentation we defined that our API will produce "application/json".
	// Therefore we need to set the header of the response.
	rw.Header().Add("Content-Type", "application/json")

	p.logger.Println("Handle GET Products request")

	lp := data.GetProducts()
	err := data.ToJSON(lp, rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}
