package service

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/application/dto"
	"github.com/Chengxufeng1994/go-ddd/internal/application/usecase"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/repository"
)

type MenuService struct {
	menuRepository repository.MenuRepository
}

func NewMenuService(menuRepository repository.MenuRepository) usecase.MenuUseCase {
	return &MenuService{
		menuRepository: menuRepository,
	}
}

func (m *MenuService) GetMenu(ctx context.Context, menuID uint) (*dto.Menu, error) {
	rmenu, err := m.menuRepository.Get(ctx, menuID)
	if err != nil {
		return nil, err
	}

	return &dto.Menu{
		ID:           rmenu.ID,
		Name:         rmenu.Name,
		Slug:         rmenu.Slug,
		Description:  rmenu.Description,
		Path:         rmenu.Path,
		ParentID:     rmenu.ParentID,
		CreatedAt:    rmenu.CreatedAt,
		UpdatedAt:    rmenu.UpdatedAt,
		ChildrenMenu: nil,
	}, nil
}

func (m *MenuService) CreateMenu(ctx context.Context, req *dto.MenuCreationRequest) (*dto.Menu, error) {
	rmenu, err := m.menuRepository.Create(ctx, &entity.Menu{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		Path:        req.Path,
		ParentID:    req.ParentID,
	})
	if err != nil {
		return nil, err
	}

	return &dto.Menu{
		ID:           rmenu.ID,
		Name:         rmenu.Name,
		Slug:         rmenu.Slug,
		Description:  rmenu.Description,
		Path:         rmenu.Path,
		ParentID:     rmenu.ParentID,
		CreatedAt:    rmenu.CreatedAt,
		UpdatedAt:    rmenu.UpdatedAt,
		ChildrenMenu: nil,
	}, nil
}

func (m *MenuService) UpdateMenu(ctx context.Context, menuID uint, req *dto.Menu) (*dto.Menu, error) {
	panic("unimplemented")
}

func (m *MenuService) DeleteMenu(ctx context.Context, menuID uint) error {
	panic("unimplemented")
}
