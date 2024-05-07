package service

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/dto"
	"github.com/Chengxufeng1994/go-ddd/internal/application/usecase"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/repository"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/valueobject"
)

type UserService struct {
	customerRepository repository.CustomerRepository
	accountRepository  repository.AccountRepository
}

func NewUserService(customerRepository repository.CustomerRepository, accountRepository repository.AccountRepository) usecase.UserUseCase {
	return &UserService{
		customerRepository: customerRepository,
		accountRepository:  accountRepository,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *dto.UserCreationRequest) (*dto.UserCreationResponse, error) {
	email, _ := valueobject.NewEmail(req.Email)
	entity := &entity.User{
		Email:          email,
		HashedPassword: req.Password,
		CustomerInfo:   valueobject.NewCustomerInfo(req.Age, req.FirstName, req.LastName),
	}

	rcustomer, err := s.customerRepository.CreateUser(ctx, entity)
	if err != nil {
		return nil, err
	}

	return &dto.UserCreationResponse{
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

func (s *UserService) GetUser(ctx context.Context, ID uint) (*dto.User, error) {
	rcustomer, err := s.customerRepository.GetUser(ctx, ID)
	if err != nil {
		return nil, err
	}

	return &dto.User{
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

func (s *UserService) ListUsers(ctx context.Context) (*[]dto.User, error) {
	criteria := repository.CustomerSearchCriteria{}
	res, err := s.customerRepository.SearchUsers(ctx, criteria)
	if err != nil {
		return nil, err
	}
	dtos := make([]dto.User, 0, len(*res))
	for _, cus := range *res {
		dtos = append(dtos, dto.User{
			ID: cus.ID,
		})
	}

	return &dtos, nil
}

func (s *UserService) AddAccountWithUser(ctx context.Context, req *dto.AccountCreationRequest) (*dto.AccountCreationResponse, error) {
	existed, err := s.customerRepository.GetUser(ctx, req.UserID)
	if existed == nil || err != nil {
		return nil, err
	}

	money, _ := valueobject.NewMoney(req.Amount, req.Currency)
	entity := &entity.Account{
		UserID: req.UserID,
		Money:  money,
	}

	raccount, err := s.accountRepository.CreateAccount(ctx, entity)
	if err != nil {
		return nil, err
	}

	return &dto.AccountCreationResponse{
		ID:        raccount.ID,
		UserID:    raccount.UserID,
		Amount:    raccount.Money.Amount(),
		Currency:  raccount.Money.Currency(),
		CreatedAt: raccount.CreatedAt,
		UpdatedAt: raccount.UpdatedAt,
	}, nil
}
