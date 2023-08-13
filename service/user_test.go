package service_test

import (
	"context"
	"errors"
	"testing"

	"dating-app/entity"
	mock "dating-app/repository/mock_repository"
	"dating-app/service"
	"dating-app/utils"

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

	err := userService.Register(context.Background(), &user)
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

	err := userService.Register(context.Background(), &user)
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

	err := userService.Register(context.Background(), &user)
	assert.Equal(t, service.ErrEmailInvalid, err)
}

func TestUserServiceLogin_UsernameSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockUserRepository(ctrl)
	userService := service.NewUserService(userRepo)

	pass, _ := utils.HashPassword("Ary4123#")

	existingUser := entity.User{
		Model: gorm.Model{
			ID: 1,
		},
		Password: pass,
	}

	userRepo.EXPECT().GetUserByUsername(gomock.Any(), "admin").Return(existingUser, nil)

	userLogin := entity.User{
		Username: "admin",
		Password: "Ary4123#",
	}

	user, err := userService.LoginUser(context.Background(), &userLogin)
	assert.Nil(t, err)
	assert.Equal(t, existingUser, user)
}

func TestUserServiceLogin_EmailSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockUserRepository(ctrl)
	userService := service.NewUserService(userRepo)

	pass, _ := utils.HashPassword("Ary4123#")

	existingUser := entity.User{
		Model: gorm.Model{
			ID: 1,
		},
		Password: pass,
	}

	userRepo.EXPECT().GetUserByEmail(gomock.Any(), "aryadevara@gmail.com").Return(existingUser, nil)

	userLogin := entity.User{
		Email:    "aryadevara@gmail.com",
		Password: "Ary4123#",
	}

	user, err := userService.LoginUser(context.Background(), &userLogin)
	assert.Nil(t, err)
	assert.Equal(t, existingUser, user)
}

func TestUserServiceLogin_PasswordNotMatchErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockUserRepository(ctrl)
	userService := service.NewUserService(userRepo)

	pass, _ := utils.HashPassword("Ary4123#")

	existingUser := entity.User{
		Model: gorm.Model{
			ID: 1,
		},
		Password: pass,
	}

	userRepo.EXPECT().GetUserByEmail(gomock.Any(), "aryadevara@gmail.com").Return(existingUser, nil)

	userLogin := entity.User{
		Email:    "aryadevara@gmail.com",
		Password: "Ary4123!",
	}

	_, err := userService.LoginUser(context.Background(), &userLogin)
	assert.Equal(t, service.ErrUserPasswordDontMatch, err)

}

func TestUserServiceLogin_UserNotFoundErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockUserRepository(ctrl)
	userService := service.NewUserService(userRepo)

	existingUser := entity.User{}

	userRepo.EXPECT().GetUserByEmail(gomock.Any(), "aryadevara@gmail.com").Return(existingUser, errors.New("db error"))

	userLogin := entity.User{
		Email:    "aryadevara@gmail.com",
		Password: "Ary4123#",
	}

	_, err := userService.LoginUser(context.Background(), &userLogin)
	assert.Equal(t, service.ErrUserNotFound, err)
}
