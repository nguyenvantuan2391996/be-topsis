package usecase

import (
	"context"

	"topsis/internal/domain/model"
	"topsis/internal/domain/repository"
)

type UserDomain struct {
	userRepo repository.UserRepositoryInterface
}

func NewUserDomain(
	userRepo repository.UserRepositoryInterface,
) *UserDomain {
	return &UserDomain{
		userRepo: userRepo,
	}
}

func (u *UserDomain) CreateUser(ctx context.Context, name string) (*model.User, error) {
	return u.userRepo.CreateUser(ctx, &model.User{
		Name: name,
	})
}
