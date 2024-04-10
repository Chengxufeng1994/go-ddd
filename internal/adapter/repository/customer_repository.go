package repository

import (
	"context"
	"math"

	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/repository"
	"github.com/Chengxufeng1994/go-ddd/internal/infrastructure/persistence/po"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormCustomerRepository struct {
	db             *gorm.DB
	customerMapper *CustomerMapper
}

func NewGormCustomerRepository(db *gorm.DB) repository.CustomerRepository {
	return &GormCustomerRepository{
		db:             db,
		customerMapper: NewCustomMapper(),
	}
}

// CreateCustomer implements repository.CustomerRepository.
func (r *GormCustomerRepository) CreateCustomer(ctx context.Context, entity *entity.Customer) (*entity.Customer, error) {
	model := r.customerMapper.ToDatabaseModel(entity)
	err := r.db.WithContext(ctx).Model(&po.Customer{}).Create(model).Error
	if err != nil {
		return nil, err
	}

	return r.customerMapper.ToDomainEntity(model), nil
}

// ListCustomers implements repository.CustomerRepository.
func (r *GormCustomerRepository) ListCustomers(ctx context.Context, page repository.PaginationCriteria) (*repository.PaginationResult, error) {
	var rows []po.Customer
	err := r.db.WithContext(ctx).Scopes(repository.Pagination(&page)).Model(&po.Customer{}).Find(&rows).Error
	if err != nil {
		return nil, err
	}

	totalRows, err := r.Count(ctx)
	page.TotalRows = totalRows

	totalPages := int(math.Ceil(float64(totalRows) / float64(page.Limit)))
	page.TotalPages = totalPages

	if err != nil {
		return nil, err
	}

	entities := make(entity.Customers, 0, len(rows))
	for i := 0; i < len(rows); i++ {
		entities = append(entities, *r.customerMapper.ToDomainEntity(&rows[i]))
	}

	return &repository.PaginationResult{
		PaginationCriteria: page,
		Rows:               entities,
	}, nil
}

// Count returns the total of sources.
func (r *GormCustomerRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.Table("customers").WithContext(ctx).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// SearchCustomers implements repository.CustomerRepository.
func (r *GormCustomerRepository) SearchCustomers(context.Context, repository.CustomerSearchCriteria) (*entity.Customers, error) {
	panic("unimplemented")
}

// GetCustomer implements repository.CustomerRepository.
func (r *GormCustomerRepository) GetCustomer(ctx context.Context, id uint) (*entity.Customer, error) {
	var row po.Customer
	err := r.db.WithContext(ctx).Model(&po.Customer{}).Where("id = ?", id).First(&row).Error
	if err != nil {
		return nil, err
	}

	return r.customerMapper.ToDomainEntity(&row), nil
}

// GetCustomerByEmail implements repository.CustomerRepository.
func (r *GormCustomerRepository) GetCustomerByEmail(ctx context.Context, email string) (*entity.Customer, error) {
	var row po.Customer
	err := r.db.WithContext(ctx).Model(&po.Customer{}).Where("email = ?", email).First(&row).Error
	if err != nil {
		return nil, err
	}

	return r.customerMapper.ToDomainEntity(&row), nil
}

// UpdateCustomer implements repository.CustomerRepository.
func (r *GormCustomerRepository) UpdateCustomer(ctx context.Context, id uint, entity *entity.Customer) (*entity.Customer, error) {
	model := r.customerMapper.ToDatabaseModel(entity)
	err := r.db.WithContext(ctx).Clauses(clause.Returning{}).Model(&po.Customer{}).Where("id = ?", id).Updates(model).Error
	if err != nil {
		return nil, err
	}

	return r.customerMapper.ToDomainEntity(model), nil
}

// DeleteCustomer implements repository.CustomerRepository.
func (r *GormCustomerRepository) DeleteCustomer(ctx context.Context, id uint) (*entity.Customer, error) {
	model := &po.Customer{}
	model.ID = uint(id)
	err := r.db.WithContext(ctx).Clauses(clause.Returning{}).Model(&po.Customer{}).Delete(model).Error
	if err != nil {
		return nil, err
	}

	return r.customerMapper.ToDomainEntity(model), nil
}
