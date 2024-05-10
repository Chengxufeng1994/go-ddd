package repository

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/repository"
	"github.com/Chengxufeng1994/go-ddd/internal/infrastructure/persistence/po"
	"gorm.io/gorm"
)

type RBACRepository struct {
	db               *gorm.DB
	permissionMapper *PermissionMapper
}

func NewRBACRepository(db *gorm.DB) repository.RBACRepository {
	return &RBACRepository{
		db:               db,
		permissionMapper: NewPermissionMapper(),
	}
}

// ListRolePermissions implements repository.RBACRepository.
func (r *RBACRepository) ListRolePermissions(ctx context.Context, roleID uint) ([]*entity.Permission, error) {
	var role po.Role
	if err := r.db.WithContext(ctx).Preload("Permissions").Model(&po.Role{}).Where("id = ?", roleID).First(&role).Error; err != nil {
		return nil, err
	}

	var entities []*entity.Permission
	perms := role.Permissions
	for _, perm := range perms {
		entities = append(entities, r.permissionMapper.ToDomainEntity(&perm))
	}

	return entities, nil
}
