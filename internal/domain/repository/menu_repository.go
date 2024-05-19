package repository

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
)

type MenuRepository interface {
	Get(context.Context, uint) (*entity.Menu, error)
	Create(context.Context, *entity.Menu) (*entity.Menu, error)
	Update(context.Context, uint, *entity.Menu) (*entity.Menu, error)
	Delete(context.Context, uint) (*entity.Menu, error)
}
