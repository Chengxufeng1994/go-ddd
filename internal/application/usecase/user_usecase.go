package usecase

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/application/dto"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/repository"
)

type UserUseCase interface {
	CreateUser(context.Context, *dto.UserCreationRequest) (*dto.UserCreationResponse, error)
	GetUser(context.Context, uint) (*dto.User, error)
	ListUsers(context.Context, *dto.UserQueryParams) (*[]dto.User, *repository.PaginationResult, error)
	SearchUsers(context.Context, *dto.UserQueryParams) (*[]dto.User, *repository.PaginationResult, error)
	AddAccountWithUser(context.Context, *dto.AccountCreationRequest) (*dto.AccountCreationResponse, error)
}
