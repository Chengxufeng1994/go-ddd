package repository

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/repository"
	"github.com/Chengxufeng1994/go-ddd/internal/infrastructure/persistence/po"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormAccountRepository struct {
	db            *gorm.DB
	accountMapper *AccountMapper
}

func NewGormAccountRepository(db *gorm.DB) repository.AccountRepository {
	return &GormAccountRepository{
		db:            db,
		accountMapper: NewAccountMapper(),
	}
}

// CreateAccount implements repository.AccountRepository.
func (r *GormAccountRepository) CreateAccount(ctx context.Context, entity *entity.Account) (*entity.Account, error) {
	model := r.accountMapper.ToDatabaseModel(entity)
	err := r.db.WithContext(ctx).Model(&po.Account{}).Create(model).Error
	if err != nil {
		return nil, err
	}

	return r.accountMapper.ToDomainEntity(model), nil
}

// ListAccounts implements repository.AccountRepository.
func (r *GormAccountRepository) ListAccounts(context.Context) (entity.Accounts, error) {
	panic("unimplemented")
}

// GetAccount implements repository.AccountRepository.
func (r *GormAccountRepository) GetAccount(ctx context.Context, ID uint) (*entity.Account, error) {
	var row *po.Account
	if err := r.db.WithContext(ctx).Model(&po.Account{}).Where("ID = ?", ID).First(&row).Error; err != nil {
		return nil, err
	}
	// if err := r.db.WithContext(ctx).Preload("Customer").Model(&po.Account{}).Where("ID = ?", ID).First(&row).Error; err != nil {
	// 	return nil, err
	// }

	return r.accountMapper.ToDomainEntity(row), nil
}

// UpdateAccount implements repository.AccountRepository.
func (r *GormAccountRepository) UpdateAccount(ctx context.Context, ID uint, entity *entity.Account) (*entity.Account, error) {
	model := r.accountMapper.ToDatabaseModel(entity)
	err := r.db.WithContext(ctx).Clauses(clause.Returning{}).Model(&po.Account{}).Where("id = ?", ID).Updates(model).Error
	if err != nil {
		return nil, err
	}

	return r.accountMapper.ToDomainEntity(model), nil
}

// DeleteAccount implements repository.AccountRepository.
func (r *GormAccountRepository) DeleteAccount(context.Context, uint) (*entity.Account, error) {
	panic("unimplemented")
}
