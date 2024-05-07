package entity

import (
	"time"

	"github.com/Chengxufeng1994/go-ddd/internal/domain/valueobject"
)

type User struct {
	ID             uint
	Active         bool
	Email          *valueobject.Email
	HashedPassword string
	UserInfo       *valueobject.CustomerInfo
	CreatedAt      time.Time
	UpdatedAt      time.Time
	RoleID         uint
	Roles          []Role
}

type Users []User
