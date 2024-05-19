package repository

import (
	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
	"github.com/Chengxufeng1994/go-ddd/internal/infrastructure/persistence/po"
	"gorm.io/gorm"
)

type MenuMapper struct{}

func NewMenuMapper() *MenuMapper {
	return &MenuMapper{}
}

func (m *MenuMapper) ToDatabaseModel(entity *entity.Menu) *po.Menu {
	return &po.Menu{
		Model: gorm.Model{
			ID:        entity.ID,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		},
		Name:        entity.Name,
		Slug:        entity.Slug,
		Description: entity.Description,
		Path:        entity.Path,
		ParentID:    entity.ParentID,
	}
}

func (m *MenuMapper) ToDomainEntity(model *po.Menu) *entity.Menu {

	return &entity.Menu{
		ID:          model.ID,
		Name:        model.Name,
		Slug:        model.Slug,
		Description: model.Description,
		Path:        model.Path,
		ParentID:    model.ParentID,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}
