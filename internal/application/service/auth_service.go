package service

import (
	"context"
	"errors"

	"github.com/Chengxufeng1994/go-ddd/internal/application/dto"
	"github.com/Chengxufeng1994/go-ddd/internal/application/usecase"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/repository"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/valueobject"
	"github.com/Chengxufeng1994/go-ddd/pkg"
)

var ErrUserNotFound = errors.New("user not found")

type AuthService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) usecase.AuthUseCase {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (svc *AuthService) SignUp(ctx context.Context, req *dto.SignUpRequest) error {
	email, err := valueobject.NewEmail(req.Email)
	if err != nil {
		return err
	}

	hashedPassword, err := pkg.HashedPassword(req.Password)
	if err != nil {
		return err
	}

	userInfo := valueobject.NewUserInfo(0, req.FirstName, req.LastName)

	toCreate := &entity.User{
		Email:          email,
		HashedPassword: hashedPassword,
		UserInfo:       userInfo,
		RoleID:         3,
	}

	_, err = svc.userRepository.CreateUser(ctx, toCreate)
	if err != nil {
		return err
	}

	return nil
}

func (svc *AuthService) SignIn(ctx context.Context, req *dto.SignInRequest) error {
	email, err := valueobject.NewEmail(req.Email)
	if err != nil {
		return err
	}

	existed, err := svc.userRepository.GetUserByEmail(ctx, email.String())
	if err != nil {
		return err
	}
	if existed == nil {
		return ErrUserNotFound
	}

	if ok := pkg.ComparePassword(existed.HashedPassword, req.Password); !ok {
		return err
	}

	return nil
}

func (svc *AuthService) SignOut(ctx context.Context, req *dto.SignOutRequest) error {
	panic("unimplemented")
}
