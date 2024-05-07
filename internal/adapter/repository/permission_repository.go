package repository

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/repository"
	"github.com/Chengxufeng1994/go-ddd/internal/infrastructure/persistence/po"
	"gorm.io/gorm"
)

type GormPermissionRepository struct {
	db               *gorm.DB
	permissionMapper *PermissionMapper
}

func NewGormPermissionRepository(db *gorm.DB) repository.PermissionRepository {
	return &GormPermissionRepository{
		db:               db,
		permissionMapper: NewPermissionMapper(),
	}
}

func (r *GormPermissionRepository) CreatePermission(ctx context.Context, entity *entity.Permission) (*entity.Permission, error) {
	model := r.permissionMapper.ToDatabaseModel(entity)
	if err := r.db.WithContext(ctx).Model(&po.Permission{}).Create(model).Error; err != nil {
		return nil, err
	}

	return r.permissionMapper.ToDomainEntity(model), nil
}

func (r *GormPermissionRepository) AssignPermissionsToRole(ctx context.Context, roleID uint, permIDs []uint) error {
	tx := r.db.Begin()
	for _, permID := range permIDs {
		if err := r.db.WithContext(ctx).Model(&po.RolePermission{}).Create(&po.RolePermission{
			RoleID:       roleID,
			PermissionID: permID,
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
