package main

import (
	"net"
	"os"

	"google.golang.org/grpc/reflection"

	"github.com/hashicorp/go-hclog"
	"github.com/undergg/go-microservices-tutorial/currency/protos/currency"
	"github.com/undergg/go-microservices-tutorial/currency/server"
	"google.golang.org/grpc"
)

func main() {

	log := hclog.Default()

	gs := grpc.NewServer()
	cs := server.NewCurrencyServer(log)

	currency.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")

	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	log.Info("Starting grpc server..")
	gs.Serve(l)
}
