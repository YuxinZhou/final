package account

import (
	"context"
	"fmt"
	"time"

	"final/api"
	"final/internal/biz"
)

// Server is used to implement api.AccountServer.
type Server struct {
	api.UnimplementedAccountServer
	UserIf biz.UserIf
}

// CreateUser implements AccountServer.CreateUser
func (s *Server) CreateUser(ctx context.Context, in *api.UserInfo) (*api.Empty, error) {
	user := biz.User{
		Name:         in.Name,
		Email:        in.Email,
		CreationTime: time.Now().Unix(),
		Status:       biz.UserStatus_ACTIVATED,
	}

	err := s.UserIf.CreateUser(ctx, user)
	if err != nil {
		fmt.Printf("err: %+v", err) // log the stack trace
	}

	return &api.Empty{}, err
}
