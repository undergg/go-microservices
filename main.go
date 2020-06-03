package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"simple-microservice1/handlers"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting the server..")

	logger := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProducts(logger)

	// Gorilla framework serveMux.
	serveMux := mux.NewRouter()

	// Define subroutes.
	getProductsRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getProductsRouter.HandleFunc("/products", ph.GetProducts)

	addProductsRouter := serveMux.Methods(http.MethodPost).Subrouter()
	addProductsRouter.HandleFunc("/products", ph.AddProduct)
	// Use Middleware to intercept the request and validate the JSON object of the request.
	addProductsRouter.Use(ph.ValidateProductMiddleware)

	putProductsRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putProductsRouter.HandleFunc("/products", ph.UpdateProduct)
	putProductsRouter.Use(ph.ValidateProductMiddleware)

	deleteProductsRouter := serveMux.Methods(http.MethodDelete).Subrouter()
	deleteProductsRouter.HandleFunc("/products/{id:[0-9]+}", ph.DeleteProduct)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Use a go function so that it won't block here.
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Receive terminate, graceful shutdown", sig)

	cxt, _ := context.WithTimeout(context.Background(), 30*time.Second)

	server.Shutdown(cxt)

}
