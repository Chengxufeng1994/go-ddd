package usecase

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/dto"
)

type CustomerUseCase interface {
	CreateCustomer(context.Context, *dto.CustomerCreationRequest) (*dto.CustomerCreationResponse, error)
}
