package usecase

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/application/dto"
)

type MenuUseCase interface {
	GetMenu(context.Context, uint) (*dto.Menu, error)
	CreateMenu(context.Context, *dto.MenuCreationRequest) (*dto.Menu, error)
	UpdateMenu(context.Context, uint, *dto.Menu) (*dto.Menu, error)
	DeleteMenu(context.Context, uint) error
}
