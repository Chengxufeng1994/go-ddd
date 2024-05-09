package repository

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
)

type RBACRepository interface {
	ListRolePermissions(ctx context.Context, roleID uint) ([]*entity.Permission, error)
}
