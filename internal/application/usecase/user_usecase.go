package usecase

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/application/dto"
)

type UserUseCase interface {
	CreateUser(context.Context, *dto.UserCreationRequest) (*dto.UserCreationResponse, error)
	GetUser(context.Context, uint) (*dto.User, error)
	ListUsers(context.Context) (*[]dto.User, error)
	AddAccountWithUser(context.Context, *dto.AccountCreationRequest) (*dto.AccountCreationResponse, error)
}
