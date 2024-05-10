package service

import (
	"context"
	"errors"

	"github.com/Chengxufeng1994/go-ddd/internal/application/dto"
	"github.com/Chengxufeng1994/go-ddd/internal/application/usecase"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/repository"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/valueobject"
	infra_repository "github.com/Chengxufeng1994/go-ddd/internal/infrastructure/persistence/repository"
)

type TransactionService struct {
	trxRepository     repository.TransferRepository
	accountRepository repository.AccountRepository
	unitOfWork        infra_repository.UnitOfWork
}

func NewTransactionService(trxRepository repository.TransferRepository, accountRepository repository.AccountRepository, unitOfWork infra_repository.UnitOfWork) usecase.TransactionUseCase {
	return &TransactionService{
		trxRepository:     trxRepository,
		accountRepository: accountRepository,
		unitOfWork:        unitOfWork,
	}
}

// Transfer implements usecase.TransactionUseCase.
func (s *TransactionService) Transfer(ctx context.Context, req *dto.TransferRequest) (*dto.TransferResponse, error) {
	s.trxRepository.CreateTransfer(ctx, &entity.Transfer{
		FromAccountId: req.FromAccountId,
		ToAccountId:   req.ToAccountId,
		Amount:        req.Amount,
	})

	fromAccount, err := s.accountRepository.GetAccount(ctx, req.FromAccountId)
	if err != nil {
		return nil, errors.New("from account not found")
	}

	toAccount, err := s.accountRepository.GetAccount(ctx, req.ToAccountId)
	if err != nil {
		return nil, errors.New("to account not found")
	}

	if ok := fromAccount.Money.IsCurrencyEquals(*toAccount.Money); !ok {
		return nil, errors.New("currency not equal")
	}

	reqMoney, err := valueobject.NewMoney(req.Amount, fromAccount.Money.Currency())
	if err != nil {
		return nil, err
	}

	fromAccount.Money, err = fromAccount.Money.Deduct(*reqMoney)
	if err != nil {
		return nil, err
	}

	toAccount.Money, err = toAccount.Money.Add(*reqMoney)
	if err != nil {
		return nil, err
	}

	rFromAccount, err := s.accountRepository.UpdateAccount(ctx, req.FromAccountId, fromAccount)
	if err != nil {
		return nil, err
	}

	rToAccount, err := s.accountRepository.UpdateAccount(ctx, req.ToAccountId, toAccount)
	if err != nil {
		return nil, err
	}

	return &dto.TransferResponse{
		FromAccountId: rFromAccount.ID,
		ToAccountId:   rToAccount.ID,
	}, nil
}

// TransferWithTrx implements usecase.TransactionUseCase.
func (s *TransactionService) TransferWithTrx(ctx context.Context, req *dto.TransferRequest) (*dto.TransferResponse, error) {
	var rFromAccount *entity.Account
	var rToAccount *entity.Account

	fn := func(store infra_repository.IUnitOfWorkStore) error {
		store.Transfer().CreateTransfer(ctx, &entity.Transfer{
			FromAccountId: req.FromAccountId,
			ToAccountId:   req.ToAccountId,
			Amount:        req.Amount,
		})

		fromAccount, err := store.Account().GetAccount(ctx, req.FromAccountId)
		if err != nil {
			return errors.New("from account not found")
		}

		toAccount, err := store.Account().GetAccount(ctx, req.ToAccountId)
		if err != nil {
			return errors.New("to account not found")
		}

		if ok := fromAccount.Money.IsCurrencyEquals(*toAccount.Money); !ok {
			return errors.New("currency not equal")
		}

		reqMoney, err := valueobject.NewMoney(req.Amount, fromAccount.Money.Currency())
		if err != nil {
			return err
		}

		fromAccount.Money, err = fromAccount.Money.Deduct(*reqMoney)
		if err != nil {
			return err
		}

		toAccount.Money, err = toAccount.Money.Add(*reqMoney)
		if err != nil {
			return err
		}

		rFromAccount, err = store.Account().UpdateAccount(ctx, req.FromAccountId, fromAccount)
		if err != nil {
			return err
		}

		rToAccount, err = store.Account().UpdateAccount(ctx, req.ToAccountId, toAccount)
		if err != nil {
			return err
		}

		return nil
	}

	if err := s.unitOfWork.Do(ctx, fn); err != nil {
		return nil, err
	}

	return &dto.TransferResponse{
		FromAccountId: rFromAccount.ID,
		ToAccountId:   rToAccount.ID,
	}, nil
}
