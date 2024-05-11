package usecase

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/application/dto"
)

type AuthUseCase interface {
	SignUp(context.Context, *dto.SignUpRequest) error
	SignIn(context.Context, *dto.SignInRequest) error
	SignOut(context.Context, *dto.SignOutRequest) error
}
