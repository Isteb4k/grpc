package grpc

import (
	"context"
	"google.golang.org/grpc"
	"grpc/client/internal/entity"
	"grpc/protos/user/v1"
	"log"
)

// Client - provides interaction with the server
type Client interface {
	CreateUser(ctx context.Context, firstName, lastName string) (entity.User, error)
	UpdateUser(ctx context.Context, id, firstName, lastName string) (entity.User, error)
	GetUser(ctx context.Context, id string) (entity.User, error)
}

type client struct {
	v1 v1.UserServiceClient
}

// NewGRPC - create a new gRPC transport
func NewGRPC() Client {
	return &client{
		v1: getClient(),
	}
}

// CreateUser - create a new user using gRPC
func (c client) CreateUser(ctx context.Context, firstName, lastName string) (entity.User, error) {
	resp, err := c.v1.CreateUser(ctx, &v1.CreateUserRequest{
		FirstName: firstName,
		LastName:  lastName,
	})
	if err != nil {
		return entity.User{}, err
	}

	u := c.convertUserProtoToUser(resp.User)

	return u, nil
}

// UpdateUser - update a user using gRPC
func (c client) UpdateUser(ctx context.Context, id, firstName, lastName string) (entity.User, error) {
	resp, err := c.v1.UpdateUser(ctx, &v1.UpdateUserRequest{
		Id:        id,
		LastName:  firstName,
		FirstName: lastName,
	})
	if err != nil {
		return entity.User{}, err
	}

	u := c.convertUserProtoToUser(resp.User)

	return u, nil
}

// GetUser - get a user using gRPC
func (c client) GetUser(ctx context.Context, id string) (entity.User, error) {
	resp, err := c.v1.GetUser(ctx, &v1.GetUserRequest{
		Id: id,
	})
	if err != nil {
		return entity.User{}, err
	}

	u := c.convertUserProtoToUser(resp.User)

	return u, nil
}

func (c client) convertUserProtoToUser(up *v1.User) entity.User {
	return entity.User{
		ID:        up.Id,
		FirstName: up.FirstName,
		LastName:  up.LastName,
	}
}

func getClient() v1.UserServiceClient {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	return v1.NewUserServiceClient(conn)
}
