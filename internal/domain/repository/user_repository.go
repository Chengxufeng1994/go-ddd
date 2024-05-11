package repository

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
)

type UserSearchCriteria struct {
	Age int
}

type UserRepository interface {
	ListUsers(context.Context, PaginationCriteria) (*PaginationResult, error)
	SearchUsers(context.Context, UserSearchCriteria) (*entity.Users, error)
	GetUser(context.Context, uint) (*entity.User, error)
	GetUserByEmail(context.Context, string) (*entity.User, error)
	CreateUser(context.Context, *entity.User) (*entity.User, error)
	UpdateUser(context.Context, uint, *entity.User) (*entity.User, error)
	DeleteUser(context.Context, uint) (*entity.User, error)
}
