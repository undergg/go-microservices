package server

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/undergg/go-microservices-tutorial/currency/protos/currency"
)

type CurrencyServer struct {
	log hclog.Logger
}

func NewCurrencyServer(l hclog.Logger) *CurrencyServer {
	return &CurrencyServer{log: l}
}

// That way it implements the CurrencyServer interface.
func (c *CurrencyServer) GetRate(ctx context.Context, rr *currency.RateRequest) (*currency.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())
	return &currency.RateResponse{Rate: 0.5}, nil
}
