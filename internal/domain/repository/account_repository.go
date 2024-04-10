package repository

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
)

type AccountRepository interface {
	ListAccounts(context.Context) (entity.Accounts, error)
	GetAccount(context.Context, uint) (*entity.Account, error)
	CreateAccount(context.Context, *entity.Account) (*entity.Account, error)
	UpdateAccount(context.Context, uint, *entity.Account) (*entity.Account, error)
	DeleteAccount(context.Context, uint) (*entity.Account, error)
}
