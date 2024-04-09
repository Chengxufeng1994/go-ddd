package repository

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
)

type CustomerRepository interface {
	ListCustomers(context.Context) (entity.Customers, error)
	GetCustomer(context.Context, uint) (*entity.Customer, error)
	CreateCustomer(context.Context, *entity.Customer) (*entity.Customer, error)
	UpdateCustomer(context.Context, uint, *entity.Customer) (*entity.Customer, error)
	DeleteCustomer(context.Context, uint) (*entity.Customer, error)
}
