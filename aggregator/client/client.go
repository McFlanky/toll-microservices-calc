package client

import (
	"context"

	"github.com/McFlanky/toll-microservices-calc/types"
)

type Client interface {
	Aggregate(context.Context, *types.AggregateRequest) error
}
