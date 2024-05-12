package repository

import (
	"context"
	"fmt"
	"math"

	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/repository"
	"github.com/Chengxufeng1994/go-ddd/internal/infrastructure/persistence/po"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormUserRepository struct {
	db         *gorm.DB
	userMapper *UserMapper
}

func NewGormUserRepository(db *gorm.DB) repository.UserRepository {
	return &GormUserRepository{
		db:         db,
		userMapper: NewUserMapper(),
	}
}

func (r *GormUserRepository) CreateUser(ctx context.Context, entity *entity.User) (*entity.User, error) {
	model := r.userMapper.ToDatabaseModel(entity)
	err := r.db.WithContext(ctx).Model(&po.User{}).Create(model).Error
	if err != nil {
		return nil, err
	}

	return r.userMapper.ToDomainEntity(model), nil
}

func (r *GormUserRepository) ListUsers(ctx context.Context, page repository.PaginationCriteria) (*entity.Users, *repository.PaginationResult, error) {
	var rows []po.User
	err := r.db.WithContext(ctx).Scopes(repository.Pagination(&page)).Model(&po.User{}).Order("created_at asc").Find(&rows).Error
	if err != nil {
		return nil, nil, err
	}

	totalRows, err := r.Count(ctx)
	page.TotalRows = totalRows

	totalPages := int(math.Ceil(float64(totalRows) / float64(page.Limit)))
	page.TotalPages = totalPages

	if err != nil {
		return nil, nil, err
	}

	entities := make(entity.Users, 0, len(rows))
	for i := 0; i < len(rows); i++ {
		entities = append(entities, *r.userMapper.ToDomainEntity(&rows[i]))
	}

	return &entities, &repository.PaginationResult{
		PaginationCriteria: page,
	}, nil
}

func (r *GormUserRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.Table("users").WithContext(ctx).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *GormUserRepository) SearchUsers(ctx context.Context, criteria repository.UserSearchCriteria) (*entity.Users, *repository.PaginationResult, error) {
	db := r.db.WithContext(ctx).Model(&po.User{})
	if criteria.Email != "" {
		db = db.Where("email LIKE ?", "%"+criteria.Email+"%")
	}

	var rows []po.User
	err := db.WithContext(ctx).
		Model(&po.User{}).
		Scopes(repository.Pagination(&criteria.PaginationCriteria)).
		Order(fmt.Sprintf("%s %s", criteria.OrderByCriteria.SortBy, criteria.OrderByCriteria.OrderBy)).
		Find(&rows).Error
	if err != nil {
		return nil, nil, err
	}

	totalRows, err := r.Count(ctx)
	criteria.PaginationCriteria.TotalRows = totalRows

	totalPages := int(math.Ceil(float64(totalRows) / float64(criteria.PaginationCriteria.Limit)))
	criteria.PaginationCriteria.TotalPages = totalPages

	if err != nil {
		return nil, nil, err
	}

	entities := make(entity.Users, 0, len(rows))
	for i := 0; i < len(rows); i++ {
		entities = append(entities, *r.userMapper.ToDomainEntity(&rows[i]))
	}

	return &entities, &repository.PaginationResult{
		PaginationCriteria: criteria.PaginationCriteria,
	}, nil
}

func (r *GormUserRepository) GetUser(ctx context.Context, id uint) (*entity.User, error) {
	var row po.User
	err := r.db.WithContext(ctx).Model(&po.User{}).Where("id = ?", id).First(&row).Error
	if err != nil {
		return nil, err
	}

	return r.userMapper.ToDomainEntity(&row), nil
}

func (r *GormUserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var row po.User
	err := r.db.WithContext(ctx).Model(&po.User{}).Where("email = ?", email).First(&row).Error
	if err != nil {
		return nil, err
	}

	return r.userMapper.ToDomainEntity(&row), nil
}

func (r *GormUserRepository) UpdateUser(ctx context.Context, id uint, entity *entity.User) (*entity.User, error) {
	model := r.userMapper.ToDatabaseModel(entity)
	err := r.db.WithContext(ctx).Clauses(clause.Returning{}).Model(&po.User{}).Where("id = ?", id).Updates(model).Error
	if err != nil {
		return nil, err
	}

	return r.userMapper.ToDomainEntity(model), nil
}

func (r *GormUserRepository) DeleteUser(ctx context.Context, id uint) (*entity.User, error) {
	model := &po.User{}
	model.ID = uint(id)
	err := r.db.WithContext(ctx).Clauses(clause.Returning{}).Model(&po.User{}).Delete(model).Error
	if err != nil {
		return nil, err
	}

	return r.userMapper.ToDomainEntity(model), nil
}
