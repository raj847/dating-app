package repository

import (
	"context"
	"dating-app/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	Register(ctx context.Context, user *entity.User) error
	GetUserByUsername(ctx context.Context, username string) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) Register(ctx context.Context, user *entity.User) error {
	err := u.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetUserByUsername(ctx context.Context, username string) (entity.User, error) {
	var res entity.User
	err := u.db.WithContext(ctx).Table("users").Where("username = ?", username).Find(&res).Error
	if err != nil {
		return entity.User{}, err
	}

	return res, nil
}

func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var res entity.User
	err := u.db.WithContext(ctx).Table("users").Where("email = ?", email).Find(&res).Error
	if err != nil {
		return entity.User{}, err
	}

	return res, nil
}
