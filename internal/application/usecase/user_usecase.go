package usecase

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/application/dto"
	"github.com/Chengxufeng1994/go-ddd/internal/application/service/command"
	"github.com/Chengxufeng1994/go-ddd/internal/application/service/query"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/repository"
)

type UserUseCase interface {
	CreateUser(context.Context, command.CreateUserCommand) (*command.CreateUserCommandResult, error)
	GetUser(context.Context, query.GetUserQuery) (*query.GetUserQueryResult, error)
	ListUsers(context.Context, *dto.UserQueryParams) (*[]dto.User, *repository.PaginationResult, error)
	SearchUsers(context.Context, *dto.UserQueryParams) (*[]dto.User, *repository.PaginationResult, error)
	AddAccountWithUser(context.Context, *dto.AccountCreationRequest) (*dto.AccountCreationResponse, error)
}
