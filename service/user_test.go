package service_test

import (
	"context"
	"testing"

	"dating-app/entity"
	mock "dating-app/repository/mock_repository"
	"dating-app/service"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestUserServiceRegister_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockUserRepository(ctrl)
	userService := service.NewUserService(userRepo)

	existingUser := entity.User{
		Model: gorm.Model{
			ID: 0,
		},
	}

	userRepo.EXPECT().GetUserByUsername(gomock.Any(), "admin").Return(existingUser, nil)
	userRepo.EXPECT().Register(gomock.Any(), gomock.Any()).Return(nil)

	user := entity.User{
		Username: "admin",
		Email:    "aryadevara@gmail.com",
		Password: "Ary4123#",
	}

	err := userService.Register(context.Background(), user)
	assert.Nil(t, err)
}

func TestUserServiceRegister_ErrAlreadyExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockUserRepository(ctrl)
	userService := service.NewUserService(userRepo)

	existingUser := entity.User{
		Model: gorm.Model{
			ID: 101,
		},
	}

	userRepo.EXPECT().GetUserByUsername(gomock.Any(), "admin").Return(existingUser, nil)

	user := entity.User{
		Username: "admin",
		Email:    "aryadevara@gmail.com",
		Password: "Ary4123#",
	}

	err := userService.Register(context.Background(), user)
	assert.Equal(t, service.ErrUserAlreadyExists, err)
}

func TestUserServiceRegister_ErrDomainNotValid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockUserRepository(ctrl)
	userService := service.NewUserService(userRepo)

	existingUser := entity.User{
		Model: gorm.Model{
			ID: 0,
		},
	}

	userRepo.EXPECT().GetUserByUsername(gomock.Any(), "admin").Return(existingUser, nil)

	user := entity.User{
		Username: "admin",
		Email:    "aryadevara@unvalidxyz.cozz",
		Password: "Ary4123#",
	}

	err := userService.Register(context.Background(), user)
	assert.Equal(t, service.ErrEmailInvalid, err)
}
