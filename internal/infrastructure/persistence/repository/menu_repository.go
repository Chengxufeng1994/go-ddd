package repository

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/repository"
	"github.com/Chengxufeng1994/go-ddd/internal/infrastructure/persistence/po"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormMenuRepository struct {
	db         *gorm.DB
	menuMapper *MenuMapper
}

func NewGormMenuRepository(db *gorm.DB) repository.MenuRepository {
	return &GormMenuRepository{
		db:         db,
		menuMapper: NewMenuMapper(),
	}
}

// Create implements repository.MenuRepository.
func (r *GormMenuRepository) Create(ctx context.Context, entity *entity.Menu) (*entity.Menu, error) {
	model := r.menuMapper.ToDatabaseModel(entity)
	err := r.db.WithContext(ctx).Model(&po.Menu{}).Create(&model).Error
	if err != nil {
		return nil, err
	}

	return r.menuMapper.ToDomainEntity(model), nil
}

// Get implements repository.MenuRepository.
func (r *GormMenuRepository) Get(ctx context.Context, id uint) (*entity.Menu, error) {
	var row po.Menu
	err := r.db.WithContext(ctx).Model(&po.Menu{}).Where("id = ?", id).First(&row).Error
	if err != nil {
		return nil, err
	}

	return r.menuMapper.ToDomainEntity(&row), nil
}

// Update implements repository.MenuRepository.
func (r *GormMenuRepository) Update(ctx context.Context, id uint, entity *entity.Menu) (*entity.Menu, error) {
	model := r.menuMapper.ToDatabaseModel(entity)
	err := r.db.WithContext(ctx).Clauses(clause.Returning{}).Model(&po.Menu{}).Where("id = ?", id).Updates(model).Error
	if err != nil {
		return nil, err
	}

	return r.menuMapper.ToDomainEntity(model), nil
}

// Delete implements repository.MenuRepository.
func (r *GormMenuRepository) Delete(ctx context.Context, id uint) (*entity.Menu, error) {
	model := &po.Menu{}
	model.ID = uint(id)
	err := r.db.WithContext(ctx).Clauses(clause.Returning{}).Model(&po.Menu{}).Delete(model).Error
	if err != nil {
		return nil, err
	}

	return r.menuMapper.ToDomainEntity(model), nil
}
