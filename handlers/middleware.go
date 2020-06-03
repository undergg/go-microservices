package handlers

import (
	"context"
	"net/http"
	"simple-microservice1/data"
)

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

		err = prod.Validate()
		if err != nil {
			http.Error(rw, "Error validating product.", http.StatusBadRequest)
			return
		}

		context := context.WithValue(r.Context(), data.Product{}, prod)
		r = r.WithContext(context)

		// If you don't call this, then it will not call the next http handler.
		next.ServeHTTP(rw, r)

	})
}
