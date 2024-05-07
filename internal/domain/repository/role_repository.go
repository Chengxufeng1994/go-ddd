package repository

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
)

type RoleRepository interface {
	CreateRole(context.Context, *entity.Role) (*entity.Role, error)
}
