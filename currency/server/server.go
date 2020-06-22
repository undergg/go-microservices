package server

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/undergg/go-microservices-tutorial/currency/data"
	"github.com/undergg/go-microservices-tutorial/currency/protos/currency"
)

type CurrencyServer struct {
	rates *data.ExchangeRates
	log   hclog.Logger
}

func NewCurrencyServer(r *data.ExchangeRates, l hclog.Logger) *CurrencyServer {
	return &CurrencyServer{rates: r, log: l}
}

// That way it implements the CurrencyServer interface.
func (c *CurrencyServer) GetRate(ctx context.Context, rr *currency.RateRequest) (*currency.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())

	if err != nil {
		return nil, err
	}

	return &currency.RateResponse{Rate: rate}, nil
}
