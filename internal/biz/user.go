package biz

import "context"

type UserStatus int32

const (
	UserStatus_ACTIVATED   UserStatus = 0
	UserStatus_DEACTIVATED UserStatus = 1
)

// User defines the Domain Object for a user. For simplification purposes, it is used as PO (Persistence object) as well.
type User struct {
	Name         string     `gorm:"not null"`
	ID           string     `gorm:"not null"`
	Email        string     `gorm:"not null"`
	CreationTime int64      `gorm:"not null;type:bigint"`
	Status       UserStatus `gorm:"not null"`
}

// UserIf defines crud functions to operate user object(s).
type UserIf interface {
	CreateUser(ctx context.Context, user User) error
	ListUser(ctx context.Context) ([]User, error)
}

func NewUserImpl(repository Repository) UserIf {
	return &UserImpl{repo: repository}
}

type UserImpl struct {
	repo Repository
}

func (u *UserImpl) CreateUser(ctx context.Context, user User) error {
	return u.repo.CreateUser(ctx, user)
}

func (u *UserImpl) ListUser(ctx context.Context) ([]User, error) {
	return u.repo.ListAllUsers(ctx)
}
