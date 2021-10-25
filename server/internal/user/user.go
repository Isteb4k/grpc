package user

import (
	"context"
	"time"
)

// User - should contain the definition of user
type User struct {
	ID        string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

// Service - defines the interface we expect
type Service interface {
	CreateUser(ctx context.Context, firstName, lastName string) (User, error)
	UpdateUser(ctx context.Context, id, firstName, lastName string) (User, error)
	GetUser(ctx context.Context, id string) (User, error)
}

type service struct {
}

// New - returns a new instance of our user service
func New() Service {
	return &service{}
}

// CreateUser - create a new user
func (s service) CreateUser(ctx context.Context, firstName, lastName string) (User, error) {
	return User{
		ID:        "test_0",
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}, nil
}

// UpdateUser - update a user data by id
func (s service) UpdateUser(ctx context.Context, id, firstName, lastName string) (User, error) {
	updatedAt := time.Now()

	return User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: time.Now(),
		UpdatedAt: &updatedAt,
	}, nil
}

// GetUser - retrieves a user based on the ID
func (s service) GetUser(ctx context.Context, id string) (User, error) {
	return User{
		ID:        id,
		FirstName: "Tom",
		LastName:  "Sawyer",
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}, nil
}
