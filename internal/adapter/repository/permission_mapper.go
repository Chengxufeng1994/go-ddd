package repository

import (
	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
	"github.com/Chengxufeng1994/go-ddd/internal/infrastructure/persistence/po"
	"gorm.io/gorm"
)

type PermissionMapper struct{}

func NewPermissionMapper() *PermissionMapper {
	return &PermissionMapper{}
}

func (m *PermissionMapper) ToDatabaseModel(entity *entity.Permission) *po.Permission {
	return &po.Permission{
		Model: gorm.Model{
			ID:        entity.ID,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		},
		Name: entity.Name,
		Slug: entity.Slug,
	}
}

func (m *PermissionMapper) ToDomainEntity(model *po.Permission) *entity.Permission {
	return &entity.Permission{
		ID:        model.ID,
		Name:      model.Name,
		Slug:      model.Slug,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
