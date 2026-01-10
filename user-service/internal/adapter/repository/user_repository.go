package repository

import (
	"context"
	"user-service/internal/core/domain/entity"
	"user-service/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetByUserEmail(ctx context.Context, email string) (*entity.UserEntity, error)
}

type userRepository struct{
	db *gorm.DB
}

func (u *userRepository) GetByUserEmail(ctx context.Context, email string) (*entity.UserEntity, error){
	modelUser := model.User{}
	if err := u.db.Where("email = ? && is_verified = ?", email, true).Preload("Roles").First(&modelUser).Error; err != nil{
		log.Errorf("[UserRepository-1] GetByUserEmail: %v", err)
		return nil, err
	}
	return &entity.UserEntity{
		ID:         modelUser.ID,
		Name:       modelUser.Name,
		Email:      email,
		Password:   modelUser.Password,
		RoleName: modelUser.Roles[0].Name,
		IsVerified: modelUser.IsVerified,
		Photo: 		modelUser.Photo,
		Address:    modelUser.Address,
		Phone:      modelUser.Phone,
		Lat:        modelUser.Lat,
		Lng:        modelUser.Lng,
		
	}, nil
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}