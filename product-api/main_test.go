package main

import (
	"testing"

	"github.com/undergg/go-microservices-tutorial/product-api/sdk/client/products"

	"github.com/undergg/go-microservices-tutorial/product-api/sdk/client"
)

func TestOurClient(t *testing.T) {
	// Default host is "localhost" which (since its http) it listens to port 80.
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := products.NewListProductsParams()
	_, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}
}
