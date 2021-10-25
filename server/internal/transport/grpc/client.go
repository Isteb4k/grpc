package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc/server/internal/user"
	"log"
	"net"

	"google.golang.org/grpc"
	"grpc/protos/user/v1"
)

// Client - will handle incoming gRPC requests
type Client struct {
	UserService user.Service
}

// NewGRPC - create a new gRPC transport client
func NewGRPC(userService user.Service) Client {
	return Client{
		UserService: userService,
	}
}

// Run - run transport, listen port and serve gRPC
func (c Client) Run() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Print("could not listen on port 50051")
		return err
	}

	grpcServer := grpc.NewServer()
	v1.RegisterUserServiceServer(grpcServer, &c)

	log.Print("start to serve grpc")
	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %s\n", err)
		return err
	}

	return nil
}

// GetUser - retrieves a user by id and returns the response.
func (c Client) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	log.Print("GetUser gRPC Endpoint Hit")

	u, err := c.UserService.GetUser(ctx, req.Id)
	if err != nil {
		log.Print("Failed to retrieve user by ID")
		return &v1.GetUserResponse{}, err
	}

	userProto := c.convertUserToUserProto(u)

	return &v1.GetUserResponse{
		User: &userProto,
	}, nil
}

// CreateUser - handler for creating a new user
func (c Client) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	log.Print("CreateUser gRPC endpoint hit")
	u, err := c.UserService.CreateUser(ctx, req.FirstName, req.LastName)
	if err != nil {
		log.Print("failed to insert rocket into database")
		return &v1.CreateUserResponse{}, err
	}

	userProto := c.convertUserToUserProto(u)

	return &v1.CreateUserResponse{
		User: &userProto,
	}, nil
}

// UpdateUser - handler for updating existing user
func (c Client) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	log.Print("UpdateUser gRPC endpoint hit")
	u, err := c.UserService.UpdateUser(ctx, req.Id, req.FirstName, req.LastName)
	if err != nil {
		return &v1.UpdateUserResponse{}, err
	}

	userProto := c.convertUserToUserProto(u)

	return &v1.UpdateUserResponse{
		User: &userProto,
	}, nil
}

func (c Client) convertUserToUserProto(u user.User) v1.User {
	var up v1.User

	up = v1.User{
		Id:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		CreatedAt: timestamppb.New(u.CreatedAt),
	}

	if u.UpdatedAt != nil {
		up.UpdatedAt = timestamppb.New(*u.UpdatedAt)
	}

	return up
}
