package user

import (
	"context"
	"grpc/client/internal/entity"
	"grpc/client/internal/transport/grpc"
)

// Service - defines the interface we expect
type Service interface {
	CreateUser(ctx context.Context, firstName, lastName string) (entity.User, error)
	UpdateUser(ctx context.Context, id, firstName, lastName string) (entity.User, error)
	GetUser(ctx context.Context, id string) (entity.User, error)
}

type service struct {
	grpc grpc.Client
}

// New - create a new user service
func New(grpc grpc.Client) Service {
	return &service{
		grpc: grpc,
	}
}

// CreateUser - create a new user using gRPC transport
func (s service) CreateUser(ctx context.Context, firstName, lastName string) (entity.User, error) {
	return s.grpc.CreateUser(ctx, firstName, lastName)
}

// UpdateUser - update a user using gRPC transport
func (s service) UpdateUser(ctx context.Context, id, firstName, lastName string) (entity.User, error) {
	return s.grpc.UpdateUser(ctx, id, firstName, lastName)
}

// GetUser - get a user using gRPC transport
func (s service) GetUser(ctx context.Context, id string) (entity.User, error) {
	return s.grpc.GetUser(ctx, id)
}
