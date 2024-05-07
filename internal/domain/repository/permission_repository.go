package repository

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
)

type PermissionRepository interface {
	CreatePermission(context.Context, *entity.Permission) (*entity.Permission, error)
	AssignPermissionsToRole(context.Context, uint, []uint) error
}
