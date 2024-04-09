package repository

import (
	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/valueobject"
	"github.com/Chengxufeng1994/go-ddd/internal/infrastructure/persistence/po"
	"gorm.io/gorm"
)

type CustomerMapper struct{}

func NewCustomMapper() *CustomerMapper {
	return &CustomerMapper{}
}

func (m *CustomerMapper) ToDatabaseModel(entity *entity.Customer) *po.Customer {
	return &po.Customer{
		Model: gorm.Model{
			ID:        entity.ID,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		},
		Active:    entity.Active,
		Email:     entity.Email.String(),
		Age:       entity.CustomerInfo.Age(),
		FirstName: entity.CustomerInfo.FirstName(),
		LastName:  entity.CustomerInfo.LastName(),
	}
}

func (m *CustomerMapper) ToDomainEntity(model *po.Customer) *entity.Customer {
	email, _ := valueobject.NewEmail(model.Email)
	return &entity.Customer{
		ID:           model.ID,
		Active:       model.Active,
		Email:        email,
		CustomerInfo: valueobject.NewCustomerInfo(model.Age, model.FirstName, model.LastName),
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
	}
}
