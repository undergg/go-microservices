package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	protos "github.com/undergg/go-microservices-tutorial/currency/protos/currency"
	"github.com/undergg/go-microservices-tutorial/product-api/data"
)

// swagger:route GET /products/{id} products listSingle
// Returns a list of products.
// responses:
// 200: productsResponse
// 404: errorResponse
// 501: errorResponse

// GetSingleProduct returns a list of products from the database.
func (p *Products) GetSingleProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	p.logger.Println("Handle GET Single Product request")

	// Get the ID from the URL parameters.
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	prod, err := data.GetProduct(id)

	if err == data.ErrProductNotFound {
		p.logger.Println("Unable to delete record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.logger.Println("Unexpected error from GetProduct")

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
	}

	rr := &protos.RateRequest{
		Base:        protos.Currencies_EUR,
		Destination: protos.Currencies_GBP,
	}

	if err != nil {
		p.logger.Println("[Error] error getting new rate", err)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	resp, err := p.cc.GetRate(context.Background(), rr)

	prod.Price = prod.Price * resp.Rate

	err = data.ToJSON(prod, rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

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
