package biz

import "context"

// Repository defines DB operations
type Repository interface {
	// CreateUser creates a new user in DB
	CreateUser(ctx context.Context, user User) error
	// ListAllUsers list all registered users
	ListAllUsers(ctx context.Context) ([]User, error)
	// GetUserByName gets the user information of the mentioned user
	GetUserByName(ctx context.Context, name string) (User, error)
}
