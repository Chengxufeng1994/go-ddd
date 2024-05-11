package repository

import (
	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/valueobject"
	"github.com/Chengxufeng1994/go-ddd/internal/infrastructure/persistence/po"
	"gorm.io/gorm"
)

type UserMapper struct{}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (m *UserMapper) ToDatabaseModel(entity *entity.User) *po.User {
	var roles []po.Role
	for _, r := range entity.Roles {
		roles = append(roles, po.Role{
			Model: gorm.Model{
				ID: r.ID,
			},
			Name: r.Name,
			Slug: r.Slug,
		})
	}
	return &po.User{
		Model: gorm.Model{
			ID:        entity.ID,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		},
		Active:         entity.Active,
		Email:          entity.Email.String(),
		HashedPassword: entity.HashedPassword,
		Age:            entity.UserInfo.Age(),
		FirstName:      entity.UserInfo.FirstName(),
		LastName:       entity.UserInfo.LastName(),
		RoleID:         entity.RoleID,
		Roles:          roles,
	}
}

func (m *UserMapper) ToDomainEntity(model *po.User) *entity.User {
	email, _ := valueobject.NewEmail(model.Email)
	var roles []entity.Role
	for _, r := range model.Roles {
		roles = append(roles, entity.Role{
			ID:   r.ID,
			Name: r.Name,
			Slug: r.Slug,
		})
	}

	return &entity.User{
		ID:             model.ID,
		Active:         model.Active,
		Email:          email,
		HashedPassword: model.HashedPassword,
		UserInfo:       valueobject.NewUserInfo(model.Age, model.FirstName, model.LastName),
		CreatedAt:      model.CreatedAt,
		UpdatedAt:      model.UpdatedAt,
		RoleID:         model.RoleID,
		Roles:          roles,
	}
}
