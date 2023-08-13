package service

import (
	"context"
	"dating-app/entity"
	"dating-app/repository"
	"dating-app/utils"
	"errors"
	"net"
	"strings"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrUserPasswordDontMatch = errors.New("password not match")
	ErrUserAlreadyExists     = errors.New("user already exists")
	ErrEmailInvalid          = errors.New("domain of email invalid")
)

func (s *UserService) Register(ctx context.Context, userReq entity.User) error {
	existingUser, err := s.userRepository.GetUserByUsername(ctx, userReq.Username)
	if err != nil {
		return err
	}

	if existingUser.ID != 0 {
		return ErrUserAlreadyExists
	}

	domain := strings.Split(userReq.Email, "@")
	_, err = net.LookupMX(domain[1])
	if err != nil {
		return ErrEmailInvalid
	}

	userReq.Password, err = utils.HashPassword(userReq.Password)
	if err != nil {
		return err
	}

	err = s.userRepository.Register(ctx, &userReq)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) LoginUser(ctx context.Context, userReq entity.User) (user entity.User, err error) {
	if userReq.Username != "" {
		user, err = s.userRepository.GetUserByUsername(ctx, userReq.Username)
	} else {
		user, err = s.userRepository.GetUserByEmail(ctx, userReq.Email)
	}

	if err != nil {
		return entity.User{}, ErrUserNotFound
	}

	if utils.CheckPassword(userReq.Password, user.Password) != nil {
		return entity.User{}, ErrUserPasswordDontMatch
	}

	return user, nil
}
