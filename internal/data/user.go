package data

import (
	"context"
	"fmt"

	"final/internal/biz"
)


func (r *repositoryImpl) CreateUser(ctx context.Context, user biz.User) error {
	fmt.Println("zyuxin create a user")
	return nil
}

func (r *repositoryImpl) ListAllUsers(ctx context.Context,) ([]biz.User, error) {
	return []biz.User{}, nil
}

func (r *repositoryImpl) GetUserByName(ctx context.Context, name string) (biz.User, error) {
	return biz.User{}, nil
}
