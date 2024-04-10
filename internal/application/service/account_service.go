package service

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/dto"
	"github.com/Chengxufeng1994/go-ddd/internal/application/usecase"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/repository"
)

type AccountService struct {
	accountRepository  repository.AccountRepository
	customerRepository repository.CustomerRepository
}

func NewAccountService(accountRepository repository.AccountRepository, customerRepository repository.CustomerRepository) usecase.AccountUseCase {
	return &AccountService{
		accountRepository:  accountRepository,
		customerRepository: customerRepository,
	}
}

// CreateAccount implements usecase.AccountUseCase.
// func (s *AccountService) CreateAccount(ctx context.Context, req *dto.AccountCreationRequest) (*dto.AccountCreationResponse, error) {
// 	existed, err := s.customerRepository.GetCustomerByEmail(ctx, req.Owner)
// 	if existed == nil || err != nil {
// 		return nil, err
// 	}

// 	owner, _ := valueobject.NewEmail(req.Owner)
// 	money, _ := valueobject.NewMoney(req.Amount, req.Currency)
// 	entity := &entity.Account{
// 		Owner: owner,
// 		Money: money,
// 	}

// 	raccount, err := s.accountRepository.CreateAccount(ctx, entity)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &dto.AccountCreationResponse{
// 		ID:        raccount.ID,
// 		Owner:     raccount.Owner.String(),
// 		Amount:    raccount.Money.Amount(),
// 		Currency:  raccount.Money.Currency(),
// 		CreatedAt: raccount.CreatedAt,
// 		UpdatedAt: raccount.UpdatedAt,
// 	}, nil
// }

// GetAccount implements usecase.AccountUseCase.
func (a *AccountService) GetAccount(ctx context.Context, ID uint) (*dto.AccountResponse, error) {
	raccount, err := a.accountRepository.GetAccount(ctx, ID)
	if err != nil {
		return nil, err
	}

	return &dto.AccountResponse{
		ID:         raccount.ID,
		CustomerID: raccount.CustomerID,
		Amount:     raccount.Money.Amount(),
		Currency:   raccount.Money.Currency(),
		CreatedAt:  raccount.CreatedAt,
		UpdatedAt:  raccount.UpdatedAt,
	}, nil
}
