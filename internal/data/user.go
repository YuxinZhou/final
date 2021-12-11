package data

import (
	"context"
	"fmt"

	"final/internal/biz"

	"github.com/pkg/errors"
)

const userInsert = `INSERT INTO ` + userTableName + `(name, email, creation_time, status) VALUES (?, ?, ?, ?)`

func (r *repositoryImpl) CreateUser(ctx context.Context, user biz.User) error {
	stmt, err := r.db.PrepareContext(ctx, userInsert)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to prepare context when create the user '%s'", user.Name))
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.CreationTime, user.Status)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to execute query when create the user '%s'", user.Name))
	}

	return nil
}

func (r *repositoryImpl) ListAllUsers(ctx context.Context, ) ([]biz.User, error) {
	return []biz.User{}, nil
}

func (r *repositoryImpl) GetUserByName(ctx context.Context, name string) (biz.User, error) {
	return biz.User{}, nil
}
