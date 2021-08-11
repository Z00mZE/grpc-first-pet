package adder

import (
	"context"
	"grpc-first-pet/pkg/api"
)

type GRPCServer struct{}

func (G GRPCServer) Add(ctx context.Context, request *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: request.GetX() + request.GetY()}, nil
}
