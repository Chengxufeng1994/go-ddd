package repository

import (
	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/valueobject"
	"github.com/Chengxufeng1994/go-ddd/internal/infrastructure/persistence/po"
	"gorm.io/gorm"
)

type AccountMapper struct{}

func NewAccountMapper() *AccountMapper {
	return &AccountMapper{}
}

func (m *AccountMapper) ToDatabaseModel(entity *entity.Account) *po.Account {
	return &po.Account{
		Model: gorm.Model{
			ID:        entity.ID,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		},
		CustomerID: entity.CustomerID,
		Amount:     entity.Money.Amount(),
		Currency:   entity.Money.Currency(),
	}
}

func (m *AccountMapper) ToDomainEntity(model *po.Account) *entity.Account {
	money, _ := valueobject.NewMoney(model.Amount, model.Currency)
	return &entity.Account{
		ID:         model.ID,
		CustomerID: model.CustomerID,
		Money:      money,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	}
}
