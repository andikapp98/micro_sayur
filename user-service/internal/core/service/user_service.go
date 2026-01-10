package service

import (
	"context"
	"errors"
	"user-service/internal/adapter/repository"
	"user-service/internal/core/domain/entity"
	"user-service/utils/conv"

	"github.com/labstack/gommon/log"
)

type UserServiceInterface interface {
	SignIn(ctx context.Context, req entity.UserEntity) (*entity.UserEntity, string, error)
}

type UserService struct {
	repo repository.UserRepositoryInterface
}

func (u *UserService) SignIn(ctx context.Context, req entity.UserEntity) (*entity.UserEntity, string, error) {
	user, err := u.repo.GetByUserEmail(ctx, req.Email)
	if err != nil {
		log.Errorf("[UserService-1] SignIn: %v", err)
		return nil, "", err
	}

	if checkPassword := conv.CheckPasswordHash(req.Password, user.Password); !checkPassword {
		err = errors.New("password is incorrect")
		log.Errorf("[UserService-1] SignIn: %v", err)
		return nil, "", err
	}

	return user, "", nil
}