package repository

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/domain/repository"
	"gorm.io/gorm"
)

// UnitOfWorkBlock is a func that accepts a UnitOfWorkStore, which creates
// a relationship between UnitOfWork (generic struct) with our application's
// specific datastores (repositories)
type UnitOfWorkBlock func(IUnitOfWorkStore) error

type UnitOfWork interface {
	Do(context.Context, UnitOfWorkBlock) error
}

type unitOfWork struct {
	db *gorm.DB
}

func New(db *gorm.DB) UnitOfWork {
	return &unitOfWork{
		db: db,
	}
}

// Do implements UnitOfWork.
// Do executes the given UnitOfWorkBlock atomically (inside a DB transaction).
func (u *unitOfWork) Do(ctx context.Context, fn UnitOfWorkBlock) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		// create a new UnitOfWorkStore, it holds references to the application's
		// repositories that must share that same transaction
		accountRepository := NewGormAccountRepository(tx)
		transferRepository := NewGormTransferRepository(tx)
		store := NewUnitOfWorkStore(accountRepository, transferRepository)

		// execute the given UnitOfWorkBlock passing in a UnitOfWorkStore which,
		// in turn, gives access to all repositories sharing the same DB transaction
		return fn(store)
	})
}

// UnitOfWorkStore provides access to datastores that can be
// used inside an UnitOfWorkBlock. All data changes done through
// them will share the same DB transaction.
type IUnitOfWorkStore interface {
	Account() repository.AccountRepository
	Transfer() repository.TransferRepository
}

type UnitOfWorkStore struct {
	accountRepository  repository.AccountRepository
	transferRepository repository.TransferRepository
}

func NewUnitOfWorkStore(
	accountRepository repository.AccountRepository,
	transferRepository repository.TransferRepository,
) IUnitOfWorkStore {
	return &UnitOfWorkStore{
		accountRepository:  accountRepository,
		transferRepository: transferRepository,
	}
}

// Account implements IUnitOfWorkStore.
func (u *UnitOfWorkStore) Account() repository.AccountRepository {
	return u.accountRepository
}

// Transfer implements IUnitOfWorkStore.
func (u *UnitOfWorkStore) Transfer() repository.TransferRepository {
	return u.transferRepository
}
