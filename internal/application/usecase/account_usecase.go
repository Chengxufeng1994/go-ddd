package usecase

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/application/dto"
)

type AccountUseCase interface {
	// CreateAccount(context.Context, *dto.AccountCreationRequest) (*dto.AccountCreationResponse, error)
	GetAccount(context.Context, uint) (*dto.AccountResponse, error)
}
