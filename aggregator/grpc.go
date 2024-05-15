package main

import "github.com/McFlanky/toll-microservices-calc/types"

type GRPCAggregatorServer struct {
	types.UnimplementedAggregatorServer
	svc Aggregator
}

func NewAggregatorGRPCServer(svc Aggregator) *GRPCAggregatorServer {
	return &GRPCAggregatorServer{
		svc: svc,
	}
}

// transport layer
// JSON -> types.Distance -> all done (same types)
// GRPC -> types.AggregateRequest -> types.Distance
// Webpack -> types.Webpack -> types.Distance

// business layer -> business layer type (main type all needs to convert to)

func (s *GRPCAggregatorServer) AggregateDistance(req *types.AggregateRequest) error {
	distance := types.Distance{
		OBUID: int(req.ObuID),
		Value: req.Value,
		Unix:  req.Unix,
	}
	return s.svc.AggregateDistance(distance)
}
