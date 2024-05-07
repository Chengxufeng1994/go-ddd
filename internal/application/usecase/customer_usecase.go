package usecase

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/dto"
)

type CustomerUseCase interface {
	CreateCustomer(context.Context, *dto.CustomerCreationRequest) (*dto.CustomerCreationResponse, error)
	GetCustomer(context.Context, uint) (*dto.Customer, error)
	ListCustomers(context.Context) (*[]dto.Customer, error)
	AddAccountWithCustomer(context.Context, *dto.AccountCreationRequest) (*dto.AccountCreationResponse, error)
}