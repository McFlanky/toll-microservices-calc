package main

import (
	"context"

	"github.com/McFlanky/toll-microservices-calc/types"
)

type Service interface {
	Aggregate(context.Context, types.Distance) error
	Calculate(context.Context, int) (*types.Invoice, error)
}

type BasicService struct {
}

func newBasicService() Service {
	return &BasicService{}
}

func (svc *BasicService) Aggregate(_ context.Context, dist types.Distance) error {
	return nil
}

func (svc *BasicService) Calculate(_ context.Context, dist int) (*types.Invoice, error) {
	return nil, nil
}

// NewAggregatorService will construct a complete microservice
// with logging and instrumentation middleware
func NewAggregatorService() Service {
	svc := newBasicService()

	return nil
}
