package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	protos "github.com/undergg/go-microservices-tutorial/currency/protos/currency"
	"github.com/undergg/go-microservices-tutorial/product-api/handlers"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting the server..")

	logger := log.New(os.Stdout, "product-api", log.LstdFlags)

	// grpc.WithInsecure ---> HTTP/2 with no TLS
	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	// create the Currency Client.
	cc := protos.NewCurrencyClient(conn)

	ph := handlers.NewProducts(logger, cc)

	// Gorilla framework serveMux.
	serveMux := mux.NewRouter()

	// Define subroutes per HTTP Verb.
	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", ph.GetProducts)
	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.GetSingleProduct)

	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", ph.AddProduct)
	// Use Middleware to intercept the request and validate the JSON object of the request.
	postRouter.Use(ph.ValidateProductMiddleware)

	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products", ph.UpdateProduct)
	putRouter.Use(ph.ValidateProductMiddleware)

	deleteRouter := serveMux.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", ph.DeleteProduct)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

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
