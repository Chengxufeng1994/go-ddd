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
	accountRepository  repository.AccountRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository, accountRepository repository.AccountRepository) usecase.CustomerUseCase {
	return &CustomerService{
		customerRepository: customerRepository,
		accountRepository:  accountRepository,
	}
}

// CreateCustomer implements usecase.CustomerUseCase.
func (s *CustomerService) CreateCustomer(ctx context.Context, req *dto.CustomerCreationRequest) (*dto.CustomerCreationResponse, error) {
	email, _ := valueobject.NewEmail(req.Email)
	entity := &entity.Customer{
		Email:          email,
		HashedPassword: req.Password,
		CustomerInfo:   valueobject.NewCustomerInfo(req.Age, req.FirstName, req.LastName),
	}

	rcustomer, err := s.customerRepository.CreateCustomer(ctx, entity)
	if err != nil {
		return nil, err
	}

	return &dto.CustomerCreationResponse{
		ID:        rcustomer.ID,
		Active:    rcustomer.Active,
		Email:     rcustomer.Email.String(),
		Age:       rcustomer.CustomerInfo.Age(),
		FirstName: rcustomer.CustomerInfo.FirstName(),
		LastName:  rcustomer.CustomerInfo.LastName(),
		CreatedAt: rcustomer.CreatedAt,
		UpdatedAt: rcustomer.UpdatedAt,
	}, nil
}

// GetCustomer implements usecase.CustomerUseCase.
func (s *CustomerService) GetCustomer(ctx context.Context, ID uint) (*dto.Customer, error) {
	rcustomer, err := s.customerRepository.GetCustomer(ctx, ID)
	if err != nil {
		return nil, err
	}

	return &dto.Customer{
		ID:        rcustomer.ID,
		Active:    rcustomer.Active,
		Email:     rcustomer.Email.String(),
		Age:       rcustomer.CustomerInfo.Age(),
		FirstName: rcustomer.CustomerInfo.FirstName(),
		LastName:  rcustomer.CustomerInfo.LastName(),
		CreatedAt: rcustomer.CreatedAt,
		UpdatedAt: rcustomer.UpdatedAt,
	}, nil
}

// SearchCustomers implements usecase.CustomerUseCase.
func (s *CustomerService) ListCustomers(ctx context.Context) (*[]dto.Customer, error) {
	criteria := repository.CustomerSearchCriteria{}
	res, err := s.customerRepository.SearchCustomers(ctx, criteria)
	if err != nil {
		return nil, err
	}
	dtos := make([]dto.Customer, 0, len(*res))
	for _, cus := range *res {
		dtos = append(dtos, dto.Customer{
			ID: cus.ID,
		})
	}

	return &dtos, nil
}

// AddAccountWithCustomer implements usecase.CustomerUseCase.
func (s *CustomerService) AddAccountWithCustomer(ctx context.Context, req *dto.AccountCreationRequest) (*dto.AccountCreationResponse, error) {
	existed, err := s.customerRepository.GetCustomer(ctx, req.CustomerID)
	if existed == nil || err != nil {
		return nil, err
	}

	money, _ := valueobject.NewMoney(req.Amount, req.Currency)
	entity := &entity.Account{
		CustomerID: req.CustomerID,
		Money:      money,
	}

	raccount, err := s.accountRepository.CreateAccount(ctx, entity)
	if err != nil {
		return nil, err
	}

	return &dto.AccountCreationResponse{
		ID:         raccount.ID,
		CustomerID: raccount.CustomerID,
		Amount:     raccount.Money.Amount(),
		Currency:   raccount.Money.Currency(),
		CreatedAt:  raccount.CreatedAt,
		UpdatedAt:  raccount.UpdatedAt,
	}, nil
}
