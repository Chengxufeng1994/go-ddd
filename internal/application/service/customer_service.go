package service

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/dto"
	"github.com/Chengxufeng1994/go-ddd/internal/application/usecase"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/repository"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/valueobject"
)

type CustomerService struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository) usecase.CustomerUseCase {
	return &CustomerService{
		customerRepository: customerRepository,
	}
}

// CreateCustomer implements usecase.CustomerUseCase.
func (c *CustomerService) CreateCustomer(ctx context.Context, req *dto.CustomerCreationRequest) (*dto.CustomerCreationResponse, error) {
	email, _ := valueobject.NewEmail(req.Email)
	entity := &entity.Customer{
		Email:        email,
		CustomerInfo: valueobject.NewCustomerInfo(req.Age, req.FirstName, req.LastName),
	}
	_, err := c.customerRepository.CreateCustomer(ctx, entity)
	if err != nil {
		return nil, err
	}

	return &dto.CustomerCreationResponse{}, nil
}
