package repository

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/repository"
	"github.com/Chengxufeng1994/go-ddd/internal/infrastructure/persistence/po"
	"gorm.io/gorm"
)

type GormRoleRepository struct {
	db         *gorm.DB
	roleMapper *RoleMapper
}

func NewGormRoleRepository(db *gorm.DB) repository.RoleRepository {
	return &GormRoleRepository{
		db:         db,
		roleMapper: NewRoleMapper(),
	}
}

func (r *GormRoleRepository) CreateRole(ctx context.Context, entity *entity.Role) (*entity.Role, error) {
	model := r.roleMapper.ToDatabaseModel(entity)
	if err := r.db.WithContext(ctx).Model(&po.Role{}).Create(model).Error; err != nil {
		return nil, err
	}

	return r.roleMapper.ToDomainEntity(model), nil
}
